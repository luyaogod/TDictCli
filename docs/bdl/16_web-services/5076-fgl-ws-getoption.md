---
title: "fgl_ws_getOption()"
source: "fgl-topics/c_gws_configuration_API_fgl_ws_getOption.html"
breadcrumb: "Web services > Reference > Configuration API functions - version 1.3 only > fgl_ws_getOption()"
type: "concept"
---

# fgl_ws_getOption()

> Returns the value of an option flag.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_getOption(
   optionName VARCHAR)
  RETURNS INTEGER
```

1. optionName is one of the global [option flags](5077-option-flags.md).

## Example

```
DEFINE value INTEGER
LET value=fgl_ws_getOption("http_invoketimeout")
```

## Possible runtime errors

- -15511: INVALID\_OPTION\_NAME
