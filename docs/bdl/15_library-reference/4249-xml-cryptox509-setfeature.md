---
title: "xml.CryptoX509.setFeature"
source: "fgl-topics/c_gws_XmlCryptoX509_setFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.setFeature"
type: "concept"
---

# xml.CryptoX509.setFeature

> Sets or resets the given feature for this CryptoX509 object.

## Syntax

```
setFeature(
   feature STRING,
   value STRING )
```

1. feature defines the [feature](4250-cryptox509-features.md "Features of the xml.CryptoX509 class.") to be set.
2. value defines the value to set.

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related reference**  

[CryptoX509 Features](4250-cryptox509-features.md "Features of the xml.CryptoX509 class.")
