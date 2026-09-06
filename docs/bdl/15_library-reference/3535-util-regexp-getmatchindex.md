---
title: "util.Regexp.getMatchIndex"
source: "fgl-topics/c_fgl_ext_util_Regexp_getMatchIndex.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.getMatchIndex"
type: "concept"
---

# util.Regexp.getMatchIndex

> Returns the starting and ending position of the first (leftmost) match of the regular expression in the subject.

## Syntax

```
getMatchIndex(
  subject STRING
 )
  RETURNS (INTEGER, INTEGER)
```

1. subject is the subject to be scanned.

## Usage

The `getMatchIndex()` method returns the position of the first occurrence of the
string matching the regular expression, in the string passed as parameter.

The method returns (`NULL,NULL`), if the passed string does not contain any
matching string for the current regular expression.

The positions are expressed in byte units or character units, depending on the
[current length semantics (FGL\_LENGTH\_SEMANTICS)](../09_advanced-features/0881-length-semantics-settings.md).

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    DEFINE s,e INTEGER
    DEFINE arr DYNAMIC ARRAY OF STRING
    LET re = util.Regexp.compile(`[ABC]XX`)
    CALL re.getMatchIndex("ZZZZBXXZZ") RETURNING s,e
    DISPLAY s, e
END MAIN
```

Output:

```
          5          7
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.getMatchIndexAll](3536-util-regexp-getmatchindexall.md "Returns an array with indexes to all matches of the regular expression in the subject.")
