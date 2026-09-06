---
title: "util.Regexp.getMatch"
source: "fgl-topics/c_fgl_ext_util_Regexp_getMatch.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.getMatch"
type: "concept"
---

# util.Regexp.getMatch

> Returns the text of the first (leftmost) match of the regular expression in the subject.

## Syntax

```
getMatch(
  subject STRING
 )
  RETURNS STRING
```

1. subject is the string to be scanned.

## Usage

The `getMatch()` method returns the first occurrence of the string matching the
regular expression, in the string passed as parameter.

The method returns `NULL`, if the passed string does not contain any matching
string for the current regular expression.

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    LET re = util.Regexp.compile(`[ABC]XX`)
    DISPLAY re.getMatch("ZZZZBXXZZZ")
END MAIN
```

Output:

```
BXX
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.getMatchAll](3534-util-regexp-getmatchall.md "Returns an array with all matches of the regular expression in the subject.")
