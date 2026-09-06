---
title: "security.Base64.FromByte"
source: "fgl-topics/c_gws_SecurityBase64_FromByte.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.FromByte"
type: "concept"
---

# security.Base64.FromByte

> Encodes the given BYTE data in base64.

## Syntax

```
security.Base64.FromByte(
   data BYTE )
  RETURNS STRING
```

1. data defines the data of type BYTE to be encoded.

## Usage

This method encodes the given BYTE data in base64 and returns the string.

> **Important:**
>
> The `BYTE` must be located in
> memory.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
