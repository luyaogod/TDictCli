---
title: "util.Strings.base64DecodeToString"
source: "fgl-topics/c_fgl_ext_util_Strings_base64DecodeToString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.base64DecodeToString"
type: "concept"
---

# util.Strings.base64DecodeToString

> Decodes a base64 encoded string and returns the corresponding string.

## Syntax

```
util.Strings.base64DecodeToString(
  base64 STRING
 )
  RETURNS STRING
```

1. base64 is the Base64 encoded string.

## Usage

The `util.Strings.base64DecodeToString()` method converts the Base64 encoded
string passed as parameter to an array of bytes, then it converts the byte array to a string
representation in the current locale, and returns that string.

If the Base64 source string contains a sequence of bytes that does not represent a valid
character in the [current application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."), the
function returns `NULL`.

> **Note:**
>
> In contrast to [`util.Strings.urlDecode()`](3559-util-strings-urldecode.md "Converts the URL-encoded string to a string in the current application locale."),
> the original string is not converted from UTF-8 to the application character encoding: The Base64
> source string must represent valid characters in the current application locale.

## Example

(Character encoding is UTF-8)

```
IMPORT util
MAIN
    DEFINE base64 STRING
    LET base64 = util.Strings.base64EncodeFromString( "Forêt" )
    DISPLAY base64
    DISPLAY util.Strings.base64DecodeToString( base64 )
END MAIN
```

Output:

```
Rm9yw6p0
Forêt
```

## Related links

**Related concepts**  

[util.Strings.base64EncodeFromString](3556-util-strings-base64encodefromstring.md "Converts the string passed as parameter to a Base64 encoded string.")
