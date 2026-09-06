---
title: "util.Strings.base64EncodeFromHexString"
source: "fgl-topics/c_fgl_ext_util_Strings_base64EncodeFromHexString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.base64EncodeFromHexString"
type: "concept"
---

# util.Strings.base64EncodeFromHexString

> Converts the hexadecimal string passed as parameter to a Base64 encoded string.

## Syntax

```
util.Strings.base64EncodeFromHexString(
  s STRING
 )
  RETURNS STRING
```

1. s is the source hexadecimal string to convert in Base64.

## Usage

The `util.Strings.base64EncodeFromHexString()` method first converts the
hexadecimal string passed as parameter to an array of bytes, then it converts the array of bytes to
a Base64 representation, and returns the resulting Base64 encoded string.

> **Tip:**
>
> Use the [`util.Integer.toHexString()`](3524-util-integer-tohexstring.md "Returns the string representation of an INTEGER, as an unsigned integer in base 16.")
> method to build an hexadecimal string from a set of integer/byte values, then convert the hexadecimal
> string to a Base64 representation.

## Example

```
IMPORT util
MAIN
    DEFINE hexa VARCHAR(50)
    DISPLAY util.Strings.base64EncodeFromHexString( "00" )
    DISPLAY util.Strings.base64EncodeFromHexString( "0000" )
    DISPLAY util.Strings.base64EncodeFromHexString( "FFFF" )
    DISPLAY util.Strings.base64EncodeFromHexString( "42EFB3E5" )
    LET hexa = util.Integer.toHexString(65)
            || util.Integer.toHexString(66)
            || util.Integer.toHexString(67)
    DISPLAY hexa
    DISPLAY util.Strings.base64EncodeFromHexString( hexa )
END MAIN
```

Output:

```
AA==
AAA=
//8=
Qu+z5Q==
414243
QUJD
```

## Related links

**Related concepts**  

[util.Strings.base64DecodeToHexString](3552-util-strings-base64decodetohexstring.md "Decodes a base64 encoded string and returns the corresponding hexadecimal string.")
