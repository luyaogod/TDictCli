---
title: "security.Base64.FromHexBinary"
source: "fgl-topics/c_gws_SecurityBase64_FromHexBinary.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.FromHexBinary"
type: "concept"
---

# security.Base64.FromHexBinary

> Decodes the given hexadecimal string to base64.

## Syntax

```
security.Base64.FromHexBinary(
   hexBinVal STRING )
  RETURNS STRING
```

1. hexBinVal defines the string in hexadecimal form.

## Usage

This method decodes the given hexadecimal string and returns it in base64.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
