---
title: "Verify an enveloped signature using a RSA key"
source: "fgl-topics/c_gws_XmlSignature_example_verify_enveloping_signature_RSA.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Verify an enveloped signature using a RSA key"
type: "concept"
---

# Verify an enveloped signature using a RSA key

> In this example, you verify the document (MyDocumentEnvelopedSignature.xml) signed with a RSA key.

The RSA key was created in [Create an enveloped signature using a RSA key](4302-create-an-enveloped-signature-using-a-rsa-key.md "In this code sample, a XML document (\"MyDocument.xml\") is loaded and signed with a RSA key.").

You need the RSA key ("RSAKey.pem") you used to sign the document in order
to verify it. Copy the file named "RSAKey.pem" to a directory where you test
the sample code.

All keys or certificates in PEM or DER format were created with the OpenSSL tool.
For information on how the OpenSSL tool works, refer to the [openssl](https://www.openssl.org/) (external link) documentation.

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE node xml.DomNode
  DEFINE sig xml.Signature
  DEFINE key xml.CryptoKey
  DEFINE list xml.DomNodeList
  DEFINE isVerified INTEGER
  # Create DomDocument object
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography, 
  # therefore it is recommended to remove unnecessary ones 
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load original document with enveloped signature into a DomDocument object
    CALL doc.load("MyDocumentEnvelopedSignature.xml")
    # Because the signature can be anywhere in the original document, 
    # we must first retrieve it
    LET list = doc.getElementsByTagNameNS("Signature",
      "http://www.w3.org/2000/09/xmldsig#")
    IF list.getCount() != 1 THEN
      DISPLAY "Unable to find one Signature node"
      EXIT PROGRAM (-1)
    ELSE
      LET node = list.getItem(1)
    END IF
    # Create RSA key
    LET key = xml.CryptoKey.Create(
      "http://www.w3.org/2000/09/xmldsig#rsa-sha1")
    CALL key.loadPEM("RSAKey.pem")
    # Create signature object from DomNode object and set RSA key to use
    LET sig = xml.Signature.CreateFromNode(node)
    CALL sig.setKey(key)
    # Verify enveloped signature validity of original document
    LET isVerified = sig.verify(doc)
    # Notice that if something has been modified in the node with 
    # attribute 'xml:id="code"' of the original XML document with the 
    # enveloped signature, the program will display "FAILED".
    IF isVerified THEN
      DISPLAY "Signature OK"
    ELSE
      DISPLAY "Signature FAILED"
    END IF
  CATCH
    DISPLAY "Unable to verify the enveloped signature :",status
  END TRY
END MAIN
```

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")
