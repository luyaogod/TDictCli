---
title: "xml.CryptoKey.savePublicToString"
source: "fgl-topics/c_gws_XmlCryptoKey_savePublicToString.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.savePublicToString"
type: "concept"
---

# xml.CryptoKey.savePublicToString

> Save the current xml.CryptoKey's public part in the returned base64 string.

## Syntax

```
savePublicToString()
  RETURNS STRING
```

## Usage

Returns the public part of the key in base64 form (`STRING`).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
