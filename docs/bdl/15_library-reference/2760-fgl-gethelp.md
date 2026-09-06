---
title: "fgl_gethelp()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETHELP.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_gethelp()"
type: "concept"
---

# fgl_gethelp()

> Reads the current help file, returning help text based on the help identifier.

## Syntax

```
FUNCTION fgl_gethelp(
   id INTEGER )
  RETURNS STRING
```

1. id is the help text identifier.

## Usage

The `fgl_gethelp()` function returns the text corresponding to
the help message number passed as parameter.

The text is read from the current help file. The current help file is
defined by the [`OPTIONS HELP FILE`](../09_advanced-features/0931-defining-the-message-file.md "The OPTIONS HELP FILE instruction defines the name of the message file.")
instruction.
