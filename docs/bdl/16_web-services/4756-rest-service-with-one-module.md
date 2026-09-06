---
title: "Publish REST service module"
source: "fgl-topics/c_gws_restful_high_level_publish_module.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Publish a REST service > REST service with one module"
type: "concept"
---

# Publish REST service module

> Publish a service with one resource.

To publish a service, in the service's main module you code to:

1. Import the service module.
2. Call on the `com.WebServiceEngine.RegisterRestService()` method to register the service.
3. Call on the `com.WebServiceEngine.Start` method to start
   the service.

## Web service main module

This is an example of a module that runs a Web service. The `IMPORT FGL` statement
imports the service module, "serviceModule".

In the call to the `RegisterRestService()` method, the GWS REST engine registers
the Web service module as a REST service. In the example, "MyService" is the public name of the REST
service, which users will see in the service's URI endpoints.

In the call to `ProcessServices()`, the service process starts within a
`WHILE` loop. The GWS REST engine processes requests for resources and handles [Web service engine errors](../15_library-reference/3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods."). The Web service runs
until interrupted.

```
IMPORT com
IMPORT FGL serviceModule

MAIN
  DEFINE ret INTEGER
  DEFER INTERRUPT
  CALL com.WebServiceEngine.RegisterRestService("serviceModule", "MyService")
  DISPLAY "Server started"
  CALL com.WebServiceEngine.Start()
  CONNECT TO "myDatabase""+driver='dbmsqt'"
  WHILE TRUE
    LET ret = com.WebServiceEngine.ProcessServices(-1)
    CASE ret
       WHEN 0
         DISPLAY "Request processed." 
       WHEN -1
         DISPLAY "Timeout reached."
       WHEN -2
         DISPLAY "Disconnected from application server."
         EXIT PROGRAM   # The Application server has closed the connection
       WHEN -3
         DISPLAY "Client Connection lost."
       WHEN -4
         DISPLAY "Server interrupted with Ctrl-C."
       WHEN -9
         DISPLAY "Unsupported operation."
       WHEN -10
         DISPLAY "Internal server error."
       WHEN -23
         DISPLAY "Deserialization error."
       WHEN -35
         DISPLAY "No such REST operation found."
       WHEN -36
         DISPLAY "Missing REST parameter."
       OTHERWISE 
         DISPLAY "Unexpected server error " || ret || "."
        EXIT WHILE 
     END CASE
     IF int_flag<>0 THEN
       LET int_flag=0
       EXIT WHILE
     END IF     
  END WHILE
  DISPLAY "Server stopped"
END MAIN
```

## Related links

**Related concepts**  

[Publish REST service with multiple resources](4757-rest-service-with-several-modules.md "Publish a service with several resources.")

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")
