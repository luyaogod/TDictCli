---
title: "util.Strings.urlEncode"
source: "fgl-topics/c_fgl_ext_util_Strings_urlEncode.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.urlEncode"
type: "concept"
---

# util.Strings.urlEncode

> Converts a string from the current codeset to a URL-encoded string.

## Syntax

```
util.Strings.urlEncode(
  source STRING
 )
  RETURNS STRING
```

1. source is the source string to url-encode.

## Usage

The `util.Strings.urlEncode()` method converts the character string passed as
parameter to a URL-encoded string.

All characters not matching `[-_.~a-zA-Z0-9]` are "percent encoded":
Percent-encoding involves converting those characters to UTF-8 and representing its
corresponding byte values by a percent sign ("%") and a pair of hexadecimal digits.

## Example

```
IMPORT util
MAIN
    DISPLAY util.Strings.urlEncode("abcÄ€")
END MAIN
```

Output:

```
abc%C3%84%E2%82%AC
```

## Related links

**Related concepts**  

[Defining the application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.")

[util.Strings.urlDecode](3559-util-strings-urldecode.md "Converts the URL-encoded string to a string in the current application locale.")
