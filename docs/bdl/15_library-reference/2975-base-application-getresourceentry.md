---
title: "base.Application.getResourceEntry"
source: "fgl-topics/c_fgl_ClassApplication_getResourceEntry.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Application class > base.Application methods > base.Application.getResourceEntry"
type: "concept"
---

# base.Application.getResourceEntry

> Returns the value of a FGLPROFILE entry.

## Syntax

```
base.Application.getResourceEntry(
   name STRING )
  RETURNS STRING
```

1. name is the name of a FGLPROFILE entry.

## Usage

The `base.Application.getResourceEntry()` method reads the [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") file(s) and returns the value defined for the
entry passed as parameter.

If the entry does not exist in the configuration file, the method returns [NULL](../08_language-basics/0572-null.md "The NULL constant defines a non-value.").

Note that FGLPROFILE entry names are case insensitive.

If multiple entries are defined with the same name (this can happen especially when several
profile files are defined in the FGLPROFILE environment variable), the last entry found wins.

The `fgl_getresource()` built-in function is equivalent to
`base.Application.getResouceEntry()`.

## Example

```
MAIN
  DISPLAY base.Application.getResourceEntry("mycompany.params.logmode")
END MAIN
```

## Related links

**Related concepts**  

[fgl\_getresource()](2765-fgl-getresource.md "Returns the value of an FGLPROFILE entry.")
