---
title: "xml.CryptoKey.saveToString"
source: "fgl-topics/c_gws_XmlCryptoKey_saveToString.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.saveToString"
type: "concept"
---

# xml.CryptoKey.saveToString

> Saves the CryptoKey object into a BASE64 string format.

## Syntax

```
saveToString()
  RETURNS STRING
```

## Usage

For Diffie-Hellman, this method returns the Diffie-Hellman key's modulus and generator in a
base64 encoded string. This is used for the parameters exchange step between the two peers.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
