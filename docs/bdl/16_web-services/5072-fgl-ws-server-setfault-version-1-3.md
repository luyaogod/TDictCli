---
title: "fgl_ws_server_setFault() (version 1.3)"
source: "fgl-topics/c_gws_server_API_fgl_ws_server_setFault.html"
breadcrumb: "Web services > Reference > Server API functions - version 1.3 only > fgl_ws_server_setFault() (version 1.3)"
type: "concept"
---

# fgl_ws_server_setFault() (version 1.3)

> Return a SOAP fault string to the client at the end of the function's execution.

This function can be called in a published Web-Function.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_server_setFault(
   faultMessage VARCHAR )
```

1. faultMessage is a string containing the SOAP Fault string that will be
   returned to the client.

## Example

```
 CALL fgl_ws_server_setFault(
    "The server is not able to manage this request.")
```
