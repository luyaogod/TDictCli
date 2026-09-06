---
title: "Create a detached signature using a HMAC key"
source: "fgl-topics/c_gws_XmlSignature_example_create_detached_signature_HMAC.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Create a detached signature using a HMAC key"
type: "concept"
---

# Create a detached signature using a HMAC key

> In the example, an XML document ("MyDocument.xml") is loaded and signed with a HMAC key.

You can use the sample content provided in [XML document (unsigned)](4305-xml-document-unsigned.md "Sample content provided for the purpose of testing examples.") for the purpose of testing the code. Copy
the content to a file named "MyDocument.xml" in a directory where you test the sample code.

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE sig xml.Signature
  DEFINE key xml.CryptoKey
  DEFINE index INTEGER
  # Create DomDocument object
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography, 
  # therefore it is recommended that you remove unnecessary ones 
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load document to be signed
    CALL doc.load("MyDocument.xml")
    # Create HMAC key
    LET key = xml.CryptoKey.Create("http://www.w3.org/2000/09/xmldsig#hmac-sha1")
    CALL key.setKey("secretpassword")
     # Create signature object with the key to use
    LET sig = xml.Signature.Create()
    CALL sig.setKey(key)
    # Set XML node to be signed. In our case, the node with attribute
    # 'xml:id="code"'
    LET index = sig.createReference("#code",
      "http://www.w3.org/2000/09/xmldsig#sha1")
    # Set canonicalization method on the XML fragment to be signed.
    CALL sig.appendReferenceTransformation(index,
      "http://www.w3.org/2001/10/xml-exc-c14n#")
    # Compute detached signature
    CALL sig.compute(doc)
    # Retrieve signature document
    LET doc=sig.getDocument()
    # Save signature on disk
    CALL doc.setFeature("format-pretty-print",TRUE)
    CALL doc.save("MyDocumentDetachedSignature.xml")
  CATCH
    DISPLAY "Unable to create a detached signature :",status
  END TRY
END MAIN
```

All keys or certificates in PEM or DER format were created with the OpenSSL tool.
For information on how the OpenSSL tool works, refer to the [openssl](https://www.openssl.org/) (external link) documentation.

For an example of the output produced by this code, see [XML document (signed with HMAC key)](4306-xml-document-signed-with-hmac-key.md "The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a HMAC key) might look like.").

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")

[XML document (signed with HMAC key)](4306-xml-document-signed-with-hmac-key.md "The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a HMAC key) might look like.")
