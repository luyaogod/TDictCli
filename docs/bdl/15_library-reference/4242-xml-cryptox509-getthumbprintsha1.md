---
title: "xml.CryptoX509.getThumbprintSHA1"
source: "fgl-topics/c_gws_XmlCryptoX509_getThumbprintSHA1.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.getThumbprintSHA1"
type: "concept"
---

# xml.CryptoX509.getThumbprintSHA1

> Gets the SHA1 encoded thumbprint identifying the X509 certificate.

## Syntax

```
getThumbprintSHA1()
  RETURNS STRING
```

## Usage

This method returns the SHA1 encoded thumbprint identifying the X509 certificate in a
BASE64-encoded string.

Example: `CM4y6z7zzLnTGMe1lE46RKIKAPI=`

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
