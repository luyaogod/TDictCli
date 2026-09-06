---
title: "xml.CryptoX509.getFeature"
source: "fgl-topics/c_gws_XmlCryptoX509_getFeature.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.getFeature"
type: "concept"
---

# xml.CryptoX509.getFeature

> Get the value of a given feature of a CryptoX509 object.

## Syntax

```
getFeature(
   feature STRING )
  RETURNS STRING
```

1. feature defines a [feature](4250-cryptox509-features.md "Features of the xml.CryptoX509 class.") of
   the CryptoX509 object.

## Usage

This method returns the value of the given feature for the CryptoX509 object, or NULL if feature
is not set.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related reference**  

[CryptoX509 Features](4250-cryptox509-features.md "Features of the xml.CryptoX509 class.")
