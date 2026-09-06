---
title: "Set up and test your environment"
source: "fgl-topics/t_fgl_set_up_restws_highapi.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Set up and test your environment"
type: "task"
description: "Complete these steps to set up your system to execute a REST Web Services server when in development. Set the runtime environment. Run the script file envcomp in the Genero BDL installation directory ..."
---

# Set up and test your environment

Complete these steps to set up your system to execute a REST Web Services server when in
development.

1. Set the runtime environment.

   Run the script file envcomp in the Genero BDL installation directory to
   ensure that FGLDIR and PATH are set correctly. They are required to run compiler and runtime system
   tools.
2. Set FGLAPPSERVER to 8090. 

   This will start the web service server application on port 8090. If this environment variable
   is not set, port number 80 is used.

   > **Warning:**
   >
   > Do not set the FGLAPPSERVER variable in production environments. In
   > production, the Genero Application Server selects the port number.
3. Set FGLWSDEBUG to 3.

   This ensures there is detail in the output of the HTTP requests and responses showing the
   interaction between the server and its clients. It is useful for debugging purposes. See [FGLWSDEBUG](../07_configuration/0537-fglwsdebug.md "The FGLWSDEBUG environment variable enables web services library debugging.").
4. Create an application module.

   Add service functions using the GWS REST function attributes. See the examples in the [Define functions in a module](4723-define-functions-in-a-module.md "A GWS REST service is defined in a module.") section.

   To perform an initial test, create a file named myService.4gl. For an
   example, see [Quick start 1: RESTful server application](4707-quick-start-1-restful-server-application.md "This quick start provides step-by-step instruction for creating a RESTful Web service server application using the high-level framework. The application will manage access to customers stored in a database.").
5. Create the application's main service module.

   To perform an initial test, create a file named wsserver.4gl to contain
   your main module. For a code sample for starting Web services, see [Quick start 1: RESTful server application](4707-quick-start-1-restful-server-application.md "This quick start provides step-by-step instruction for creating a RESTful Web service server application using the high-level framework. The application will manage access to customers stored in a database.").
6. Compile the application.

   To continue the initial test, execute :fglcomp wsserver myService.
7. Start the Genero RESTful Web Service server.

   Use the fglrun command to execute the web service server application. The
   web service server starts on the port specified by FGLAPPSERVER.

   ```
   fglrun <wsserverapp>
   ```
8. Retrieve the OpenAPI description.

   For a web service server started with FGLAPPSERVER=8090 (as specified in 2) the OpenAPI description
   URL would be:

   http://localhost:8090/myService?openapi.json.
