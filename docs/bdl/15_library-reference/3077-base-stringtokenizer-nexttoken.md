---
title: "base.StringTokenizer.nextToken"
source: "fgl-topics/c_fgl_ClassStringTokenizer_nextToken.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class > base.StringTokenizer methods > base.StringTokenizer.nextToken"
type: "concept"
---

# base.StringTokenizer.nextToken

> Returns the next token found in the source string.

## Syntax

```
nextToken()
  RETURNS STRING
```

## Usage

The `nextToken()` method parses the source string for tokens, following the
creation method used, and returns the next token if found.

The method returns `NULL`, if there are no more tokens to be returned.

When the tokenizer object was created from the [`base.StringTokenizer.createExt()`](3074-base-stringtokenizer-createext.md "Create a string tokenizer object with escape char and null handling.") method, with the
withNulls parameter set to `TRUE`, the
`nextToken()` method returns an empty string (non-NULL), when an empty token is
found.

Use the `hasMoreTokens()` method to check if more tokens are to be read.
