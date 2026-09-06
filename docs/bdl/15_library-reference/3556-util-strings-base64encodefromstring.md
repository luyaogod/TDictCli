---
title: "util.Strings.base64EncodeFromString"
source: "fgl-topics/c_fgl_ext_util_Strings_base64EncodeFromString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.base64EncodeFromString"
type: "concept"
---

# util.Strings.base64EncodeFromString

> Converts the string passed as parameter to a Base64 encoded string.

## Syntax

```
util.Strings.base64EncodeFromString(
  s STRING
 )
  RETURNS STRING
```

1. s is the source string to convert in Base64.

## Usage

The `util.Strings.base64EncodeFromString()` method first converts the string
passed as parameter (in the current character encoding of the application) to an array of bytes,
then it converts the array of bytes to a Base64 representation, and returns the resulting Base64
encoded string.

> **Note:**
>
> In contrast to [`util.Strings.urlEncode()`](3560-util-strings-urlencode.md "Converts a string from the current codeset to a URL-encoded string."),
> the original string is not converted from the application locale to UTF-8, before performing the
> encoding to Base64: The resulting Base64 encoded string will contain byte sequences representing
> characters in the current application locale.

## Example

(Character encoding is UTF-8)

```
IMPORT util
MAIN
    DISPLAY util.Strings.base64EncodeFromString( "Forêt" )
END MAIN
```

Output:

```
Rm9yw6p0
```

## Related links

**Related concepts**  

[util.Strings.base64DecodeToString](3553-util-strings-base64decodetostring.md "Decodes a base64 encoded string and returns the corresponding string.")
