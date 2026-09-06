---
title: "fgl_ws_server_generateWSDL() (version 1.3)"
source: "fgl-topics/c_gws_server_API_fgl_ws_server_generateWSDL.html"
breadcrumb: "Web services > Reference > Server API functions - version 1.3 only > fgl_ws_server_generateWSDL() (version 1.3)"
type: "concept"
---

# fgl_ws_server_generateWSDL() (version 1.3)

> Generates the WSDL file based on the BDL-server program.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_server_generateWSDL(
   serviceName VARCHAR,
   serviceLocation VARCHAR,
   fileName VARCHAR )
  RETURNS INTEGER
```

1. *serviceName* is the name of the web service.
2. *serviceLocation* is the URL of the server.
3. *fileName* is the name of the file that will be generated.

## Usage

The function returns:

- 0 if the file has been correctly generated.
- Any other values if the operation has failed.

## Example

```
DEFINE mystatus INTEGER

LET mystatus=fgl_ws_server_generateWSDL(
 "CustomerService",
 "http://localhost:8080",
 "C:/mydirectory/myfile.wsdl")

IF mystatus=0 THEN
  DISPLAY "Generation of WSDL done..."
ELSE
  DISPLAY "Generation of WSDL failed!"
END IF
```
