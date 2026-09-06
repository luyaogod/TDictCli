---
title: "Writing your functions"
source: "fgl-topics/c_gws_handlers_server_005.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > WS server stubs and handlers > Writing your functions"
type: "concept"
---

# Writing your functions

> Write functions that implement the functions in the stub modules. It allows you to create your own version of the function.

The ws\_calculator.inc WSDL file provides you with the global input and
output records and function names that allow you to write your own code implementing the
operations.

Your new code should not be written in the generated modules. For example, do not add your
own version of the Add function to the generated ws\_calculator.4gl
module; it can be included in your module containing the `MAIN` program block,
or in a separate module to be included as part of the Web server application. The function
must use the generated definitions for the global input and output records.

In your version of the Add operation, for example, this function adds 100 to the sum of the
variables in the input
record:

```
FUNCTION Add()
  LET AddResponse.r = (Add.a + Add.b) + 100
END FUNCTION
```

See [Tutorial: Writing a
Server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .") for more information. The demo/WebServices
subdirectory of your Genero installation directory contains complete examples of Server
Applications.
