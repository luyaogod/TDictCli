---
title: "security.Base64.LoadBinary"
source: "fgl-topics/c_gws_SecurityBase64_LoadBinary.html"
breadcrumb: "Library reference > Extension packages > The security package > The Base64 class > Base64 methods > security.Base64.LoadBinary"
type: "concept"
---

# security.Base64.LoadBinary

> Reads data from a file and encodes to base64.

## Syntax

```
security.Base64.LoadBinary(
   path STRING )
  RETURNS STRING
```

1. path defines the path to the binary file.

## Usage

Reads the file located at path and encodes this binary data in Base64
format.

For example, this method can be used to send images through a network.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
