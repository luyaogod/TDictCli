---
title: "xml.CryptoKey.getFeature"
source: "fgl-topics/c_gws_XmlCryptoKey_getFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.getFeature"
type: "concept"
---

# xml.CryptoKey.getFeature

> Returns the value of the given feature for this CryptoKey object, or NULL.

## Syntax

```
getFeature(
   feature STRING )
  RETURNS STRING
```

1. [feature](4224-cryptokey-features.md "Features of the xml.CryptoKey class.") defines the
   CryptoKey feature.

## Usage

Returns NULL if the feature is not set.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related reference**  

[CryptoKey Features](4224-cryptokey-features.md "Features of the xml.CryptoKey class.")
