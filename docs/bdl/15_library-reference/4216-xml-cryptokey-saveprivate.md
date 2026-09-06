---
title: "xml.CryptoKey.savePrivate"
source: "fgl-topics/c_gws_XmlCryptoKey_savePrivate.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.savePrivate"
type: "concept"
---

# xml.CryptoKey.savePrivate

> Saves the private key part of an asymmetric RSA CryptoKey object into a XML document according to the XKMS2.0 specification.

## Syntax

```
savePrivate()
  RETURNS xml.DomDocument
```

## Usage

Returns a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.")
object containing the private key part of an asymmetric RSA key.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
