---
title: "xml.Encryption.decryptKey"
source: "fgl-topics/c_gws_XmlEncryption_decryptKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.decryptKey"
type: "concept"
---

# xml.Encryption.decryptKey

> Decrypts the EncryptedKey as root in the given XML document, and returns a new CryptoKey of the given kind.

## Syntax

```
decryptKey(
   doc xml.DomDocument,
   url STRING )
  RETURNS xml.CryptoKeY
```

1. doc defines the [DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.") object.
2. url defines the string.

## Usage

Returns a new [CryptoKey](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") of the
given kind.

Only symmetric or HMAC keys are allowed.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
