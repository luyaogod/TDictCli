---
title: "xml.Encryption.encryptKey"
source: "fgl-topics/c_gws_XmlEncryption_encryptKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.encryptKey"
type: "concept"
---

# xml.Encryption.encryptKey

> Encrypts the given symmetric or HMAC key as an encrypted-key node and returns it as root node of a new XML document.

## Syntax

```
encryptKey(
   symkey xml.CryptoKey )
  RETURNS xml.DomDocument
```

1. symkey defines the given symmetric or
   HMAC [key](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") as an [EncryptedKey](http://www.w3.org/TR/xmlenc-core/#sec-EncryptedKey) node.

## Usage

This method encrypts the symmetric or HMAC key given in symkey as an
encrypted-key node and returns it as root node of a new [XML document](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards."). The encryption key must
have been set otherwise it will fail.

Depending on the feature set on the key-encryption key, the returned XML document will contain an
additional `KeyInfo` node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
