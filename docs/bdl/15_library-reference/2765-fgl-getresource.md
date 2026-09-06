---
title: "fgl_getresource()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETRESOURCE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getresource()"
type: "concept"
---

# fgl_getresource()

> Returns the value of an FGLPROFILE entry.

## Syntax

```
FUNCTION fgl_getresource(
   name STRING )
  RETURNS STRING
```

1. *name* is the FGLPROFILE entry name to be read.

## Usage

The `fgl_getresource()` function reads the
[FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") file(s) and returns
the value defined for the entry passed as parameter.

If the entry does not exist in the configuration file, the function returns [NULL](../08_language-basics/0572-null.md "The NULL constant defines a non-value.").

> **Important:**
>
> FGLPROFILE entry names are not case sensitive.

If multiple entries are defined with the same name (this can happen especially when several
profile files are defined in the FGLPROFILE environment variable), the last entry found wins.

`fgl_getresource()` is a global built-in function, consider using the new [`base.Application.getResourceEntry()`](2975-base-application-getresourceentry.md "Returns the value of a FGLPROFILE entry.") method instead.

## Related links

**Related concepts**  

[base.Application.getResourceEntry](2975-base-application-getresourceentry.md "Returns the value of a FGLPROFILE entry.")
