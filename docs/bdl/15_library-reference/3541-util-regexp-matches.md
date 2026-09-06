---
title: "util.Regexp.matches"
source: "fgl-topics/c_fgl_ext_util_Regexp_matches.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.matches"
type: "concept"
---

# util.Regexp.matches

> Tests if a string contains any match of this regular expression.

## Syntax

```
matches(
  subject STRING
 )
  RETURNS BOOLEAN
```

1. subject is the subject to be scanned.

## Usage

The `matches()` method returns `TRUE`, if the string passed as
parameter matches the current regular expression.

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    LET re = util.Regexp.compile(`^[ABC]XX$`)
    DISPLAY re.matches("ZZZ")
    DISPLAY re.matches("ZAXXZ")
    DISPLAY re.matches("AXX")
END MAIN
```

Output:

```
     0
     0
     1
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[STRING.matches](2925-string-matches.md "Tests if the string matches a regular expression.")
