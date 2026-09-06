---
title: "xml.CryptoKey.loadFromString"
source: "fgl-topics/c_gws_XmlCryptoKey_loadFromString.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.loadFromString"
type: "concept"
---

# xml.CryptoKey.loadFromString

> Loads the given key in BASE64 string format into a CryptoKey object.

## Syntax

```
loadFromString(
   str STRING )
```

1. str defines the string to load.

## Usage

For Diffie-Hellman, the input parameter is a base64 encoded string containing the Diffie-Hellman
parameters. This method populates the Diffie-Hellman key with the modulus and generator in the
base64 encoded string. This is useful for the parameters exchange step between two peers.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
