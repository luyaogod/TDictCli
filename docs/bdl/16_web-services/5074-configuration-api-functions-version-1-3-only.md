---
title: "Configuration API functions - version 1.3 only"
source: "fgl-topics/c_gws_configuration_API_functions.html"
breadcrumb: "Web services > Reference > Configuration API functions - version 1.3 only"
type: "concept"
---

# Configuration API functions - version 1.3 only

> Configuration API functions can modify the behavior of the Web Services engine for the client and/or server.

> **Warning:**
>
> These functions are valid for backward compatibility, but they are not the preferred way to
> handle Genero Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for
> the preferred classes and methods for handling Web services.

| Name | Description |
| --- | --- |
| FUNCTION fgl_ws_setOption( optionName VARCHAR, optionValue INTEGER) | Sets an option flag with a given value, changing the global behavior of the Web Services engine. |
| FUNCTION fgl_ws_getOption( optionName VARCHAR) RETURNS INTEGER | Returns the value of an option flag. |

## Child topics

- [fgl_ws_setOption()](5075-fgl-ws-setoption.md): Sets an option flag with a given value, changing the global behavior of the Web Services engine.
- [fgl_ws_getOption()](5076-fgl-ws-getoption.md): Returns the value of an option flag.
- [Option flags](5077-option-flags.md)
- [WSDL generation option notes](5078-wsdl-generation-option-notes.md)
