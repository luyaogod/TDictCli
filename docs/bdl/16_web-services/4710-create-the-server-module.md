---
title: "Create the server module"
source: "fgl-topics/t_gws_restful_high_level_quick_start_main_module.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 1: RESTful server application > Create the server module"
type: "task"
---

# Create the server module

> The server module registers the Web Service application with the Genero Web Services (GWS) server that starts the Web service.

In this task you create a module that:

1. Imports our [service
   module](4709-create-a-web-service-module.md "Resources are identified as a set of functions within the Web services module that provides the service.").
2. Calls the `com.WebServiceEngine.RegisterRestService()` method to register the service with the
   Web server.
3. Calls the `com.WebServiceEngine.ProcessServices()` method to start the service.

## Steps

1. Create your application server module. For example, create a file named
   wsserver.4gl
2. Add the code from the example.

   In this example, you have the code you need. This code acts as a template for starting your Web
   services. In the `CONNECT TO` statement, we add the name of the database created in
   [Prepare a database](4708-prepare-a-database.md "In preparation for working with the Web service, create a database with data from the code provided and extract the schema file.") and we use the SQLite driver. For
   more information on what the code does, see the [Publish REST service module](4756-rest-service-with-one-module.md "Publish a service with one resource.") page.

   ```
   IMPORT com
   IMPORT FGL myservice

   MAIN
     DEFINE ret INTEGER
     CALL com.WebServiceEngine.RegisterRestService("myservice", "MyService")
     DISPLAY "Server started"
     CALL com.WebServiceEngine.Start()
     CONNECT TO "test1.db+driver='dbmsqt'"
     WHILE TRUE
       LET ret = com.WebServiceEngine.ProcessServices(-1)
       CASE ret
         WHEN 0
           DISPLAY "Request processed."
         WHEN -1
           DISPLAY "Timeout reached."
         WHEN -2
           DISPLAY "Disconnected from application server."
           EXIT PROGRAM # The Application server has closed the connection
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

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")
