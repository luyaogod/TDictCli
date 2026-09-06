---
title: "fgl_ws_setOption()"
source: "fgl-topics/c_gws_configuration_API_fgl_ws_setOption.html"
breadcrumb: "Web services > Reference > Configuration API functions - version 1.3 only > fgl_ws_setOption()"
type: "concept"
---

# fgl_ws_setOption()

> Sets an option flag with a given value, changing the global behavior of the Web Services engine.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_setOption(
   optionName VARCHAR,
   optionValue INTEGER)
```

1. optionName is one of the global [option flags](5077-option-flags.md).
2. optionValue is the value of the flag.

## Example

```
CALL fgl_ws_setOption("http_invoketimeout",5)
```

## Possible runtime errors

- -15511: INVALID\_OPTION\_NAME
