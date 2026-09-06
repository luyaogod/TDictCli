---
title: "Generating the stub file for a GWS client"
source: "fgl-topics/c_gws_handlers_client_003.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > Generate the stub file"
type: "concept"
---

# Generating the stub file for a GWS client

> Use the fglwsdl tool to generate the Genero BDL stub from a WSDL URL or file.

This example requests the MyCalculator Web service information from the specified URL. The output
file will have the base name
"ws\_calculator":

```
fglwsdl -c -soap12 -o ws_calculator http://localhost:8090/MyCalculator?WSDL
```

The [fglwsdl](../13_programming-tools/2523-fglwsdl.md) generates filename.4gl.
It contains the definitions of the functions that can be used in your GWS client application to
perform the requested Web service operation, and the code that manages the Web service request. In
our example, the file is ws\_calculator.4gl.

Do not modify this file. Compile and import this file into your GWS client application. For
examples, see [Call the web service](4598-call-the-web-service.md "Import the required libraries and write the program to call a SOAP Web Service.").

## Related links

**Related concepts**  

[SOAP 1.1 and 1.2](4513-soap-1-1-and-1-2.md "GWS is able to communicate with Web services provided with SOAP 1.1 or SOAP 1.2.")

[Generating stub files for a legacy GWS client](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.")

## Child topics

- [Generating stub files for a legacy GWS client](4618-generate-the-stub-files-legacy.md): Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.
