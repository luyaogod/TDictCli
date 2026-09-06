---
title: "Text literals"
source: "fgl-topics/c_fgl_literals_STRING.html"
breadcrumb: "Language basics > Literals > Text literals"
type: "concept"
---

# Text literals

> Text literals define a character string in an expression.

## Syntax 1: Using double quote delimiters

```
" char [...] "
```

## Syntax 2: Using single quote delimiters

```
' char [...] '
```

## Syntax 3: Using the % localized string marker

```
%" char [...] "
```

## Syntax 4: Using back quote delimiters

```
` char [...] `
```

1. In all above syntaxes, char can be any character supported in the [current locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.").
2. Syntax 3 defines a [localized string](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.").
3. In syntax 1, 2 and 3, char can be a `\` backslash escape
   combination as described in the following table:

   | Escape sequence | Description |
| --- | --- |
| `\\` | The backslash character |
| `\"` | double-quote character (to be used with double quote delimiters) |
| `\'` | single-quote character (to be used with single quote delimiters) |
| `\n` | The newline character |
| `\r` | The carriage-return character |
| `\0` | The null character |
| `\f` | The form-feed character |
| `\t` | The tab character |
| `\xNN` | ASCII character defined by the hexadecimal code NN |

## Usage

A text literal (or character string literal) defines a character string constant containing valid
characters in the current application character set.

The application character set is defined by the [current
locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.").

An empty string (`""`) is equivalent to [`NULL`](0572-null.md "The NULL constant defines a non-value.").

A text literal can be written on multiple lines, the compiler merges lines by removing the newline
character.

With double quoted and single quoted string literals, the escape character is the backslash
character (`\`). When using raw string literals with back quote delimiters, there is
no escape character.

When using single quotes as delimiters, double quotes can be used as is inside the string,
while single quotes must be doubled or escaped with a
backslash:

```
DISPLAY ' " "  '' \' '      -- shows:  " "  ' '
```

When using double quotes as delimiters, single quotes can be used as is inside the string,
while double quotes must be doubled or escaped with a
backslash:

```
DISPLAY " "" \"  ' ' "      -- shows:  " "  ' '
```

When using back quotes as delimiters, you can specify any character inside the string literal,
except the back quote character: Single and double quotes can just be specified as is (not backslash
escaping needed):

```
DISPLAY ` " " ' ' `         -- shows:  " " ' '
```

With single or double quote delimiters, special characters can be specified with backslash
escape symbols. Use for example `\n` to insert a new-line character in a string
literal:

```
DISPLAY "First line\nSecond line"
```

When using back quote delimiters, the backslash escape symbol has no specific
meaning:

```
DISPLAY ` \n \\ \ `    -- shows:  \n \\ \
```

When using single or double quotes and the string literal is written on multiple lines, the
string pieces are concatenated:

```
DISPLAY "piece1
 piece2
 piece3" -- shows:    piece1 piece2 piece3
```

When using back quote delimiters and the string literal is written on multiple lines, the line
breaks are kept:

```
DISPLAY `line1
line2
line3` -- shows:    line1<NL>line2<NL>line3
```

With single and double quote delimiters, the `\xNN`
hexadecimal notation allows you to specify control characters in a string literal. Only ASCII codes
(<=0x7F) are allowed:

```
DISPLAY "\x65"    -- shows:   e
DISPLAY `\x65`    -- shows:   \x65
```

When prefixing the string with a `%` percent sign, you specify the key for a
localized string. This key will be replaced at runtime by another text, found in string resource
files:

```
MESSAGE %"orders.message.shipped"
```

For more details, see [Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")

## Example

```
MAIN
  
  DISPLAY "Some text in double quotes"
  DISPLAY 'Some text in single quotes'
  DISPLAY `Some text in back quotes`

  DISPLAY "  \"  ""  "
  DISPLAY '  \'  ''  '
  DISPLAY `  "  '  `

  DISPLAY 'Newline character here:\n ... and continue on next line.'

  DISPLAY "This is a double quoted string
 on multiple lines, that
 will show concatenated."

  DISPLAY `This is a back quoted string
 on multiple lines, where newlines
 are kept.`

  DISPLAY ( "" IS NULL )
  DISPLAY ( '' IS NULL )
  DISPLAY ( `` IS NULL )

END MAIN
```

Output:

```
Some text in double quotes
Some text in single quotes
Some text in back quotes
  "  "  
  '  '  
  "  '  
Newline character here:
 ... and continue on next line.
This is a double quoted string on multiple lines, that will show concatenated.
This is a back quoted string
 on multiple lines, where newlines
 are kept.
     1
     1
     1
```

## Related links

**Related concepts**  

[String delimiters](0547-string-delimiters.md "String literals need to be delimited by specific characters.")

[String expressions](0597-string-expressions.md "This section covers string expression evaluation rules.")

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")
