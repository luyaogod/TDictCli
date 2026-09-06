---
title: "security.HexBinary.FromString"
source: "fgl-topics/c_gws_SecurityHexBinary_FromString.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.FromString"
type: "concept"
---

# security.HexBinary.FromString

> Encodes a given string in hexadecimal.

## Syntax

```
security.HexBinary.FromString(
   clearVal STRING )
  RETURNS STRING
```

1. clearVal defines the string to be
   encoded.

## Usage

Encodes the given string and returns it in hexadecimal.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
