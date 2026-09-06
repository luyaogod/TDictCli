---
title: "util.Regexp.replaceAll"
source: "fgl-topics/c_fgl_ext_util_Regexp_replaceAll.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.replaceAll"
type: "concept"
---

# util.Regexp.replaceAll

> Substitutes all matches of the regular expression with the replacement string.

## Syntax

```
replaceAll(
  subject STRING,
  replacement STRING
 )
  RETURNS STRING
```

1. subject is the string to be scanned.
2. replacement is the replacement string.

## Usage

The `replaceAll()` method scans the string passed as parameter, replaces each
substring matching the current regular expression with the replacement string, and returns the new
resulting string.

The replacement string can reference captured groups with the
`$num` notation, where num is the ordinal
position of the group delimited by parentheses in the regular expression, and where
`$0` represents the whole matching string.

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    LET re = util.Regexp.compile(`[@#%]([0-9]+)`)
    DISPLAY re.replaceAll("ABC#33 DE@891 FG#45","-$1")
END MAIN
```

Output:

```
ABC-33 DE-891 FG-45
```

## Related links

**Related concepts**  

[util.Regexp.compile](3532-util-regexp-compile.md "Compiles a regular expression and returns a util.Regexp object.")

[util.Regexp.replaceFirst](3543-util-regexp-replacefirst.md "Substitutes the match of the regular expression with the replacement string.")

[STRING.replaceAll](2926-string-replaceall.md "Replace all substrings matching a regular expression.")
