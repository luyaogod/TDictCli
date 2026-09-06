---
title: "security.HexBinary.LoadBinary"
source: "fgl-topics/c_gws_SecurityHexBinary_LoadBinary.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.LoadBinary"
type: "concept"
---

# security.HexBinary.LoadBinary

> Reads binary data from a file and converts it to hexadecimal.

## Syntax

```
security.HexBinary.LoadBinary(
   path STRING )
  RETURNS STRING
```

1. path defines the path to the binary file.

## Usage

Reads the file located at path and returns this binary data in hexadecimal
format.

For example, this method can be used to send images through a network.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
