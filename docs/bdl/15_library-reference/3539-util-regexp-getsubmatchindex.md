---
title: "util.Regexp.getSubmatchIndex"
source: "fgl-topics/c_fgl_ext_util_Regexp_getSubmatchIndex.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.getSubmatchIndex"
type: "concept"
---

# util.Regexp.getSubmatchIndex

> Returns an array with the indexes of the sub-matches of the first match of regular expression in the subject.

## Syntax

```
getSubmatchIndex(
  subject STRING
 )
  RETURNS DYNAMIC ARRAY OF INTEGER
```

1. subject is the subject to be scanned.

## Usage

The `getSubmatchIndex()` method scans the string passed as parameter to find the
first matching string and all subexpressions matching strings inside the main matching string, and
returns the position of these matches.

The method returns an empty array, if the passed string does not contain any
matching string for the current regular expression.

In the returned array, elements at index 1 and 2 are the start and end positions of the whole
match. The elements at index 3 and 4 are the start and end positions of the first subexpression. The
elements at index 5 and 6 are the start and end positions of the second subexpression, etc.

The positions are expressed in byte units or character units, depending on the
[current length semantics (FGL\_LENGTH\_SEMANTICS)](../09_advanced-features/0881-length-semantics-settings.md).

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    DEFINE arr DYNAMIC ARRAY OF INTEGER
    LET re = util.Regexp.compile(`([A-Z]*)([0-9]*)`)
    LET arr = re.getSubmatchIndex("ABC678EFG99")
    DISPLAY arr[1], arr[2]
    DISPLAY arr[3], arr[4]
    DISPLAY arr[5], arr[6]
END MAIN
```

Output:

```
          1          6
          1          3
          4          6
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.getSubmatchIndexAll](3540-util-regexp-getsubmatchindexall.md "Returns a 2-dimensional array with the indexes of the sub-matches of all matches of regular expression in the subject.")
