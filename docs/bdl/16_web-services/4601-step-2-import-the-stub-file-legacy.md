---
title: "Step 2: Import the stub file (legacy)"
source: "fgl-topics/c_gws_client_tutorial_005_legacy_320.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Call the web service > Step 2: Import the stub file > Step 2: Import the stub file (legacy)"
type: "concept"
description: "If you have generated stub files for compatibility with a legacy GWS client application (an application created with Genero 3.20 or prior), add a GLOBALS statement at the top of the .4gl module to ..."
---

# Step 2: Import the stub file (legacy)

If you have [generated stub
files](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.") for compatibility with a legacy GWS client application (an application created with
Genero 3.20 or prior), add a `GLOBALS` statement at the top of the
.4gl module to specify the globals file generated from the WSDL using the
fglwsdl command.

To explain how you code a Genero client application that accesses a
SOAP Web service, this documentation uses the example of a calculator client application. This
application accesses the Add operation in the **MyCalculator** Web service. To write the client,
you simply need to access the WSDL for the Web service. To learn how to program the service, see
[Writing a Web server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .").

In the client calculator example, the client stub files created were named
ws\_calculator.inc and ws\_calculator.4gl. The
`GLOBALS` statement references the file with the .inc
extension:

```
GLOBALS "ws_calculator.inc"
```

## Related links

**Related concepts**  

[Generate a GWS client stub file from the WSDL](4596-generate-the-client-stub-file.md "To access a SOAP Web service, retrieve the WSDL from the service provider to learn about the functions provided by the Web service. The fglwsdl command-line tool can use the WSDL to generate much of the client code for you.")
