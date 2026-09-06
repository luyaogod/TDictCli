---
title: "Create an enveloped signature using a RSA key"
source: "fgl-topics/c_gws_XmlSignature_example_create_enveloping_signature_RSA.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Create an enveloped signature using a RSA key"
type: "concept"
---

# Create an enveloped signature using a RSA key

> In this code sample, a XML document ("MyDocument.xml") is loaded and signed with a RSA key.

You can use the sample content provided in [XML document (unsigned)](4305-xml-document-unsigned.md "Sample content provided for the purpose of testing examples.") for the purpose of testing the code. Copy
the content to a file named "MyDocument.xml" in a directory where you test the sample code.

A RSA key ("RSAKey.pem") is used to sign the document. You will need to
generate this key. The key can be created with the OpenSSL
tool.

```
openssl genrsa -out RSAKey.pem 2048
```

The number
2048 is the size of the key in bits.

Copy the file "RSAKey.pem" to a directory where you test the sample code.

All keys or certificates in PEM or DER format were created with the OpenSSL tool.
For information on how the OpenSSL tool works, refer to the [openssl](https://www.openssl.org/) (external link) documentation.

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE doc2 xml.DomDocument
  DEFINE root xml.DomNode
  DEFINE node xml.DomNode
  DEFINE signNode xml.DomNode
  DEFINE sig xml.Signature
  DEFINE key xml.CryptoKey
  DEFINE index INTEGER
  # Create DomDocument object
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography, 
  # therefore it is recommended to remove unnecessary ones 
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load document to be signed
    CALL doc.load("MyDocument.xml")
    # Create rsa key
    LET key = xml.CryptoKey.Create("http://www.w3.org/2000/09/xmldsig#rsa-sha1")
    CALL key.loadPEM("RSAKey.pem")
    # Create signature object with the key to use
    LET sig = xml.Signature.Create()
    CALL sig.setKey(key)
    # Set XML node to be signed. In our case, the node with 
    # attribute 'xml:id="code"'
    LET index = sig.createReference("#code",
      "http://www.w3.org/2000/09/xmldsig#sha1")
    # Add enveloped method to not take the XML signature node into account 
    # when computing the entire document.
    CALL sig.appendReferenceTransformation(index,
            "http://www.w3.org/2000/09/xmldsig#enveloped-signature",doc.getDocumentElement())
    # Set canonicalization method on the XML fragment to be signed.
    CALL sig.appendReferenceTransformation(index,
            "http://www.w3.org/2001/10/xml-exc-c14n#")
    # Compute enveloped signature
    CALL sig.compute(doc)
    # Retrieve signature document
    LET doc2=sig.getDocument()
    # Append the signature node to the original document to get 
    # a valid enveloped signature
    # Notice that the enveloped signature can be added anywhere in the 
    # original document
    LET signNode = doc2.getDocumentElement() # Get Signature node
    # Import it into the original document
    LET node = doc.importNode(signNode,true) 
    # Retrieve the original document root node
    LET root = doc.getDocumentElement() 
    # Append the signature node as last child of the original document
    CALL root.appendChild(node) 
    # Save document with enveloped signature back to disk
    CALL doc.setFeature("format-pretty-print",TRUE)
    CALL doc.save("MyDocumentEnvelopedSignature.xml")
  CATCH
    DISPLAY "Unable to create an enveloped signature :",status
  END TRY
END MAIN
```

For an example of the output produced by this code, see [XML document (signed with RSA key)](4308-xml-document-signed-with-rsa-key.md "The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a RSA key) might look like.").

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")

[XML document (signed with RSA key)](4308-xml-document-signed-with-rsa-key.md "The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a RSA key) might look like.")
