---
title: "xml.Encryption.getEmbeddedKey"
source: "fgl-topics/c_gws_XmlEncryption_getEmbeddedKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.getEmbeddedKey"
type: "concept"
---

# xml.Encryption.getEmbeddedKey

> Get a copy of the embedded symmetric key that was used in the last decryption operation.

## Syntax

```
getEmbeddedKey()
  RETURNS xml.CryptoKey
```

## Usage

Returns a copy of the embedded symmetric [key](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") that was used in the last decryption operation, or NULL if there is none.

An embedded symmetric key is always encrypted, and needs therefore a key-encryption key to be set
in order to decrypt it.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
