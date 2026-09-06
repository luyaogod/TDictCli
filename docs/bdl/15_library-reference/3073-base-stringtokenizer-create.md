---
title: "base.StringTokenizer.create"
source: "fgl-topics/c_fgl_ClassStringTokenizer_create.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class > base.StringTokenizer methods > base.StringTokenizer.create"
type: "concept"
---

# base.StringTokenizer.create

> Create a string tokenizer object.

## Syntax

```
base.StringTokenizer.create(
   str STRING,
   delimiters STRING )
  RETURNS base.StringTokenizer
```

1. str is the character string to be parsed.
2. delimiters defines the delimiters to be used.

## Usage

Use the `base.StringTokenizer.create()` class method to create a string
tokenizer object.

The created object must be assigned to a program variable defined with the
`base.StringTokenizer` type.

The method can take a unique or multiple delimiters into account. A delimiter is always one
character long.

To specify a backslash as a delimiter, use double backslashes in
both the source string and as the delimiter.

The empty tokens are not taken into account, and no escape character is defined for the
delimiters. In order to return empty tokens, use the [`base.StringTokenizer.createExt()`](3074-base-stringtokenizer-createext.md "Create a string tokenizer object with escape char and null handling.") method.

## Example

```
DEFINE tok base.StringTokenizer
-- Using a single pipe delimiter
LET tok = base.StringTokenizer.create("aaa|bbb|ccc","|")
-- Using several delimiters
LET tok = base.StringTokenizer.create("aaa|bbb;ccc+ddd","|+;")
```

For more detailed examples, see [Examples](3078-examples.md "base.StringTokenizer usage examples.").
