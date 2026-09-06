---
title: "base.StringTokenizer methods"
source: "fgl-topics/c_fgl_ClassStringTokenizer_methods.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class > base.StringTokenizer methods"
type: "concept"
description: "Table 1. Class methods Name Description base.StringTokenizer.create ( str STRING, delimiters STRING ) RETURNS base.StringTokenizer Create a string tokenizer object. base.StringTokenizer.createExt ( ..."
---

# base.StringTokenizer methods

| Name | Description |
| --- | --- |
| base.StringTokenizer.create( str STRING, delimiters STRING ) RETURNS base.StringTokenizer | Create a string tokenizer object. |
| base.StringTokenizer.createExt( str STRING, delimiters STRING, escapeChar STRING, withNulls BOOLEAN ) RETURNS base.StringTokenizer | Create a string tokenizer object with escape char and null handling. |

| Name | Description |
| --- | --- |
| countTokens() RETURNS INTEGER | Returns the number of tokens left to be returned. |
| hasMoreTokens() RETURNS BOOLEAN | Returns `TRUE` if there are more tokens to return. |
| nextToken() RETURNS STRING | Returns the next token found in the source string. |
