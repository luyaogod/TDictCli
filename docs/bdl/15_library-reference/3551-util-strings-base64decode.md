---
title: "util.Strings.base64Decode"
source: "fgl-topics/c_fgl_ext_util_Strings_base64Decode.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.base64Decode"
type: "concept"
---

# util.Strings.base64Decode

> Decodes a Base64 encoded string and writes the bytes to a file.

## Syntax

```
util.Strings.base64Decode(
  base64 STRING,
  filename STRING
 )
```

1. base64 is the Base64 encoded string.
2. filename is the name of the file to write to.

## Usage

The `util.Strings.base64Decode()` method converts the Base64 encoded
string passed as first parameter, and writes the bytes to file specified as second
parameter.

## Example

```
IMPORT util
MAIN
    DEFINE base64 STRING
    LET base64 = util.Strings.base64Encode( "picture1.png" )
    DISPLAY base64
    CALL util.Strings.base64Decode( base64, "picture2.png" )
END MAIN
```

## Related links

**Related concepts**  

[util.Strings.base64Encode](3554-util-strings-base64encode.md "Converts the content of a file to a Base64 encoded string.")
