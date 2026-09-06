---
title: "err_quit()"
source: "fgl-topics/c_fgl_BuiltInFunctions_ERR_QUIT.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > err_quit()"
type: "concept"
---

# err_quit()

> Prints in the error line the text corresponding to an error number and terminates the program.

## Syntax

```
FUNCTION err_quit(
   messageId INTEGER )
```

1. messageId is a runtime error or an Informix® SQL error.

## Usage

The `err_quit()` function prints the error message
corresponding to the number passed as parameter.
The message will be displayed in standard error stream and the program
will terminate.

> **Important:**
>
> IBM®
> Informix SQL message numbers can only be supported, if
> the program is connected to an Informix
> database, and an SQL error occurred just before calling this function. Do not use this function in
> the context of SQL execution, when using different type of database servers.

## Related links

**Related reference**  

[Genero BDL errors](4483-genero-bdl-errors.md "System error messages sorted by error number.")
