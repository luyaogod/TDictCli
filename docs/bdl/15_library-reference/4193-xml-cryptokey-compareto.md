---
title: "xml.CryptoKey.compareTo"
source: "fgl-topics/c_gws_XmlCryptoKey_compareTo.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.compareTo"
type: "concept"
---

# xml.CryptoKey.compareTo

> Compares a CryptoKey object to a second key.

## Syntax

```
compareTo(
   toCompare xml.CryptoKey )
  RETURNS INTEGER
```

1. toCompare defines the [xml.CryptoKey](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") object to use for
   comparison to the current CryptoKey object.

## Usage

The method verifies if the key's URL, type, size, usage and value are the same. If they are the
same, the two identical keys will produce the same encryption cipher.

The key features are not taken into account during comparison.

Returns `TRUE` if they are identical, `FALSE` if they are not
identical.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related reference**  

[CryptoKey Features](4224-cryptokey-features.md "Features of the xml.CryptoKey class.")
