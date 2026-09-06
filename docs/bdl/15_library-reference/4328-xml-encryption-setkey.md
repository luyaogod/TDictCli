---
title: "xml.Encryption.setKey"
source: "fgl-topics/c_gws_XmlEncryption_setKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.setKey"
type: "concept"
---

# xml.Encryption.setKey

> Assigns a copy of the symmetric key to this encryption object.

## Syntax

```
setKey(
   key xml.CryptoKey )
```

1. key defines the symmetric [key](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.").

## Usage

Any further XML document or node encryption or decryption will use that symmetric key.

When decrypting a XML document that has an embedded symmetric key, the embedded key will be used
instead.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
