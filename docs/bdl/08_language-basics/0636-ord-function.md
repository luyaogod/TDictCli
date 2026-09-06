---
title: "ORD() [function]"
source: "fgl-topics/c_fgl_operators_ORD.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > ORD() [function]"
type: "concept"
---

# ORD() [function]

> The ORD() operator returns the code point of a character in the current locale.

## Syntax

```
ORD ( source STRING )
```

1. source is a string expression. The processing of the character byte sequence
   depends depends on the current application locale and length semantics.

## Usage

The value returned by `ORD()` is the code point in the current locale of the
character passed as argument, or the first byte of the encoded character.

The `ORD()` operator reads only the first character (or byte) of the source string
passed as argument.

The `ORD()` operator returns `NULL`, if the first character (or
byte) of the source string is not valid in the current encoding.

The processing of the character byte sequence depends on the [application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.") and the [length semantics](../09_advanced-features/0881-length-semantics-settings.md) setting:

- With a single byte encoding like ISO-8859-15 (and byte length semantics), `ORD()`
  returns the code point of the first SBCS byte / character, in the current charset.
- With UTF-8 encoding and char length semantics, `ORD()` returns the UNICODE code
  point of the first character.
- With UTF-8 encoding and byte length semantics, `ORD()` returns the value of the
  first byte of the first character.

If the character encoding is single byte, or when using UTF-8 and char length semantics,
`ORD()` is the inverse of [`ASCII()`](0630-ascii.md "The ASCII operator produces a character from its ordinal value.").

The next example shows the behavior of `ORD()`, depending on the character set and
length semantics, on a Linux platform:

```
MAIN
    DEFINE src STRING
    DISPLAY "LC_ALL              : ", fgl_getenv("LC_ALL")
    DISPLAY "FGL_LENGTH_SEMANTICS: ", fgl_getenv("FGL_LENGTH_SEMANTICS")
    LET src = arg_val(1)
    DISPLAY "src                 : ",src
    DISPLAY "ORD(s)              : ",ORD(src)
END MAIN
```

ISO-8859-15 + BYTE length semantics (with terminal using ISO-8859-15
encoding):

```
$ export LC_ALL="en_US.iso8859-15"
$ export FGL_LENGTH_SEMANTICS="BYTE"
$ fglrun ord.42m "ôxx"
LC_ALL              : en_US.iso8859-15
FGL_LENGTH_SEMANTICS: BYTE
src                 : ôxx
ORD(s)              :         244        <-- ISO-8859-15 code point (0xF4) of ô
```

UTF-8 + CHAR length semantics (with terminal using UTF-8
encoding):

```
$ export LC_ALL="en_US.utf-8"
$ export FGL_LENGTH_SEMANTICS="CHAR"
$ fglrun ord.42m "愛xx"
LC_ALL              : en_US.utf-8
FGL_LENGTH_SEMANTICS: CHAR
src                 : 愛xx
ORD(s)              :       24859        <-- UNICODE code point of 愛
```

UTF-8 + BYTE length semantics (with terminal using UTF-8
encoding):

```
$ export LC_ALL="en_US.utf-8"
$ export FGL_LENGTH_SEMANTICS="BYTE"
$ fglrun ord.42m "愛xx"
LC_ALL              : en_US.utf-8
FGL_LENGTH_SEMANTICS: BYTE
src                 : 愛xx
ORD(s)              :         230        <-- first UTF-8 byte (0xE6) of 愛
```

## Related links

**Related concepts**  

[Defining the application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.")

[fgl\_keyval()](../15_library-reference/2771-fgl-keyval.md "Returns the key code of a logical or physical key.")

[ASCII](0630-ascii.md "The ASCII operator produces a character from its ordinal value.")
