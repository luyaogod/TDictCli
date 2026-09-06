---
title: "Create and publish the Web services operation"
source: "fgl-topics/c_gws_function_declaration_005.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web services server function > Create and publish the Web services operation"
type: "concept"
---

# Create and publish the Web services operation

> Provide your Web service and its operation to users who can access it on the net.

Methods are available in the Genero Web Services library ([com](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.")) to:

- Define the Web Service, by creating a WebService object
- Register the HTTP input and output methods
- Define the Web Services operation for your function, by creating
  a WebOperation object
- Publish the operation - associate it with the Web Service object
  that you defined.

The `com` library must be imported into each module of a Web Services Server
application.

The following abbreviated example is from the [Web Services Server
tutorial](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. ."):

```
IMPORT com
...
FUNCTION createservice()
    DEFINE service   com.WebService    # A WebService
    DEFINE operation com.WebOperation  # Operation of a WebService
    
    # Handle HTTP register methods
    CALL service.registerInputHttpVariable(HttpIn)
    CALL service.registerOutputHttpVariable(HttpOut)

    # Publish Operation : add
    LET operation = com.WebOperation.CreateDOCStyle("service_implementation.add","add",Add,AddResponse)
    CALL service.publishOperation(operation,"")

    # Register Service
    CALL com.WebServiceEngine.RegisterService(service)
   
...
END FUNCTION
```

See the [Writing a Web server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .") and [Choosing a web services style](4673-choosing-a-web-services-style.md "Genero Web Services contains style options for creating SOAP Web services. Your choice is dependent on the type of service, (Document or RPC), and the encoding mechanism (literal or encoded) required.") for complete examples and explanations.
