---
title: "Step 1: Obtain information about Web service resources"
source: "fgl-topics/c_gws_rest_client_webservice_resources.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application > Step 1: Obtain information about resources"
type: "concept"
---

# Step 1: Obtain information about Web service resources

> To access a RESTful Web services server, you must first get information about the services offered by a server from the service provider.

In REST, the name of the resources or functions and what query parameters or arguments to use
with them are generally all that you require, as this information contains everything the client
needs to interact with the required resource.

## Add Function

In the code example for the [Calculator RESTFul Web
services server application](4877-calculator-server-source.md "The source code for the server-side application included in the RESTful Web services calculator demo."), we know that the calculator demo service performs various
calculator functions. Our client application simply wants to access the function to add values.

```
# Calculator Web Service add function
 
FUNCTION add()
  LET add_out.r = add_in.a + add_in.b
  LET add_out.status.code = 0
  LET add_out.status.desc = "OK"
END FUNCTION
```

The HTTP request must therefore use a URI with a resource name and the
required parameters that matches the calculator function defined in the calculator Web service
server side.

```
http[s]://host:port/resource-name[?query-params]
```

1. host:port is the
   base URL from the server.
2. resource-name is the name of the resource.
3. query-params is a set of URL query parameters.

Therefore, **Add** becomes the resource name part of the URI we need to include in our code.

```
# Sample request to add two numbers

GET http://localhost:8090/add?a=1&b=2
```

In the next step we start to code the client application, [Step 2: Import extension packages (com, xml, util)](4857-step-2-import-extension-packages-com-xml-util.md "The functions you need to create a REST Web Service client application are contained in the classes that make up the com package of the Genero Web Services (GWS). Use the IMPORT statement to include the required packages.")
