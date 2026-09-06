---
title: "xml.Signature.setCertificate"
source: "fgl-topics/c_gws_XmlSignature_setCertificate.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.setCertificate"
type: "concept"
---

# xml.Signature.setCertificate

> Defines the X509 certificate to be added to the signature object when signing a document.

## Syntax

```
setCertificate(
   cert xml.CryptoX509 )
```

1. cert defines the [X509 certificate](4234-the-cryptox509-class.md "The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.") to be
   added.

## Usage

If NULL, no certificate is added.

During the computation of the signature, some certificate information can be added based on the
[feature](4250-cryptox509-features.md "Features of the xml.CryptoX509 class.") set on that CryptoX509 object. If no
features are set, the complete X509 certificate is automatically added.

During the verification of a signature the certificate set with the
`setCertificate` method isn't used. See [XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")
for more details.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
