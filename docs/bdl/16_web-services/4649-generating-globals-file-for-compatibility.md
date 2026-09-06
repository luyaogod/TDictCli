---
title: "Generating globals file for GWS server"
source: "fgl-topics/c_gws_handlers_server_generate_globals.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > WS server stubs and handlers > Generating stub file for a GWS server > Generating globals file for compatibility"
type: "concept"
---

# Generating globals file for GWS server

> Use the fglwsdl -legacy tool option to generate legacy code (Genero 3.20 or prior) for the server stub from a WSDL.

This command gets the WSDL information for the Calculator Service defined in [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service.") using the `-legacy`
option of the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool. This allows you to generate a
globals file and a stub file.

```
fglwsdl -s -legacy -soap12 -o ws_calculator http://localhost:8090/Calculator?WSDL
```

The `-s` option specifies the command generates server stubs to be used in a GWS
server application. It generates two files, adding `"Service`" to the file names
generated:

- filename.inc - the globals file, containing declarations
  of global variables that can be used as input or output to functions accessing the Web Service
  operations. In our example, the file is ws\_calculatorService.inc.

  This file
  must be listed in a `GLOBALS` statement at the top of any .4gl
  modules that you write for your GWS Server application. For
  example:

  ```
  GLOBALS "ws_calculatorService.inc"
  ```
- filename.4gl - containing a function that creates the
  service described in the WSDL, publishes the operations of the service, and registers the service.
  In our example, the file is ws\_calculatorService.4gl.

  This file must be
  compiled and imported into your GWS Server application.

It is not advised to modify these files.

## Related links

**Related concepts**  

[Compiling GWS server applications](4668-compiling-gws-server-applications.md "When compiling, remember to include the WSHelper library.")
