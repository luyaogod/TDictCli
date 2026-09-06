---
title: "util.Regexp.getMatchIndexAll"
source: "fgl-topics/c_fgl_ext_util_Regexp_getMatchIndexAll.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.getMatchIndexAll"
type: "concept"
---

# util.Regexp.getMatchIndexAll

> Returns an array with indexes to all matches of the regular expression in the subject.

## Syntax

```
getMatchIndexAll(
  subject STRING
 )
  RETURNS DYNAMIC ARRAY OF INTEGER
```

1. subject is the string to be scanned.

## Usage

The `getMatchIndexAll()` method returns all starting and ending positions of all
occurrences of the strings matching the regular expression, in the string passed as parameter.

The method returns an empty array, if the passed string does not contain any
matching string for the current regular expression.

In the returned array, elements at index 1 and 2 are the start and end positions of the first
match. The elements at index 3 and 4 are the start and end positions of the second match, etc.

The positions are expressed in byte units or character units, depending on the
[current length semantics (FGL\_LENGTH\_SEMANTICS)](../09_advanced-features/0881-length-semantics-settings.md).

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    DEFINE arr DYNAMIC ARRAY OF INTEGER
    LET re = util.Regexp.compile(`[ABC]XX`)
    LET arr = re.getMatchIndexAll("ZBXXZAXXZBXX")
    DISPLAY arr[1], arr[2]
    DISPLAY arr[3], arr[4]
    DISPLAY arr[5], arr[6]
END MAIN
```

Output:

```
          2          4
          6          8
         10         12
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.getMatchIndex](3535-util-regexp-getmatchindex.md "Returns the starting and ending position of the first (leftmost) match of the regular expression in the subject.")
