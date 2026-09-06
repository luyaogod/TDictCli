---
title: "Verify a detached signature using a HMAC key"
source: "fgl-topics/c_gws_XmlSignature_example_verify_detached_signature_HMAC.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Verify a detached signature using a HMAC key"
type: "concept"
---

# Verify a detached signature using a HMAC key

> In the example, you verify the document signed with an HMAC key.

The HMAC key was created in the sample [Create a detached signature using a HMAC key](4298-create-a-detached-signature-using-a-hmac-key.md "In the example, an XML document (\"MyDocument.xml\") is loaded and signed with a HMAC key."), against the original
unsigned document ("MyDocument.xml").

If you used the sample content provided in [XML document (unsigned)](4305-xml-document-unsigned.md "Sample content provided for the purpose of testing examples.") to create the signed document, then you must
verify the signature against this document to test the sample code.

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE sig xml.Signature
  DEFINE key xml.CryptoKey
  DEFINE isVerified INTEGER
  # Create DomDocument object
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography, 
  # therefore it is recommended to remove unnecessary ones 
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load Signature into a DomDocument object
    CALL doc.load("MyDocumentDetachedSignature.xml")
    # Create signature object from DomDocument root node
    LET sig = xml.Signature.CreateFromNode(doc.getDocumentElement())
    # Create HMAC key and assign it to the signature object
    LET key = xml.CryptoKey.Create("http://www.w3.org/2000/09/xmldsig#hmac-sha1")
    CALL key.setKey("secretpassword")
    CALL sig.setKey(key)
    # Load original XML document into a DomDocument object
    CALL doc.load("MyDocument.xml")
    # Verify detached signature validity of original document
    LET isVerified = sig.verify(doc)
    # Notice that if something has been modified in the node 
    # with attribute 'xml:id="code"' of the original XML document,
    # the program will display "FAILED".
    IF isVerified THEN
      DISPLAY "Signature OK"
    ELSE
      DISPLAY "Signature FAILED"
    END IF
  CATCH
    DISPLAY "Unable to verify the detached signature :",status
  END TRY
END MAIN
```

All keys or certificates in PEM or DER format were created with the OpenSSL tool.
For information on how the OpenSSL tool works, refer to the [openssl](https://www.openssl.org/) (external link) documentation.

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")
