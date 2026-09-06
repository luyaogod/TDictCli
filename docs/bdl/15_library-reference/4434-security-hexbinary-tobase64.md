---
title: "security.HexBinary.ToBase64"
source: "fgl-topics/c_gws_SecurityHexBinary_ToBase64.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.ToBase64"
type: "concept"
---

# security.HexBinary.ToBase64

> Converts an hexadecimal string to the base64 equivalent

## Syntax

```
security.HexBinary.ToBase64(
   hexBinVal STRING )
  RETURNS STRING
```

1. hexBinVal defines the string in hexadecimal
   form

## Usage

Decodes the given hexadecimal string and returns it in base64.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
