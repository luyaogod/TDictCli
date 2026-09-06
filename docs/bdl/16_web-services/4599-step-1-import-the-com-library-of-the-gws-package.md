---
title: "Step 1: Import the COM library of the GWS package"
source: "fgl-topics/c_gws_client_tutorial_004.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Call the web service > Step 1: Import the COM library of the GWS package"
type: "concept"
description: "The classes that make up the Genero Web Services Library ( com ) contain the methods associated with creating and publishing a Web Service. You must import the library to use any of these methods in ..."
---

# Step 1: Import the COM library of the GWS package

The classes that make up the Genero Web Services Library (`com`) contain the
methods associated with creating and publishing a Web Service. You must import the library to use
any of these methods in your client application.

To explain how you code a Genero client application that accesses a
SOAP Web service, this documentation uses the example of a calculator client application. This
application accesses the Add operation in the **MyCalculator** Web service. To write the client,
you simply need to access the WSDL for the Web service. To learn how to program the service, see
[Writing a Web server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .").

In our example, the calculator client application calls a method of the
`com.WebServiceEngine` class to set a timeout, defining the length of time the client
will wait for the service to respond. The module must import the library; add the following to the
top of the module:

```
IMPORT com
```

If your application uses data types from the `XML` class data types, add the
following:

```
IMPORT xml
```

If you are using the [wsError](4604-handle-gws-server-errors.md "When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.")
record, you must import the [WSHelper](4923-wshelper-library.md "The WSHelper library provides a set of functions to help with web services.") module. Add the
following:

```
IMPORT FGL WSHelper
```

## Related links

**Related concepts**  

[The com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.")
