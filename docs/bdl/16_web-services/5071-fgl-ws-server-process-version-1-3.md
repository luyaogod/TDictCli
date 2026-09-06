---
title: "fgl_ws_server_process() (version 1.3)"
source: "fgl-topics/c_gws_server_API_fgl_ws_server_process.html"
breadcrumb: "Web services > Reference > Server API functions - version 1.3 only > fgl_ws_server_process() (version 1.3)"
type: "concept"
---

# fgl_ws_server_process() (version 1.3)

> Waits for an incoming SOAP request for a given time (in seconds) and then processes the request, or returns, if there has been no request during the given time.

If a DEFER INTERRUPT or DEFER QUIT instruction has been defined, the function returns even if it
is an infinite wait.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_server_process(
   timeout INTEGER)
  RETURNS INTEGER
```

1. timeout is the maximum waiting time for an incoming request (or -1 for an
   infinite wait)

## Usage

The function can return one of the following values:

- 0 Request has been processed
- -1 Timeout has been reached
- -2 The application server asks the runner to shutdown
- -3 A client connection has been unexpectedly broken
- -4 An interruption has been raised
- -5 The HTTP header of the request was incorrect
- -6 The SOAP envelope was malformed
- -7 The XML document was malformed

## Example

```
MAIN

  DEFINE mystatus INTEGER

  DEFER INTERRUPT

  LET mystatus=fgl_ws_server_process(5)# wait for 5 seconds 

  IF mystatus=0 THEN
    DISPLAY "Request processed."
  END IF
  IF mystatus=-1 THEN
    DISPLAY "No request."
  END IF
  IF mystatus=-2 THEN # terminate the application properly
    EXIT PROGRAM      # if connected to application server
  END IF
  IF mystatus=-3 THEN
    DISPLAY "Client connection unexpectedly broken."
  END IF
  IF mystatus=-4 THEN
    DISPLAY "Server process has been interrupted."
  END IF
  IF mystatus=-5 THEN
    DISPLAY "Malformed or bad HTTP request received."
  END IF
  IF int_flag<>0 THEN
    LET int_flag=0
    EXIT PROGRAM
  END IF

END MAIN
```
