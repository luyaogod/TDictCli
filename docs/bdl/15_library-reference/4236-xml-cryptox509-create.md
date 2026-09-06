---
title: "xml.CryptoX509.Create"
source: "fgl-topics/c_gws_XmlCryptoX509_Create.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.Create"
type: "concept"
---

# xml.CryptoX509.Create

> Constructor of an empty CryptoX509 object.

## Syntax

```
xml.CryptoX509.Create()
  RETURNS xml.CryptoX509
```

## Usage

Returns a [xml.CryptoX509](4234-the-cryptox509-class.md "The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.") object or NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
