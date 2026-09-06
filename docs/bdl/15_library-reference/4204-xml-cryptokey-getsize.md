---
title: "xml.CryptoKey.getSize"
source: "fgl-topics/c_gws_XmlCryptoKey_getSize.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.getSize"
type: "concept"
---

# xml.CryptoKey.getSize

> Returns the size of the key in bits.

## Syntax

```
getSize()
  RETURNS INTEGER
```

## Usage

For a Diffie-Hellman key, it returns the size of the key; the size of a Diffie-Hellman key is
actually the size of the modulus. If the modulus is not available (null or equal to zero), the
method returns zero. In this situation, a return of zero does NOT mean the key is corrupt or
unusable.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related reference**  

[Supported kind of keys](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.")
