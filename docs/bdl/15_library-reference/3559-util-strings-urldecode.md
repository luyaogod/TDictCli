---
title: "util.Strings.urlDecode"
source: "fgl-topics/c_fgl_ext_util_Strings_urlDecode.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Strings class > util.Strings methods > util.Strings.urlDecode"
type: "concept"
---

# util.Strings.urlDecode

> Converts the URL-encoded string to a string in the current application locale.

## Syntax

```
util.Strings.urlDecode(
  s STRING
 )
  RETURNS STRING
```

1. s is the URL-encoded source string (UTF-8 bytes).

## Usage

The `util.Strings.urlDecode()` method converts the URL-encoded string
passed as parameter to a character string.

The source string must contain ASCII characters and/or `%xx` hexadecimal
representation of UTF-8 encoding bytes.

The decoder is error tolerant:

- Alphabetical characters of a `%xx` element can be uppercase
  or lowercase (`%b2` = `%B2`).
- If the source string contains a set of `%xx` elements that
  represent a UTF-8 encoded character which is not existing in the current
  application locale, it will be converted to a `?` question
  mark.
- If the percent character is not followed by two hexadecimal digits, then a
  "`%`" is copied to the result string and the decoder
  continues at the next character.

## Example

```
IMPORT util
MAIN
    DISPLAY util.Strings.urlDecode("abc%C3%84%E2%82%AC")
END MAIN
```

Output:

```
abcÄ€
```

## Related links

**Related concepts**  

[Defining the application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.")

[util.Strings.urlEncode](3560-util-strings-urlencode.md "Converts a string from the current codeset to a URL-encoded string.")
