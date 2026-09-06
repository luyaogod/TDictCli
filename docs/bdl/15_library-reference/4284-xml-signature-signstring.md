---
title: "xml.Signature.SignString"
source: "fgl-topics/c_gws_XmlSignature_signString.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.SignString"
type: "concept"
---

# xml.Signature.SignString

> Sign the passed string according to the specified key.

## Syntax

```
SignString(
  key xml.CryptoKey,
  str STRING
 )
 RETURNS STRING
```

1. key defines the [key](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") to be used for the
   signature.
2. str defines the string to be
   signed.

## Usage

The key can be a HMAC key, a RSA private key, or a DSA private key. The digest used for signing
is determined by the key algorithm.

Returns sig, or the signature in base64 format.

This method does not belong to the XML encryption specification.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
