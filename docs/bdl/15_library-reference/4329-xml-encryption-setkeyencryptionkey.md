---
title: "xml.Encryption.setKeyEncryptionKey"
source: "fgl-topics/c_gws_XmlEncryption_setKeyEncryptionKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.setKeyEncryptionKey"
type: "concept"
---

# xml.Encryption.setKeyEncryptionKey

> Assigns a copy of the key-encryption key to this encryption object.

## Syntax

```
setKeyEncryptionKey(
   key xml.CryptoKey )
```

1. key defines the key-encryption [key](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.").

## Usage

Any further XML encryption will use that key-encryption key to encrypt the symmetric key set with
`setKey()` within the resulting XML, and any further XML decryption will use that
key-encryption key to decrypt the embedded symmetric key.

- NULL is allowed, meaning that embedded symmetric keys will not be encrypted nor decrypted
  anymore, assuming that they have been exchanged in another way.
- Only public or private RSA keys, or key-wrap keys are allowed.
- Public RSA keys can encrypt but not decrypt.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
