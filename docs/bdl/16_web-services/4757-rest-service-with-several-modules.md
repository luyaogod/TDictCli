---
title: "Publish REST service with multiple resources"
source: "fgl-topics/c_gws_restful_high_level_publish_service_resources.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Publish a REST service > REST service with several modules"
type: "concept"
---

# Publish REST service with multiple resources

> Publish a service with several resources.

To publish a service that has many resources, in the service's main module you code to:

1. Import all the service resource modules.
2. Call on the `com.WebServiceEngine.RegisterRestResources`
   method to register the resources of the service.
3. Call on the `com.WebServiceEngine.Start` method to start
   the service.

## Web service main module

This is an example of a module that runs a Web service. The `IMPORT FGL` statement
imports the service resource modules.

In the call to the `RegisterRestResources()` method, the GWS REST engine registers
the Web service resource modules for the REST service. In the example, "fruits" is the public name
of the REST service, which users will see in the service's URI endpoints.

In the call to `ProcessServices()`, the service process starts within a
`WHILE` loop. The GWS REST engine processes requests for resources and handles [Web service engine errors](../15_library-reference/3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods."). The Web service runs
until interrupted.

```
IMPORT com
# resource modules
IMPORT FGL apple
IMPORT FGL peach

PUBLIC DEFINE serviceInfo RECORD
   title STRING,
   description STRING,
   termOfService STRING,
   contact RECORD
   name STRING,
   url STRING,
   email STRING
 END RECORD,
 version STRING
 END RECORD = (
     title: "Officestore RESTful Services.",
     version: "3.0",
     contact: ( email:"helpdesk@mysite.com") )

MAIN
  DEFINE ret INTEGER
  DEFER INTERRUPT
  CALL com.WebServiceEngine.RegisterRestResources(
    "fruits", serviceInfo, "scope", "apple", "peach")
 
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
     IF int_flag <> 0 THEN
       LET int_flag = 0
       EXIT WHILE
     END IF     
  END WHILE
  DISPLAY "Server stopped"
END MAIN
```

## Related links

**Related concepts**  

[REST resource URI naming practice](4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.")

[Publish REST service module](4756-rest-service-with-one-module.md "Publish a service with one resource.")

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")
