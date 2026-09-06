---
title: "security.HexBinary.FromBase64"
source: "fgl-topics/c_gws_SecurityHexBinary_FromBase64.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.FromBase64"
type: "concept"
---

# security.HexBinary.FromBase64

> Converts a base64 string to the hexadecimal equivalent.

## Syntax

```
security.HexBinary.FromBase64(
   val64 STRING )
  RETURNS STRING
```

1. val64 defines a string encoded in base64.

## Usage

This method decodes the given base64 string and returns it in its hexadecimal form.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
