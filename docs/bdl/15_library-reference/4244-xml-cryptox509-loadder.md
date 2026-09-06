---
title: "xml.CryptoX509.loadDER"
source: "fgl-topics/c_gws_XmlCryptoX509_loadDER.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.loadDER"
type: "concept"
---

# xml.CryptoX509.loadDER

> Loads a X509 certificate from a file in DER format.

## Syntax

```
loadDER(
   filename STRING )
```

1. filename defines the filename or an [entry](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") in the FGLPROFILE
   file.

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
