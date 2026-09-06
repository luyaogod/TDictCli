---
title: "STRING.split"
source: "fgl-topics/c_fgl_datatypes_STRING_split.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.split"
type: "concept"
---

# STRING.split

> Splits the current string around matches of the given regular expression.

## Syntax

```
split(
     regex STRING )
   RETURNS DYNAMIC ARRAY OF STRING
```

1. regex is the regular expression defining token separators. See [Regular expression patterns](3545-regular-expression-patterns.md "Summary of util.Regexp.compile() pattern syntax.").

## Usage

This method scans the current string value and generates a dynamic array with all tokens
terminated by separators matching the regular expression passed as parameter.

The tokens can be terminated by the substring that matches the regular expression, or is
terminated by the end of the string.

If the expression does not match any part of the input, then the resulting array has just one
element, which is a copy of this string.

> **Note:**
>
> When using single or double quoted string literals, backslash is
> interpreted as escape character. Consider using [back
> quotes](../08_language-basics/0588-text-literals.md "Text literals define a character string in an expression.") as delimiters for regular expressions strings, and use backslash characters directly
> as required by the regexp syntax ( `` `\b` `` instead of `"\\b"` or
> `'\\b'` )

If the regular expression is `NULL`, it matches an empty separator at every
position, splitting the source string into its individual characters. The resulting array then
contains one element per character, preceded and followed by an additional empty string element. For
example, the source string `"abc"` will be splitted into array elements `[ "",
"a", "b", "c", "" ]` .

Splitting with an `NULL` regular expression follows the runtime character set.
With a multibyte character set such as UTF-8, the string is split into characters and multibyte
characters are kept together; with a single-byte character set, the string is split into individual
bytes. This behavior depends on the [runtime locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.")
(LC\_ALL/LANG), not on the FGL\_LENGTH\_SEMANTICS setting.

## Example

```
MAIN
    DEFINE s STRING
    DEFINE tokens DYNAMIC ARRAY OF STRING
    DEFINE x INTEGER
    LET s = "aa bb;cc:dd"
    LET tokens = s.split(`[ ;:]`)
    FOR x=1 TO tokens.getLength()
        DISPLAY tokens[x]
    END FOR
END MAIN
```

Output:

```
aa
bb
cc
dd
```

## Related links

**Related concepts**  

[util.Regexp.split](3544-util-regexp-split.md "Splits a string around matches of the current regular expression.")
