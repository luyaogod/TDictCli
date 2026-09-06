---
title: "STRING.replaceAll"
source: "fgl-topics/c_fgl_datatypes_STRING_replaceAll.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.replaceAll"
type: "concept"
---

# STRING.replaceAll

> Replace all substrings matching a regular expression.

## Syntax

```
replaceAll(
     regex STRING,
     replacement STRING )
   RETURNS STRING
```

1. regex is the regular expression to find substrings that must be replaced. See
   [Regular expression patterns](3545-regular-expression-patterns.md "Summary of util.Regexp.compile() pattern syntax.").
2. replacement is the replacement string.

## Usage

This method scans the current string value for substrings matching the regular expression passed
as first parameter, replaces all these matching substrings by the replacement string, and returns
the new resulting string.

The replacement string can reference captured groups with the
`$num` notation, where num is the ordinal
position of the group delimited by parentheses in the regular expression, and where
`$0` represents the whole matching string.

> **Note:**
>
> When using single or double quoted string literals, backslash is
> interpreted as escape character. Consider using [back
> quotes](../08_language-basics/0588-text-literals.md "Text literals define a character string in an expression.") as delimiters for regular expressions strings, and use backslash characters directly
> as required by the regexp syntax ( `` `\b` `` instead of `"\\b"` or
> `'\\b'` )

## Example

```
MAIN
    DEFINE s STRING
    LET s = "ABC123DEF456"
    DISPLAY s.replaceAll(`[0-9]+`,"X")
    DISPLAY s.replaceAll(`[0-9]+`,"-$0-")
END MAIN
```

Output:

```
ABCXDEFX
ABC-123-DEF-456-
```

## Related links

**Related concepts**  

[util.Regexp.replaceAll](3542-util-regexp-replaceall.md "Substitutes all matches of the regular expression with the replacement string.")
