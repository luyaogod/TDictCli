---
title: "util.Regexp.getSubmatch"
source: "fgl-topics/c_fgl_ext_util_Regexp_getSubmatch.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.getSubmatch"
type: "concept"
---

# util.Regexp.getSubmatch

> Returns an array with the sub-matches of the first match of regular expression in the subject.

## Syntax

```
getSubmatch(
  subject STRING
 )
  RETURNS DYNAMIC ARRAY OF STRING
```

1. subject is the subject to be scanned.

## Usage

The `getSubmatch()` method scans the string passed as parameter to find the first
matching string and all subexpressions matching strings inside the main matching string.

Submatches are matches of parenthesized subexpressions (capturing groups). For
example, the regular expression "`([A-Z]*)([0-9]*)`" defines two subexpressions for a
first group composed of capital letters A to Z and a second group is composed of digits.

In the returned array, element at index 1 represents the whole match. The element at index 2
represents the value of the first subexpression. The element at index 3 represents the value of the
second subexpression, etc.

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    DEFINE arr DYNAMIC ARRAY OF STRING
    LET re = util.Regexp.compile(`([A-Z]*)([0-9]*)([A-Z]*)`)
    LET arr = re.getSubmatch("ABC678EFG99")
    DISPLAY arr[1]
    DISPLAY arr[2]
    DISPLAY arr[3]
    DISPLAY arr[4]
END MAIN
```

Output:

```
ABC678EFG
ABC
678
EFG
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.getSubmatchAll](3538-util-regexp-getsubmatchall.md "Returns 2-dimensional array with the sub-matches of all matches of regular expression in the subject.")
