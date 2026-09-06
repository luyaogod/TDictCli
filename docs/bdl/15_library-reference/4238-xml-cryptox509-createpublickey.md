---
title: "xml.CryptoX509.createPublicKey"
source: "fgl-topics/c_gws_XmlCryptoX509_createPublicKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.createPublicKey"
type: "concept"
---

# xml.CryptoX509.createPublicKey

> Creates a new public CryptoKey object for the given URL.

## Syntax

```
createPublicKey(
   url STRING )
  RETURNS xml.CryptoKey
```

1. url defines the given [url](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.").

## Usage

This method creates a new public CryptoKey object for the given URL, from the public key embedded
in this certificate if any; NULL otherwise.

Returns a [xml.CryptoKey](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
