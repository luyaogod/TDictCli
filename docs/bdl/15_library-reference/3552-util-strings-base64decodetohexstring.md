---
title: "util.Strings.base64DecodeToHexString"
source: "fgl-topics/c_fgl_ext_util_Strings_base64DecodeToHexString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.base64DecodeToHexString"
type: "concept"
---

# util.Strings.base64DecodeToHexString

> Decodes a base64 encoded string and returns the corresponding hexadecimal string.

## Syntax

```
util.Strings.base64DecodeToHexString(
  base64 STRING
 )
  RETURNS STRING
```

1. base64 is the Base64 encoded string.

## Usage

The `util.Strings.base64DecodeToHexString()` method converts the Base64 encoded
string passed as parameter to an array of bytes, then it converts the byte array to the hexadecimal
representation of this array of bytes, and returns that string.

> **Tip:**
>
> After converting the Base64 string to an hexadecimal string, use the [`util.parseHexString()`](3518-util-integer-parsehexstring.md "Returns an INTEGER from its hexadecimal (base 16) string representation.")
> method to convert bytes of the hexadecimal string to integers.

## Example

```
IMPORT util
MAIN
    DEFINE hexa VARCHAR(50)
    DISPLAY util.Strings.base64DecodeToHexString( "AA==" )
    DISPLAY util.Strings.base64DecodeToHexString( "AAAAAA==" )
    DISPLAY util.Strings.base64DecodeToHexString( "AQ==" )
    DISPLAY util.Strings.base64DecodeToHexString( "//8=" )
    DISPLAY util.Strings.base64DecodeToHexString( "QUJDRA==" )
    LET hexa = util.Strings.base64DecodeToHexString( "QUJDRA==" )
    DISPLAY util.Integer.parseHexString( hexa[1,2] )
    DISPLAY util.Integer.parseHexString( hexa[3,4] )
    DISPLAY util.Integer.parseHexString( hexa[5,6] )
    DISPLAY util.Integer.parseHexString( hexa[7,8] )
END MAIN
```

Output:

```
00
00000000
01
ffff
41424344
         65
         66
         67
         68
```

## Related links

**Related concepts**  

[util.Strings.base64EncodeFromHexString](3555-util-strings-base64encodefromhexstring.md "Converts the hexadecimal string passed as parameter to a Base64 encoded string.")
