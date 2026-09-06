---
title: "Step 2: Import the stub file"
source: "fgl-topics/c_gws_client_tutorial_005.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Call the web service > Step 2: Import the stub file"
type: "concept"
description: "To explain how you code a Genero client application that accesses a SOAP Web service, this documentation uses the example of a calculator client application. This application accesses the Add ..."
---

# Step 2: Import the stub file

To explain how you code a Genero client application that accesses a
SOAP Web service, this documentation uses the example of a calculator client application. This
application accesses the Add operation in the **MyCalculator** Web service. To write the client,
you simply need to access the WSDL for the Web service. To learn how to program the service, see
[Writing a Web server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .").

Include the following line at the top of the module, where ws\_calculator is
the stub file created command from the WSDL by the fglwsdl tool:

```
IMPORT FGL ws_calculator
```

## Related links

1. [Step 2: Import the stub file (legacy)](4601-step-2-import-the-stub-file-legacy.md)

**Related concepts**  

[Generate a GWS client stub file from the WSDL](4596-generate-the-client-stub-file.md "To access a SOAP Web service, retrieve the WSDL from the service provider to learn about the functions provided by the Web service. The fglwsdl command-line tool can use the WSDL to generate much of the client code for you.")
