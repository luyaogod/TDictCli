---
title: "security.HexBinary.FromByte"
source: "fgl-topics/c_gws_SecurityHexBinary_FromByte.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.FromByte"
type: "concept"
---

# security.HexBinary.FromByte

> Encodes BYTE data in hexadecimal.

## Syntax

```
security.HexBinary.FromByte(
   data BYTE )
  RETURNS STRING
```

1. data defines the data of type BYTE to be
   encoded.

## Usage

This method encodes the given BYTE data in hexadecimal and returns the string.

> **Important:**
>
> The `BYTE` must be located in
> memory.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
