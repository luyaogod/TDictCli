---
title: "String delimiters"
source: "fgl-topics/c_fgl_language_features_quotation_marks.html"
breadcrumb: "Language basics > Syntax features > String delimiters"
type: "concept"
---

# String delimiters

> String literals need to be delimited by specific characters.

In the Genero BDL language, string literals are delimited by single quote (`'`),
double quote (`"`) or back quote (`` ` ``) characters:

```
'My character string'
"My character string"
`My character string`
```

Do not mix quote as delimiters of the same string. The following is not a valid character
string:

```
'This is not a valid character string"
```

To include literal quotation marks within a character string, precede each quotation mark with
the backslash (\), or else enclose the string between a pair of quotes that :

```
'Type \'Y\' to accept changes.'
"Type 'Y' to accept changes."
'Type "Y" to accept changes.'
```

Use the back quotes, to write the string literal as raw text, without interpretation of the
backslash escape character. In the next example, the string literal will contain 4 characters: a
backslash, the letter `n`, another backslash, and the letter `r`. With
single or double quotes delimiters, `\n` represents a newline and `\r`
a carriage-return:

```
`\n\r`
```

A string literal can be written on multiple lines. The compiler
merges lines by removing the newline character.

In the SQL language, the standard specifications recommend that
you use single quotes for string literals and double quotes for database
object identifiers like table or column names. When accessing a non-Informix
database, double quotation marks might not be recognized as database
object name delimiters. As a general rule, use single quoted string
literals in SQL statements, and use non-quoted, lowercase database
object identifiers.

## Related links

**Related concepts**  

[Text literals](0588-text-literals.md "Text literals define a character string in an expression.")
