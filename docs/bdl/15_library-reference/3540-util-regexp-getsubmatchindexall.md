---
title: "util.Regexp.getSubmatchIndexAll"
source: "fgl-topics/c_fgl_ext_util_Regexp_getSubmatchIndexAll.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.getSubmatchIndexAll"
type: "concept"
---

# util.Regexp.getSubmatchIndexAll

> Returns a 2-dimensional array with the indexes of the sub-matches of all matches of regular expression in the subject.

## Syntax

```
getSubmatchIndexAll(
  subject STRING
 )
  RETURNS DYNAMIC ARRAY WITH DIMENSION 2 OF INTEGER
```

1. subject is the subject to be scanned.

## Usage

The `getSubmatchIndexAll()` method scans the string passed as parameter to find
all matching strings and all subexpressions matching strings inside the main matching string, and
returns the position of these matches.

The method returns an empty array, if the passed string does not contain any
matching string for the current regular expression.

In the returned array, elements at index `[1,1]` and `[1,2]` are
the start and end positions of the first whole match. The elements at index `[1,3]`
and `[1,4]` are the start and end positions of the first subexpression for the first
whole match. The elements at index `[1,5]` and `[1,6]` are the start
and end positions of the second subexpression for the first whole match. Elements at index
`[2,1]` and `[2,2]` are the start and end positions of the second
whole match. The elements at index `[2,3]` and `[2,4]` are the start
and the end position of the first subexpression for the second whole match, etc.

The positions are expressed in byte units or character units, depending on the
[current length semantics (FGL\_LENGTH\_SEMANTICS)](../09_advanced-features/0881-length-semantics-settings.md).

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    DEFINE arr DYNAMIC ARRAY WITH DIMENSION 2 OF INTEGER
    LET re = util.Regexp.compile(`([A-Z]*)([0-9]*)`)
    LET arr = re.getSubmatchIndexAll("ABC678EFG99")
    DISPLAY "First whole match:"
    DISPLAY arr[1,1], arr[1,2]
    DISPLAY arr[1,3], arr[1,4]
    DISPLAY arr[1,5], arr[1,6]
    DISPLAY "Second whole match:"
    DISPLAY arr[2,1], arr[2,2]
    DISPLAY arr[2,3], arr[2,4]
    DISPLAY arr[2,5], arr[2,6]
END MAIN
```

Output:

```
First whole match:
          1          6
          1          3
          4          6
Second whole match:
          7         11
          7          9
         10         11
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.getSubmatchIndex](3539-util-regexp-getsubmatchindex.md "Returns an array with the indexes of the sub-matches of the first match of regular expression in the subject.")
