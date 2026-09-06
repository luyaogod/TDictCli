---
title: "xml.CryptoKey.setKey"
source: "fgl-topics/c_gws_XmlCryptoKey_setKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.setKey"
type: "concept"
---

# xml.CryptoKey.setKey

> Defines the value of a HMAC or Symmetric key.

## Syntax

```
setKey(
   key STRING )
```

1. key defines the value.

## Usage

The value can be a password and must be of the size corresponding to the key [identifier](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
