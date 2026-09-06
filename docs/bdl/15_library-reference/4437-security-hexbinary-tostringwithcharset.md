---
title: "security.HexBinary.ToStringWithCharset"
source: "fgl-topics/c_gws_SecurityHexBinary_ToStringWithCharset.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.ToStringWithCharset"
type: "concept"
---

# security.HexBinary.ToStringWithCharset

> Decodes an hexadecimal string to a clear, human-readable string, based on a given charset.

## Syntax

```
security.HexBinary.ToStringWithCharset(
   hexVal STRING,
   charset STRING )
  RETURNS STRING
```

1. hexVal defines the string in hexadecimal.
2. charset defines the character set to be
   used.

## Usage

Decodes the given hexadecimal string and returns it in its clear human readable form, based on a
given charset.

The original hexadecimal encoded string is first decoded to a string that will then be converted
from the specified charset to the local DVM charset. In case of charset conversion error, the error
-15700 is raised.

If the hexadecimal string does not contain a human readable string, the method will raise an
exception.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
