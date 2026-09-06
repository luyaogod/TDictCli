---
title: "Source comments"
source: "fgl-topics/c_fgl_language_features_comment.html"
breadcrumb: "Language basics > Syntax features > Source comments"
type: "concept"
---

# Source comments

> The --, # and { } characters can be used to add source comments.

## Comment markers

A source comment is text in the source code to assist human readers, but which BDL ignores.

A source comment can be specified by any of the following:

- A pair of minus signs ( `--` ) indicates a comment that terminates at the end of
  the current line. This comment indicator conforms to the ANSI standard for SQL.
- The hash character ( `#` ) indicates a comment that terminates at the end of the
  current line.
- A starting left-brace ( `{` ) starts a comment. It can be followed by any
  character (including line breaks). The comment ends when the closing right-brace (
  `}` ) symbol is found.

```
MAIN
  -- DISPLAY "This line will be ignored."
  # DISPLAY "This line will be ignored."
  {
  DISPLAY "This line will be ignored."
  DISPLAY "This line will be ignored."
  }
  DISPLAY "Hello, World"
END MAIN
```

Notes:

- Within a quoted string, the compiler interprets comment indicators as literal characters, rather
  than as comment indicators.
- You cannot use curly brackets ( `{ }` ) to nest comments within comments.
- Comments cannot appear in the form section defining a layout grid, such as
  `SCREEN`, `TABLE`, `TREE`, or
  `GRID`.
- The `#` symbol cannot indicate comments in an SQL statement block nor in the text
  of a prepared statement.
- You cannot specify consecutive minus signs ( `--` ) in arithmetic expressions, as
  BDL interprets what follows as a comment. Instead, use a blank space or parentheses to separate
  consecutive arithmetic minus signs.

## I4GL/BDL language specific comments

The `--#` specific comment indicator is used to distinguish Informix® 4GL code from Genero BDL code.

When writing a comment with `--#`, the remaining line content will be taken into
account by the Genero BDL compiler, but it will be ignored by the I4GL compiler.

This conditional code compilation technique can be inverted by enclosing code blocks between
`--#{` and `--#}` comments:

```
MAIN
  --# DISPLAY "Ignored by I4GL, but compiled with BDL."
  --#{
      DISPLAY "Ignored by BDL, but compiled with I4GL."
  --#}
END MAIN
```

To summarize:

- Code lines starting with `--#` are compiled with Genero BDL, but ignored by Informix 4GL.
- Code blocks surrounded with `--#{` and `--#}` are compiled with
  Informix 4GL, but ignored by Genero BDL.

## Code beautifier control comments

Specific comments can be used to skip the formatting of a group of lines, by using the
`fgl-format off/on`
keywords:

```
# fgl-format off
DEFINE    cnt    INTEGER
DEFINE    ratio  DECIMAL(10,2)
# fgl-format on
```

For more details, see [Source code beautifier](../13_programming-tools/2636-source-code-beautifier.md "Reformat the source code for better readability.").
