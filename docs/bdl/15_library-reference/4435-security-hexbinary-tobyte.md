---
title: "security.HexBinary.ToByte"
source: "fgl-topics/c_gws_SecurityHexBinary_ToByte.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.ToByte"
type: "concept"
---

# security.HexBinary.ToByte

> Decodes an hexadecimal string into a BYTE variable.

## Syntax

```
security.HexBinary.ToByte(
   hex STRING,
   ret BYTE )
```

1. hex defines a string in hexadecimal.
2. ret defines the input parameter to fill with data of type BYTE.

## Usage

Decodes the given hexadecimal string and fills the BYTE variable with binary data.

> **Important:**
>
> The `BYTE` must be located in
> memory.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
