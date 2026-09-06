---
title: "fgl_ws_server_setNamespace() (version 1.3)"
source: "fgl-topics/c_gws_server_API_fgl_ws_server_setNamespace.html"
breadcrumb: "Web services > Reference > Server API functions - version 1.3 only > fgl_ws_server_setNamespace() (version 1.3)"
type: "concept"
---

# fgl_ws_server_setNamespace() (version 1.3)

> Defines the namespace of the service on the Web and must be called first, before all other functions of the API.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_server_setNamespace(
   namespace VARCHAR )
```

1. namespace is the name of the namespace.

## Example

```
 CALL fgl_ws_server_setNamespace("http://tempuri.org/")
```
