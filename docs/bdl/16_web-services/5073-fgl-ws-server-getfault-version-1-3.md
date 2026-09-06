---
title: "fgl_ws_server_getFault() (version 1.3)"
source: "fgl-topics/c_gws_server_API_fgl_ws_server_getFault.html"
breadcrumb: "Web services > Reference > Server API functions - version 1.3 only > fgl_ws_server_getFault() (version 1.3)"
type: "concept"
---

# fgl_ws_server_getFault() (version 1.3)

> Retrieves the last fault string the user has set in a Web-Function, or an empty string if there is none.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_server_getFault()
  RETURNS STRING
```

## Usage

The function returns a string containing the [SOAP](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") fault string.

This function is only for testing the Web Services functions before
they are published on the Web.

## Example

```
 DEFINE div_input RECORD
			a INTEGER,
			b INTEGER
			END RECORD

 DEFINE div_output RECORD
			result INTEGER
			END RECORD

 FUNCTION TestServices()
    DEFINE string VARCHAR(100)
    ...
    # Test divide by zero operation
    LET div_input.a=15
    LET div_input.b=0
    CALL service_operation_div()
    LET string=fgl_ws_server_getFault()
    DISPLAY "Operation div error: ", string
    ...
 END FUNCTION

 FUNCTION service_operation_div()
    ...
    IF div_input.b = 0 THEN
      CALL fgl_ws_server_setFault("Divide by zero")
      RETURN
    END IF
    ...
 END FUNCTION
```
