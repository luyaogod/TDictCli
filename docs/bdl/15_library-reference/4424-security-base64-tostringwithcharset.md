---
title: "security.Base64.ToStringWithCharset"
source: "fgl-topics/c_gws_SecurityBase64_ToStringWithCharset.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.ToStringWithCharset"
type: "concept"
---

# security.Base64.ToStringWithCharset

> Decodes the given base64 string, based on a given charset.

## Syntax

```
security.Base64.ToStringWithCharset(
   val64 STRING,
   charset STRING )
  RETURNS STRING
```

1. val64 defines the string in
   base64.
2. charset defines the character set to be used.

## Usage

Decodes the given base64 string and returns it in its clear human readable form, based on a given
charset.

The original base64 encoded string is first decoded to a string that will be converted from the
specified charset to the local DVM charset. In case of charset conversion error, the error -15700 is
raised.

If the base64 string does not contain human readable data, the method will raise an
exception.

If the base64 string contains bytes sequences that do not match a valid character in the current
encoding, the method raises a conversion error.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
