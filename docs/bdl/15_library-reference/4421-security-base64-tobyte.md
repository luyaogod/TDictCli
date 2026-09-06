---
title: "security.Base64.ToByte"
source: "fgl-topics/c_gws_SecurityBase64_ToByte.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.ToByte"
type: "concept"
---

# security.Base64.ToByte

> Decodes the given base64 string into a BYTE.

## Syntax

```
security.Base64.ToByte(
   val64 STRING,
   ret BYTE )
```

1. val64 defines the string in base64.
2. ret defines a parameter of type BYTE to fill with data.

## Usage

Decodes the base64 string specified in val64 and fills the BYTE variable with
binary data.

> **Important:**
>
> The `BYTE` must be located in
> memory.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
