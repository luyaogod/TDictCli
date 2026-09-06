---
title: "base.StringTokenizer.createExt"
source: "fgl-topics/c_fgl_ClassStringTokenizer_createExt.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class > base.StringTokenizer methods > base.StringTokenizer.createExt"
type: "concept"
---

# base.StringTokenizer.createExt

> Create a string tokenizer object with escape char and null handling.

## Syntax

```
base.StringTokenizer.createExt(
   str STRING,
   delimiters STRING,
   escapeChar STRING,
   withNulls BOOLEAN )
  RETURNS base.StringTokenizer
```

1. str is the character string to be parsed.
2. delimiters defines the delimiters to be used.
3. escapeChar defines the character to escape a delimiter.
4. withNulls indicates if empty tokens must be returned.

## Usage

Use the `base.StringTokenizer.createExt()` class method to create a string
tokenizer object, with escape character and empty tokens handling.

The created object must be assigned to a program variable defined with the
`base.StringTokenizer` type.

The method can take a unique or multiple delimiters into account. A delimiter is always one
character long.

To specify a backslash as a delimiter, use double backslashes in
both the source string and as the delimiter.

When defining an escape character with the third parameter, the delimiters can be escaped in
the source string.

When passing `TRUE` for the last parameter of
`base.StringTokenizer.createExt()`, empty tokens are taken into account. If an empty
token is found, the `nextToken()` method might return an empty string.

In the source string, leading and trailing delimiters or the amount of delimiters between two
tokens affects the number of tokens.

## Example

```
DEFINE tok base.StringTokenizer
LET tok = base.StringTokenizer.createExt("||\\|aaa||bbc|","|","\\",FALSE)
```

For more detailed examples, see [Examples](3078-examples.md "base.StringTokenizer usage examples.").
