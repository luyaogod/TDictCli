---
title: "util.Regexp.getSubmatchAll"
source: "fgl-topics/c_fgl_ext_util_Regexp_getSubmatchAll.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.getSubmatchAll"
type: "concept"
---

# util.Regexp.getSubmatchAll

> Returns 2-dimensional array with the sub-matches of all matches of regular expression in the subject.

## Syntax

```
getSubmatchAll(
  subject STRING
 )
  RETURNS DYNAMIC ARRAY WITH DIMENSION 2 OF STRING
```

1. subject is the subject to be scanned.

## Usage

The `getSubmatchAll()` method scans the string passed as parameter to find all
matching strings and all subexpressions matching strings inside the main matching string.

Submatches are matches of parenthesized subexpressions (capturing groups). For
example, the regular expression "`([A-Z]*)([0-9]*)`" defines two subexpressions for a
first group composed of capital letters A to Z and a second group is composed of digits.

In the returned array, element at index `[1,1]` represents the first whole match.
The element at index `[1,2]` represents the value of the first subexpression for the
first whole match. The element at index `[1,3]` represents the value of the second
subexpression first whole match. The element at index [1,n+1] is the last subexpression match for
the first whole match, where n is the number of sub-matches. The element at index
`[2,1]` represents the second whole match. The element at index
`[2,2]` represents the value of the first subexpression for the second whole match,
etc.

The method returns an empty array, if the passed string does not contain any
matching string for the current regular expression.

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    DEFINE arr DYNAMIC ARRAY WITH DIMENSION 2 OF STRING
    LET re = util.Regexp.compile(`([A-Z]*)([0-9]*)`)
    LET arr = re.getSubmatchAll("ABC678EFG99")
    DISPLAY arr[1,1]
    DISPLAY arr[1,2]
    DISPLAY arr[1,3]
    DISPLAY arr[2,1]
    DISPLAY arr[2,2]
    DISPLAY arr[2,3]
END MAIN
```

Output:

```
ABC678
ABC
678
EFG99
EFG
99
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.getSubmatch](3537-util-regexp-getsubmatch.md "Returns an array with the sub-matches of the first match of regular expression in the subject.")
