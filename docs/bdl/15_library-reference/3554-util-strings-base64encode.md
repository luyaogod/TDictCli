---
title: "util.Strings.base64Encode"
source: "fgl-topics/c_fgl_ext_util_Strings_base64Encode.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.base64Encode"
type: "concept"
---

# util.Strings.base64Encode

> Converts the content of a file to a Base64 encoded string.

## Syntax

```
util.Strings.base64Encode(
  filename STRING
 )
  RETURNS STRING
```

1. filename is the name of the file to read from.

## Usage

The `util.Strings.base64Encode()` method reads the content of the file
passed as parameter, and converts the bytes to a Base64 encoded string.

## Example

```
IMPORT util
MAIN
    DISPLAY util.Strings.base64Encode( "picture.png" )
END MAIN
```

## Related links

**Related concepts**  

[util.Strings.base64Decode](3551-util-strings-base64decode.md "Decodes a Base64 encoded string and writes the bytes to a file.")
