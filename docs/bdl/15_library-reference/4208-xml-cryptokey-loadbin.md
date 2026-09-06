---
title: "xml.CryptoKey.loadBin"
source: "fgl-topics/c_gws_XmlCryptoKey_loadBIN.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.loadBin"
type: "concept"
---

# xml.CryptoKey.loadBin

> Loads a symmetric or HMAC key from a file in raw format.

## Syntax

```
loadBin(
   filename STRING )
```

1. filename defines the filename or an [entry](../16_web-services/4916-fglprofile-entries-for-web-services.md) in the FGLPROFILE file.

## Usage

Raw format means that data in the file is read without any transformation, and will be stored as
is in the key.

For instance, if your file contains "hello", it has the same effect as calling
`xml.CryptoKey.setKey()` with "hello" as parameter.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
