---
title: "xml.CryptoX509.getCount"
source: "fgl-topics/c_gws_XmlCryptoX509_getCount.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.getCount"
type: "concept"
---

# xml.CryptoX509.getCount

> Returns the number of X509 certificates contained in one CryptoX509 object.

## Syntax

```
getCount()
  RETURNS INTEGER
```

## Usage

Returns the number of X509 certificates contained in one CryptoX509 object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
