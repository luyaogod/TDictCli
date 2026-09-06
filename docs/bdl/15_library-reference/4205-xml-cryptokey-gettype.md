---
title: "xml.CryptoKey.getType"
source: "fgl-topics/c_gws_XmlCryptoKey_getType.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.getType"
type: "concept"
---

# xml.CryptoKey.getType

> Returns the type of key.

## Syntax

```
getType()
  RETURNS STRING
```

## Usage

This method returns the key type. It corresponds to the key type used by the [identifier](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related reference**  

[Supported kind of keys](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.")
