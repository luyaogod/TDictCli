---
title: "xml.CryptoKey.loadPrivate"
source: "fgl-topics/c_gws_XmlCryptoKey_loadPrivate.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.loadPrivate"
type: "concept"
---

# xml.CryptoKey.loadPrivate

> Loads the private asymmetric RSA key from the given XML document.

## Syntax

```
loadPrivate(
   doc xml.DomDocument )
```

1. doc defines a `xml.DomDocument` object.

## Usage

This method loads the private asymmetric RSA key contained in the given
`xml.DomDocument` into the private key element of the CryptoKey object. The RSA
private key is based on the [XKMS2.0
specification](http://www.w3.org/TR/xkms2/#XKMS_2_0_Section_8_2).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
