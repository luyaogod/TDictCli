---
title: "util.Regexp.getMatchAll"
source: "fgl-topics/c_fgl_ext_util_Regexp_getMatchAll.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.getMatchAll"
type: "concept"
---

# util.Regexp.getMatchAll

> Returns an array with all matches of the regular expression in the subject.

## Syntax

```
getMatchAll(
  subject STRING
 )
  RETURNS DYNAMIC ARRAY OF STRING
```

1. subject is the string to be scanned.

## Usage

The `getMatchAll()` method returns all occurrences of the strings matching the
regular expression, in the string passed as parameter.

The method returns an empty array, if the passed string does not contain any
matching string for the current regular expression.

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    DEFINE arr DYNAMIC ARRAY OF STRING
    LET re = util.Regexp.compile(`[ABC]XX`)
    LET arr = re.getMatchAll("ZBXXZAXXZCXX")
    DISPLAY arr[1]
    DISPLAY arr[2]
    DISPLAY arr[3]
END MAIN
```

Output:

```
BXX
AXX
CXX
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.getMatch](3533-util-regexp-getmatch.md "Returns the text of the first (leftmost) match of the regular expression in the subject.")
