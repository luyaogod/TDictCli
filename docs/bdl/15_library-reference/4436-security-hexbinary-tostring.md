---
title: "security.HexBinary.ToString"
source: "fgl-topics/c_gws_SecurityHexBinary_ToString.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.ToString"
type: "concept"
---

# security.HexBinary.ToString

> Decodes an hexadecimal string to a clear, human-readable string.

## Syntax

```
security.HexBinary.ToString(
   hexVal STRING )
  RETURNS STRING
```

1. hexVal defines the string in hexadecimal.

## Usage

Decodes the given hexadecimal string and returns it in its clear, human readable, form. If the
hexadecimal string does not contain a human readable string, the method will raise an exception.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
