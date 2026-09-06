---
title: "security.Base64.Xor"
source: "fgl-topics/c_gws_SecurityBase64_Xor.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.Xor"
type: "concept"
---

# security.Base64.Xor

> Computes the exclusive disjunction between two base64 encoded strings.

## Syntax

```
security.Base64.Xor(
   clearVal1 STRING,
   clearVal2 STRING )
  RETURNS STRING
```

1. clearVal1 defines the first string encoded in base64.
2. clearVal2 defines the second string encoded in
   base64.

## Usage

Decodes the two given strings and does an exclusive disjunction between the two binary inputs.
The result is returned encoded in base64.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
