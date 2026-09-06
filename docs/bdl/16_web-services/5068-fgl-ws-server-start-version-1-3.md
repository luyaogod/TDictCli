---
title: "fgl_ws_server_start() (version 1.3)"
source: "fgl-topics/c_gws_server_API_fgl_ws_server_start.html"
breadcrumb: "Web services > Reference > Server API functions - version 1.3 only > fgl_ws_server_start() (version 1.3)"
type: "concept"
---

# fgl_ws_server_start() (version 1.3)

> Creates and starts the Web services server.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_server_start(
   tcpPort VARCHAR )
```

1. tcpPort is a string representing either:
   - the **socket port number** (for a single Web Service server)
   - the **host** and **port** value separated by a colon (for a Web Service server connecting
     to an application server). The value of **port** is an offset beginning at 6400.

If the FGLAPPSERVER environment variable is set, the tcpPort value is
ignored, and replaced by the value of FGLAPPSERVER.

## Usage

For development or testing purposes, you may start a Web Service server as a single server where
only one request at a time will be able to be processed. For deployment, you may start a Web Service
server with an application server able to handle several connections at one time using a
load-balancing algorithm. The value of the parameter passed to the function determines which method
is used.

## Examples:

To start a standalone Web Service server:

```
 CALL fgl_ws_server_start("8080") # A single Server is listening
                                  # on port number: 8080
```

To start a Web Service server attempting to connect to an application server:

```
 CALL fgl_ws_server_start("zeus:5") # The server attempt to connect
                                    # to an application server located
                                    # on host zeus and listening
                                    # on the port number 6405
```

## Possible runtime errors

- -15504: PORT\_ALREADY\_USED
- -15514: PORT\_NOT\_NUMERIC
- -15515: NO\_AS\_FOUND
- -15516: LICENSE\_ERROR
