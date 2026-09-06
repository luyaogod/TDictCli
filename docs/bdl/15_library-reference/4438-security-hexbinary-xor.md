---
title: "security.HexBinary.Xor"
source: "fgl-topics/c_gws_SecurityHexBinary_Xor.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.Xor"
type: "concept"
---

# security.HexBinary.Xor

> Computes the exclusive disjunction between two hexadecimal encoded strings.

## Syntax

```
security.HexBinary.Xor(
   hexVal1 STRING,
   hexVal2 STRING )
  RETURNS STRING
```

1. hexVal1 defines the first string in hexadecimal.
2. hexVal2 defines the second string in hexadecimal.

## Usage

Decodes the two given strings and does an exclusive disjunction between the two binary inputs.
The result is returned in hexadecimal.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
