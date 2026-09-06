---
title: "STRING.getMultibyteLength"
source: "fgl-topics/c_fgl_datatypes_STRING_getMultibyteLength.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.getMultibyteLength"
type: "concept"
---

# STRING.getMultibyteLength

> Counts the number of bytes in the string.

## Syntax

```
getMultibyteLength( )
       RETURNS INTEGER
```

## Usage

Use this method to count the number of bytes of the string, especially when using UTF-8 locale
and `CHAR` length semantics.

## Example

```
MAIN
    DEFINE s STRING
    LET s = "forêt"   -- the e with circumflex accent uses 2 bytes in UTF-8
    DISPLAY "getLength()          = ", s.getLength()
    DISPLAY "getMultibyteLength() = ", s.getMultibyteLength()
END MAIN
```

Output (Linux):

```
$ export LC_ALL=en_US.utf8

$ export FGL_LENGTH_SEMANTICS=BYTE
$ fglcomp main.4gl && fglrun main.42m
getLength()          =           6
getMultibyteLength() =           6

$ export FGL_LENGTH_SEMANTICS=CHAR
$ fglcomp main.4gl && fglrun main.42m
getLength()          =           5
getMultibyteLength() =           6
```

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")

[STRING.getLength](2923-string-getlength.md "Returns the length of the current string.")
