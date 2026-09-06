---
title: "Create an enveloping signature using a DSA key"
source: "fgl-topics/c_gws_XmlSignature_example_create_enveloping_signature_DSA.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Create an enveloping signature using a DSA key"
type: "concept"
---

# Create an enveloping signature using a DSA key

> Use this example to digitally sign an XML document with a DSA key, producing an enveloping signature that guarantees the document has not been tampered with.

An enveloping signature wraps the signed content inside the
`Signature` element, keeping the data and signature together in a
single output file. This example uses the DSA-SHA256 algorithm, which is the
recommended algorithm for DSA digital signatures.

You can use the sample content provided in [XML document (unsigned)](4305-xml-document-unsigned.md "Sample content provided for the purpose of testing examples.") for the purpose of testing the code. Copy
the content to a file named "MyDocument.xml" in a directory where you test the sample code.

The sample code signs the XML document with a DSA private key named
DSAKey.pem. Before running the example, generate the key with
OpenSSL and copy it to the directory where you will run the sample code.

1. Generate the DSA parameters:

   ```
   openssl dsaparam -out DSAparam.pem 2048
   ```

   The value `2048` specifies the key size in bits.
2. Generate the private key and save it to DSAKey.pem:

   ```
   openssl gendsa -out DSAKey.pem DSAparam.pem
   ```

All keys or certificates in PEM or DER format were created with the OpenSSL tool.
For information on how the OpenSSL tool works, refer to the [openssl](https://www.openssl.org/) (external link) documentation.

Copy DSAKey.pem to the directory where you will run the
sample code.

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE sig xml.Signature
  DEFINE key xml.CryptoKey
  DEFINE index INTEGER
  DEFINE objInd INTEGER
  # Create DomDocument object
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography, 
  # therefore it is recommended to remove unnecessary ones 
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load document to be signed
    CALL doc.load("MyDocument.xml")
    # Create DSA key and load it from file
    LET key = xml.CryptoKey.Create(
      "http://www.w3.org/2009/xmldsig11#dsa-sha256")
    CALL key.loadPEM("DSAKey.pem")
    # Create signature object with the key to use
    LET sig = xml.Signature.Create()
    CALL sig.setKey(key)
    # Create an object inside the signature to envelop the root node
    LET objInd = sig.createObject()
    # Set the object id to get a reference
    CALL sig.setObjectId(objInd,"data")
    # Copy the enveloping node from the document
    CALL sig.appendObjectData(objInd,doc.getDocumentElement())
    # Set the reference to be signed on the object node.
    # In our case, the object node with attribute 'data'
    LET index = sig.createReference("#data",
      "http://www.w3.org/2001/04/xmlenc#sha256")
    # Set canonicalization method on the enveloping object to be signed.
    CALL sig.appendReferenceTransformation(index,
      "http://www.w3.org/2001/10/xml-exc-c14n#")
    # Compute enveloping signature
    CALL sig.compute(NULL)
    # Retrieve signature document
    LET doc=sig.getDocument()
    # Save signature on disk
    CALL doc.setFeature("format-pretty-print",TRUE)
    CALL doc.save("MyDocumentEnvelopingSignature.xml")
  CATCH
    DISPLAY "Unable to create an enveloping signature :",status
  END TRY
END MAIN
```

For an example of the output produced by this code, see [XML document (signed with DSA key)](4307-xml-document-signed-with-dsa-key.md "The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a DSA key) might look like.").

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")

[XML document (signed with DSA key)](4307-xml-document-signed-with-dsa-key.md "The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a DSA key) might look like.")

**Related reference**  

[Supported kind of keys](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.")
