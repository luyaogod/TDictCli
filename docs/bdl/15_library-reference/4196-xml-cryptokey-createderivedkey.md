---
title: "xml.CryptoKey.CreateDerivedKey"
source: "fgl-topics/c_gws_XmlCryptoKey_CreateDerivedKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.CreateDerivedKey"
type: "concept"
---

# xml.CryptoKey.CreateDerivedKey

> Constructor of an empty CryptoKey object based on a URL. The crypto key must be derived before use.

## Syntax

```
xml.CryptoKey.CreateDerivedKey(
   url STRING )
  RETURNS xml.CryptoKey
```

1. [url](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.") defines a key identifier based on the XML-Signature and
   XML-Encryption specification.

## Usage

Returns a [xml.CryptoKey](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") object or
NULL. Only symmetric and HMAC keys can be derived.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[xml.CryptoKey.deriveKey](4198-xml-cryptokey-derivekey.md "Derives the symmetric or HMAC CryptoKey object using the given method identifier and concatenating the optional label, the mandatory seed value and the optional created date as initial random value.")
