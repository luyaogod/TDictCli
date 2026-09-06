---
title: "Generating stub file for a GWS server"
source: "fgl-topics/c_gws_handlers_server_003.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > WS server stubs and handlers > Generating stub file for a GWS server"
type: "concept"
---

# Generating stub file for a GWS server

> Use the fglwsdl -s tool server option to generate the BDL server stub from a WSDL.

You can write a Genero Web Services server application for a Web Service that you have created;
see [Tutorial: Writing a Server
Application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .").

If you want to make sure your Web Service is compatible with that of a third-party (an accounting
application vendor, for example), you use the server stub option of the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") command to obtain the WSDL information that complies with that
vendor's standards, and to generate corresponding file that can be used in your GWS server
application.

You must provide the name of the GWS service module (`serviceImp` in the sample)
where the service finds the functions for the operations at runtime.

The output option (`-o`) specifies the base name for the stub file
("ws\_calculator" in the sample). fglwsdl adds "Service" to the file name
generated. In our example, the file will be named
ws\_calculatorService.4gl

```
fglwsdl -s serviceImp -o ws_calculator http://localhost:8090/Calculator?WSDL
```

For
an more information on this command see [Generate server stub file](../13_programming-tools/2523-fglwsdl.md).

The generated file contains a function that creates the service described in the WSDL, publishes
the operations of the service, and registers the service.

This file must be compiled and imported into your GWS Server application.

## Child topics

- [Generating globals file for GWS server](4649-generating-globals-file-for-compatibility.md): Use the fglwsdl -legacy tool option to generate legacy code (Genero 3.20 or prior) for the server stub from a WSDL.
