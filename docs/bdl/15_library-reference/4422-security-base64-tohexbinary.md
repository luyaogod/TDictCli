---
title: "security.Base64.ToHexBinary"
source: "fgl-topics/c_gws_SecurityBase64_ToHexBinary.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.ToHexBinary"
type: "concept"
---

# security.Base64.ToHexBinary

> Decodes the given base64 string to hexadecimal.

## Syntax

```
security.Base64.ToHexBinary(
   val64 STRING )
  RETURNS STRING
```

1. val64 defines a string encoded in base64.

## Usage

Decodes the base64 string, val64, and returns it in its hexadecimal form.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
