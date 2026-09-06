---
title: "MATCHES"
source: "fgl-topics/c_fgl_operators_MATCHES.html"
breadcrumb: "Language basics > Operators > List of expression elements > Comparison operators > MATCHES"
type: "concept"
---

# MATCHES

> The MATCHES operator returns TRUE if a string matches a given mask.

## Syntax

```
expr [NOT] MATCHES mask [ ESCAPE "char" ]
```

1. expr is any character string expression.
2. mask is a character string expression defining the filter.
3. char is a single char specifying the escape symbol (default is backslash,
   don't forget to double it!).
4. The `NOT` keyword negates the comparison.

## Usage

The mask can be any combination of characters, including the `*`,
`?`, `[`, `]`, `-` and `^`
wildcards:

- The `*` star character matches any string of zero or more characters.
- The `?` question mark matches any single character.
- The `[ ]` brackets match any enclosed character.
- Inside `[ ]`, the `-` (hyphen) between characters means a
  range of characters.
- Inside `[ ]`, the `^` An initial caret matches any character
  that is not listed.

Trailing blanks in the expression are not significant. Consequently, `'a ' MATCHES
'?'` is `TRUE`, but `' ' MATCHES '?'` is
`FALSE`.

The `ESCAPE` clause can be used to define an escape character different from the
default backslash. It must be enclosed in single or double quotes.

A backslash (or the escape character specified by the `ESCAPE` clause) makes the
operator treat the next character as a literal character, even if it is one of the special symbols
in the mask list. This allows you to search for wildcard charachers such as
`*`, `?`, `[`, `]` or
`\`.

If you need to escape a wildcard character, keep in mind that a string constant in the source
code must also escape the backslash character. As a result, if you want to pass a backslash to the
`MATCHES` operator (by using backslash as default escape character), you need to
write four backslashes in the original string constant.

To include the hyphen (`-`) or caret (`^`) characters in a
`[ ]` set of characters in the pattern, put the hyphen at the beginning of the list
and the caret after the first position, there is no need to escape these characters inside `[
]`.

The next table shows some examples of string constants used in the source code and their
equivalent `MATCHES` pattern:

| String constant in source code | MATCHES pattern in memory | Description |
| --- | --- | --- |
| `"*"` | `*` | Matches any character in a non-empty string. |
| `"?"` | `?` | Matches a single character. |
| `"abc*"` | `abc*` | Starts with abc. |
| `"*abc"` | `*abc` | Ends with abc. |
| `"*abc*"` | `*abc*` | Contains abc. |
| `"abc??"` | `abc??` | Starts with abc, followed by two additional characters. |
| `"[a-z]*"` | `[a-z]*` | Starts with a letter in the range a to z. |
| `"[^0-9]*"` | `[^0-9]*` | Must not start with a digit. |
| `"\\*"` | `\*` | Contains a single star character (the \* wildcard is escaped) |
| `"*abc\\\\def*"` | `*abc\\def*` | Contains abc followed by a backslash followed by def (the backslash is escaped) |
| `"*[-^]*"` | `*[-^]*` | Contains hyphen or caret characters. |
| `` "*[ -/:-@[-^`{-~]*" `` | `` *[ -/:-@[-^`{-~]* `` | Contains characters that are not valid for an identifier (to be used with `NOT MATCHES`). Note that the pattern is using character ranges with `-` hyphen: `-/` (space to slash)`:-@` (colon to minus)`[-^` (open sb to caret)`` ` `` (backquote)`{-~` (open cb to tilde) |

## Example

```
MAIN
  IF "55f-plot" MATCHES "55[a-z]-*" THEN
     DISPLAY "Item reference format is correct."
  END IF
END MAIN
```

## Related links

**Related concepts**  

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")

[STRING.matches](../15_library-reference/2925-string-matches.md "Tests if the string matches a regular expression.")

[LIKE](0607-like.md "The LIKE operator returns TRUE if a string matches a given mask.")
