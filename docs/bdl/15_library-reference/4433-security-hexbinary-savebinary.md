---
title: "security.HexBinary.SaveBinary"
source: "fgl-topics/c_gws_SecurityHexBinary_SaveBinary.html"
breadcrumb: "Library reference > Extension packages > The security package > The HexBinary class > HexBinary methods > security.HexBinary.SaveBinary"
type: "concept"
---

# security.HexBinary.SaveBinary

> Decodes an hexadecimal strings and writes the binary data to a file.

## Syntax

```
security.HexBinary.SaveBinary(
   path STRING,
   hexBinData STRING )
```

1. path defines the path to the binary file.
2. hexBinData defines the hexadecimal string to be written.

## Usage

Decodes the given hexadecimal string and writes the binary data to the file defined by
path.

This method can be used to save data from a network on the disk.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
