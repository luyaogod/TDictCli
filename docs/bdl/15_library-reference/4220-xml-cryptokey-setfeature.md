---
title: "xml.CryptoKey.setFeature"
source: "fgl-topics/c_gws_XmlCryptoKey_setFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.setFeature"
type: "concept"
---

# xml.CryptoKey.setFeature

> Sets or resets the value of a feature for a CryptoKey object.

## Syntax

```
setFeature(
   feature STRING,
   value STRING )
```

1. [feature](4224-cryptokey-features.md "Features of the xml.CryptoKey class.") defines the name of the feature.
2. value defines the value to set for
   the named feature.

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related reference**  

[CryptoKey Features](4224-cryptokey-features.md "Features of the xml.CryptoKey class.")
