---
title: "err_print()"
source: "fgl-topics/c_fgl_BuiltInFunctions_ERR_PRINT.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > err_print()"
type: "concept"
---

# err_print()

> Prints in the error line the text corresponding to an error number.

## Syntax

```
FUNCTION err_print(
   messageId INTEGER )
```

1. messageId is a runtime error or an Informix® SQL error.

## Usage

The `err_print()` function displays to the screen the
error message corresponding to the number passed as parameter.
The message will be displayed in the error line defined by the program.

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
