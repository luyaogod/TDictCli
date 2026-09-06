---
title: "xml.CryptoKey.loadPublicFromString"
source: "fgl-topics/c_gws_XmlCryptoKey_loadPublicFromString.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.loadPublicFromString"
type: "concept"
---

# xml.CryptoKey.loadPublicFromString

> Populate the current CryptoKey object with the passed public key.

## Syntax

```
loadPublicFromString(
   str STRING )
```

1. str defines the public part of the
   key in base64 form.

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
