---
title: "security.Base64.FromString"
source: "fgl-topics/c_gws_SecurityBase64_FromString.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.FromString"
type: "concept"
---

# security.Base64.FromString

> Encodes the given string in base64.

## Syntax

```
security.Base64.FromString(
   clearVal STRING )
  RETURNS STRING
```

1. clearVal defines the string to be
   encoded.

## Usage

This method encodes the given string,clearVal, and returns it in base64.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
