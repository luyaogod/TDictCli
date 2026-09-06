---
title: "base.StringTokenizer.countTokens"
source: "fgl-topics/c_fgl_ClassStringTokenizer_countTokens.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class > base.StringTokenizer methods > base.StringTokenizer.countTokens"
type: "concept"
---

# base.StringTokenizer.countTokens

> Returns the number of tokens left to be returned.

## Syntax

```
countTokens()
  RETURNS INTEGER
```

## Usage

Use the `countTokens()` method to count the number of tokens left to be
returned by the string tokenizer.

This method can be used to get the number of tokens before processing the source string with the
`hasMoreTokens()` and `nextToken()` methods.
