---
title: "xml.CryptoX509.saveToString"
source: "fgl-topics/c_gws_XmlCryptoX509_saveToString.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.saveToString"
type: "concept"
---

# xml.CryptoX509.saveToString

> Saves the CryptoX509 certificate into a BASE64 string format.

## Syntax

```
saveToString()
  RETURNS STRING
```

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
