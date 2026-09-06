---
title: "Verify an enveloping signature using a X509 certificate"
source: "fgl-topics/c_gws_XmlSignature_example_verify_enveloping_signature_X509.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Verify an enveloping signature using a X509 certificate"
type: "concept"
---

# Verify an enveloping signature using a X509 certificate

> In this example, you verify the document (MyDocumentEnvelopingSignature.xml) signed with a DSA key.

The DSA key was created in [Create an enveloping signature using a DSA key](4300-create-an-enveloping-signature-using-a-dsa-key.md "Use this example to digitally sign an XML document with a DSA key, producing an enveloping signature that guarantees the document has not been tampered with.").

A X509 certificate ("DSACertificate.crt") can be used to verify the signed
document. The certificate will contain the public key corresponding to the private key
(DSAKey.pem) used to sign the document. You will need to generate this
certificate first. The certificate can be created with the OpenSSL tool:

```
openssl req -x509 -new -key DSAKey.pem -out DSACertificate.crt
```

You
can examine the contents of the certificate in plain text
using:

```
openssl x509 -in DSACertificate.crt --text
```

Copy
the file "DSACertificate.crt" to a directory where you test the sample
code.

All keys or certificates in PEM or DER format were created with the OpenSSL tool.
For information on how the OpenSSL tool works, refer to the [openssl](https://www.openssl.org/) (external link) documentation.

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE sig xml.Signature
  DEFINE cert xml.CryptoX509
  DEFINE pub xml.CryptoKey
  DEFINE isVerified INTEGER
  # Create DomDocument object
  LET doc = xml.DomDocument.Create()
  # Notice that whitespaces are significant in cryptography, 
  # therefore it is recommended to remove unnecessary ones 
  CALL doc.setFeature("whitespace-in-element-content",FALSE)
  TRY
    # Load Signature into a DomDocument object
    CALL doc.load("MyDocumentEnvelopingSignature.xml")
    # Create signature object from DomDocument root node
    LET sig = xml.Signature.CreateFromNode(doc.getDocumentElement())
    # Create X509 certificate 
    LET cert = xml.CryptoX509.Create()
    CALL cert.loadPEM("DSACertificate.crt")
    # Create public key from that X509 certificate 
    LET pub = cert.createPublicKey(
      "http://www.w3.org/2000/09/xmldsig#dsa-sha1")
    # Assign it to the signature 
    CALL sig.setKey(pub)
    # Verify enveloping signature validity 
    LET isVerified = sig.verify(NULL)
    # Notice that if something has been modified in the signature 
    # or if the certificate isn't associated with the 
    # private DSA key of example 3,
    # the program will display "FAILED".
    IF isVerified THEN
      DISPLAY "Signature OK"
    ELSE
      DISPLAY "Signature FAILED"
    END IF
  CATCH
    DISPLAY "Unable to verify the enveloping signature :",status
  END TRY
END MAIN
```

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")
