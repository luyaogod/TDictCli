---
title: "xml.CryptoKey.Create"
source: "fgl-topics/c_gws_XmlCryptoKey_Create.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.Create"
type: "concept"
---

# xml.CryptoKey.Create

> Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL.

## Syntax

```
xml.CryptoKey.Create(
   url STRING )
  RETURNS xml.CryptoKey
```

1. [url](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.") defines a key identifier based on the XML-Signature and
   XML-Encryption specification or the Diffie-Hellman specification.

## Usage

Returns a [xml.CryptoKey](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") object or
NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
