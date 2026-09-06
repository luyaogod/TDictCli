---
title: "security.Base64.ToString"
source: "fgl-topics/c_gws_SecurityBase64_ToString.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.ToString"
type: "concept"
---

# security.Base64.ToString

> Decodes the given base64 string.

## Syntax

```
security.Base64.ToString(
   val64 STRING )
  RETURNS STRING
```

1. val64 defines the string in
   base64.

## Usage

Decodes the given base64 string and returns it in its clear (human readable) form.

If the base64 string does not contain human readable data, the method will raise an
exception.

If the base64 string contains bytes sequences that do not match a valid character in the current
encoding, the method raises a conversion error.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
