---
title: "xml.Signature.setKey"
source: "fgl-topics/c_gws_XmlSignature_setKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.setKey"
type: "concept"
---

# xml.Signature.setKey

> Defines the key used for signing or validation.

## Syntax

```
setKey(
   key xml.CryptoKey )
```

1. key defines the [key](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") to be used for
   signing or validation.

## Usage

Only RSA, DSA or HMAC keys intended for SIGNATURE are allowed.

During the computation of the signature, some key information can be added depending on the [feature](4224-cryptokey-features.md "Features of the xml.CryptoKey class.") set on that CryptoKey object. If no features
are set, nothing is added. See [XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")
for more details.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
