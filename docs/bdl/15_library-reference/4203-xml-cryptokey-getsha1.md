---
title: "xml.CryptoKey.getSHA1"
source: "fgl-topics/c_gws_XmlCryptoKey_getSHA1.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.getSHA1"
type: "concept"
---

# xml.CryptoKey.getSHA1

> Returns the SHA1 encoded key identifier in a base64 encoded STRING.

## Syntax

```
getSHA1()
  RETURNS STRING
```

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
