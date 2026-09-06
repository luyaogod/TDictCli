---
title: "STRING.matches"
source: "fgl-topics/c_fgl_datatypes_STRING_matches.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.matches"
type: "concept"
---

# STRING.matches

> Tests if the string matches a regular expression.

## Syntax

```
matches(
     regex STRING )
   RETURNS BOOLEAN
```

1. regex is the regular expression. See [Regular expression patterns](3545-regular-expression-patterns.md "Summary of util.Regexp.compile() pattern syntax.").

## Usage

This method scans the current string and returns `TRUE`, when the string matches
the regular expression passed as parameter.

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
    LET s = "id_243"
    DISPLAY s.matches(`^[[:alpha:]_]+[[:alnum:]_]*$`)
END MAIN
```

Output:

```
     1
```

## Related links

**Related concepts**  

[The util.Regexp class](3530-the-util-regexp-class.md "The util.Regexp class provides character string pattern matching and replacements based on regular expressions.")
