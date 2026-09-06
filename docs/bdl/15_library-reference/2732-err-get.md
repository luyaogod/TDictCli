---
title: "err_get()"
source: "fgl-topics/c_fgl_BuiltInFunctions_ERR_GET.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > err_get()"
type: "concept"
---

# err_get()

> Returns the text corresponding to an error number.

## Syntax

```
FUNCTION err_get(
   messageId INTEGER ) 
  RETURNS STRING
```

1. messageId is a runtime error or an Informix® SQL error.

## Usage

The `err_get()` function returns the error message corresponding to the number
passed as parameter. This can be a Genero BDL error (like -6324), or an IBM® Informix SQL error message.

> **Important:**
>
> IBM
> Informix SQL message numbers can only be supported, if
> the program is connected to an Informix
> database, and an SQL error occurred just before calling this function. Do not use this function in
> the context of SQL execution, when using different type of database servers.

## Related links

**Related reference**  

[Genero BDL errors](4483-genero-bdl-errors.md "System error messages sorted by error number.")
