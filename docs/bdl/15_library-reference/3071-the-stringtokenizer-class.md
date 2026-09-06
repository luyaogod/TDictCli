---
title: "The StringTokenizer class"
source: "fgl-topics/c_fgl_ClassStringTokenizer.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class"
type: "concept"
---

# The StringTokenizer class

> The base.StringTokenizer class is designed to parse a string to extract tokens based on delimiters.

The steps to use a string tokenizer are:

1. Define a variable of the `base.StringTokenizer` type.
2. Create the string tokenizer object with one of the create methods, passing the string to be parsed as parameter.
3. Optionally, count the number of tokens with `countTokens()` before
   processing.
4. Use a `WHILE` loop to process the different tokens, by using `hasMoreTokens()`
   as loop condition and `nextToken()` inside the loop body to get the next token.

## Child topics

- [base.StringTokenizer methods](3072-base-stringtokenizer-methods.md)
- [Examples](3078-examples.md): base.StringTokenizer usage examples.
