---
title: "Step 5: Start the GWS server and process requests"
source: "fgl-topics/c_gws_server_tutorial_009.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 1: Writing the entire server application > Step 5: Start the GWS server and process requests"
type: "concept"
---

# Step 5: Start the GWS server and process requests

> Code to start the Genero Web Services (GWS) Server.

The GWS Server is located on the same physical machine where the application is being
executed (in other words, where fglrun executes).

This is the `MAIN` program block of your application.

## Define a variable for status

Define a variable to hold the returned status of the request:

```
MAIN
  DEFINE ret INTEGER
```

Call the function that you created, which defined
and registered the service and its operations:

```
  CALL createservice()
```

## Start the GWS Server

Use the `Start` class method of the [`WebServiceEngine`](../15_library-reference/3785-the-webserviceengine-class.md "The com.WebServiceEngine class provides an interface to manage the Web Services engine.") class to start the
server.

```
  CALL com.WebServiceEngine.Start()
```

## Process requests

This example uses the `ProcessServices` method of the [`WebServiceEngine`](../15_library-reference/3785-the-webserviceengine-class.md "The com.WebServiceEngine class provides an interface to manage the Web Services engine.") class to
process each incoming request. It returns an integer representing the status. The parameter
specifies the timeout period (in seconds) the method may wait to process a service. The
value -1 specifies an infinite waiting time.

```
  WHILE TRUE
    # Process each incoming requests (infinite loop)
    LET ret = com.WebServiceEngine.ProcessServices(-1)
    CASE ret
      WHEN 0
        DISPLAY "Request processed."
      WHEN -1
        DISPLAY "Timeout reached."
      WHEN -2
        DISPLAY "Disconnected from application server."
        EXIT PROGRAM 
      WHEN -3
        DISPLAY "Client Connection lost."
      WHEN -4
        DISPLAY "Server interrupted with Ctrl-C."
      WHEN -10
        DISPLAY "Internal server error."
        EXIT PROGRAM
      WHEN -15
        DISPLAY "Server was not started."
        EXIT PROGRAM
      OTHERWISE
        DISPLAY "ERROR: ", status, sqlca.sqlerrm
    END CASE
    IF int_flag<>0 THEN
      LET int_flag=0
      EXIT WHILE
    END IF 
  END WHILE

  DISPLAY "Server stopped"

END MAIN
```

For testing purposes only, the GWS Server can be started in [standalone mode](4669-testing-the-gws-service-in-stand-alone-mode.md "Test that your service is reachable and that it can generate the WSDL."). In
a production environment, the Genero Application Server (GAS) is required to manage your
application. For deployment, the GWS Server application must be added to the GAS
configuration. See the Genero Application Server User Guide.
