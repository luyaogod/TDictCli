---
title: "Generating stub files for a legacy GWS client"
source: "fgl-topics/c_gws_handlers_server_generate_globals_client.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > Generate the stub file > Generate the stub files (legacy)"
type: "concept"
---

# Generating stub files for a legacy GWS client

> Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.

This example command requests the Web service information for the "MyCalculator" service. The
`-legacy` option specifies to generate [legacy code](../13_programming-tools/2523-fglwsdl.md).

```
fglwsdl -c -legacy -soap12 -o ws_calculator http://localhost:8090/MyCalculator?WSDL
```

The command generates the globals file and the stub file. The output files will have the base
name "ws\_calculator".

- filename.inc - the globals file, containing declarations
  of global variables that can be used as input or output to functions accessing the Web Service
  operations. In our example, the file is ws\_calculator.inc.

  This file must be
  listed in a `GLOBALS` statement at the top of any .4gl modules
  that you write for your GWS client application. For
  example:

  ```
  GLOBALS "ws_calculator.inc"
  ```
- filename.4gl - containing the definitions of the
  functions that can be used in your GWS client application to perform the requested Web service
  operation, and the code that manages the Web service request. In our example, the file is
  ws\_calculator.4gl.

  The stub file must be compiled and imported into your GWS
  client application.

It is not advised to modify these files.

## Related links

**Related concepts**  

[Example globals file](4619-example-globals-file.md "The example WSDL file for the Calculator Web Service provides information about the service.")

## Child topics

- [Example globals file](4619-example-globals-file.md): The example WSDL file for the Calculator Web Service provides information about the service.
