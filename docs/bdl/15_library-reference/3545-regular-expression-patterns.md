---
title: "Regular expression patterns"
source: "fgl-topics/c_fgl_ext_util_Regexp_pattern_syntax.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > Regular expression patterns"
type: "concept"
---

# Regular expression patterns

> Summary of util.Regexp.compile() pattern syntax.

## Defining regular expressions in strings literals

Regular expression often use the backslash character to start a pattern element. For example, the
backslash followed by the lowercase letter d (`\d`) expresses a digit at the given
position.

Consider using back quotes as string delimiters for regular expressions: This will define a raw
string literal without having backslash interpreted as the escape character as in single or double
quoted string literals:

```
LET re = util.Regexp.compile(`^\d\d$`)
```

For more details, see [Text literals](../08_language-basics/0588-text-literals.md "Text literals define a character string in an expression.").

## Character classes

| Syntax | Description |
| --- | --- |
| `.` (dot) | Any single character except '`\n`', '`\r`'. |
| `\d` | Any digit. Equivalent to `[0-9]`. |
| `\D` | Any character that is not a digit. Equivalent to `[^0-9]`. |
| `\w` | ASCII letter, digit or underscore. Equivalent to `[A-Za-z0-9_]`. |
| `\W` | Any character that is not an ASCII letter, digit or underscore. Equivalent to `[^A-Za-z0-9_]`. |
| `\s` | Any white space character, including space, tab, form feed, line feed, and other Unicode spaces. |
| `\S` | Any character that is not a white space. |
| `\b` | The boundary of a word (start-of-word or end-of-word) |
| `\B` | Not the boundary of a word (inverse of `\b`) |
| `[\b]` | The backspace character (`BS`, `^H`, `ASCII 08`). |
| `\t` | An horizontal tab. |
| `\r` | A carriage return. |
| `\n` | A linefeed. |
| `\v` | A vertical tab. |
| `\f` | A form-feed. |
| `\cx` | The control character corresponding to x. |
| `\xhh` | The character with the code hh (two hexadecimal digits). |

## Assertions

| Syntax | Description |
| --- | --- |
| `^` | The beginning of the input. |
| `$` | The end of the input. |
| `\b` | A word boundary. |
| `\B` | Not a word boundary. |
| `\G` | The end of the previous match. |
| `x(?=y)` | Matches "x" only if "x" is followed by "y". |
| `x(?!y)` | Matches "x" only if "x" is not followed by "y". |
| `(?<=y)x` | Matches "x" only if "x" is preceded by "y". |
| `(?<!y)x` | Matches "x" only if "x" is not preceded by "y". |

## Groups and ranges

| Syntax | Description |
| --- | --- |
| `x\|y` | Matches either "x" or "y". |
| `[x]` | Matches any one of the enclosed characters. |
| `[^x]` | Matches anything that is not enclosed in the brackets. |
| `(x)` | Matches "x" and remembers the match. |
| `(?:x)` | Matches "x" but does not remember the match. |
| `(?<name>x)` | Matches "x" and stores it on the groups property of the returned matches under the name specified by name. The angle brackets (< and >) are required for the group name. |
| `\nn` | A back reference to the last substring matching the nn'st parenthetical in the regular expression (counting left parentheses). |
| `\k<name>` | A back reference to the last substring matching the named capture group specified by name. |

## Quantifiers

| Syntax | Description |
| --- | --- |
| `x*` | Matches the preceding item "x" 0 or more times. |
| `x+` | Matches the preceding item "x" 1 or more times. Equivalent to `{1,}`. |
| `x?` | Matches the preceding item "x" 0 or 1 times. If used immediately after any of the quantifiers `*`, `+`, `?`, or `{}`, makes the quantifier non-greedy (matching the minimum number of times), as opposed to the default, which is greedy (matching the maximum number of times). |
| `x{n}` | Where "n" is a positive integer, matches exactly "n" occurrences of the preceding item "x". |
| `x{n,}` | Where "n" is a positive integer, matches at least "n" occurrences of the preceding item "x". |
| `x{n,m}` | Where "n" is 0 or a positive integer, "m" is a positive integer, and `m>n`, matches at least "n" and at most "m" occurrences of the preceding item "x". |

## More regular expression patterns

For more details about supported regular expression patterns, refer to [the PRCE documentation](https://www.pcre.org/current/doc/html/pcre2pattern.html).
