---
title: "Whitespace separators"
source: "fgl-topics/c_fgl_language_features_whitespace_separators.html"
breadcrumb: "Language basics > Syntax features > Whitespace separators"
type: "concept"
---

# Whitespace separators

> Whitespace characters are used to separate language elements.

Genero BDL is free-form, like C or Pascal, and generally ignores TAB characters, LINEFEED
characters, comments, and extra blank spaces between instructions and language elements.

You can freely use these whitespace characters to enhance the readability of your source
code.

Blank (ASCII 32) characters act as delimiters in some contexts. Blank spaces must separate
successive keywords or identifiers, but cannot appear within a keyword or identifier.

Pairs of double ( " ) or single ( ' ) quotation marks must delimit any character string that
contains a blank space (ASCII 32) or other whitespace character, such as LINEFEED or RETURN.
