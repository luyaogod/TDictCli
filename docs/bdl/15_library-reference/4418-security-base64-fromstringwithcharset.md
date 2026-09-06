---
title: "security.Base64.FromStringWithCharset"
source: "fgl-topics/c_gws_SecurityBase64_FromStringWithCharset.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.FromStringWithCharset"
type: "concept"
---

# security.Base64.FromStringWithCharset

> Encodes the given string in base64, based on a given charset.

## Syntax

```
security.Base64.FromStringWithCharset(
   clearVal STRING, 
   charset STRING )
  RETURNS STRING
```

1. clearVal defines the string to be
   encoded.
2. charset defines the
   character set to be used.

## Usage

This method encodes the string, clearVal, based on the specified charset,
charset, and returns it in base64.

Before conversion, the string is converted from the local DVM charset to the specified
encoding.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
