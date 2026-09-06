---
title: "Step 2: Listen for requests"
source: "fgl-topics/c_gws_rest_server_listen_request.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web Services server application > Step 2: Listen for requests"
type: "concept"
---

# Step 2: Listen for requests

> Listen for an incoming request.

Define an object variable using the `com.HttpServiceRequest` class identifier.
This variable will reference the incoming HTTP service request. In addition, define an
`INTEGER` variable to hold an error status, if
needed.

```
DEFINE ret INTEGER
DEFINE req com.HttpServiceRequest
```

Start the server engine with the method
`com.WebServiceEngine.Start()`.

```
DEFER INTERRUPT

# Start the server

DISPLAY "Starting server..."

# Starts the server on the port number specified by the FGLAPPSERVER environment variable
#  (EX: FGLAPPSERVER=8090)
 
CALL com.WebServiceEngine.Start()
DISPLAY "The server is listening."
```

Listen for an incoming request using `com.WebServiceEngine.getHTTPServiceRequest`.
When a request is received, the variable `req` references the HTTP service request
object. You can use the same object to send a response.

Include a `CATCH` block to trap runtime exceptions. For a RESTful Web service,
include a check for error [-15565](../15_library-reference/4483-genero-bdl-errors.md), which is
returned when the GWS server is started in development mode and an incoming request
cannot be returned. Include a test also for error [-15575](../15_library-reference/4483-genero-bdl-errors.md) for when the
server will be behind a GAS in production and an incoming request cannot be
returned.

```
...
# create the server
WHILE TRUE
    TRY
        LET req = com.WebServiceEngine.getHTTPServiceRequest(-1)
        ... parse and process the request ...
    CATCH
         LET ret = status
         CASE ret
         WHEN -15565
             DISPLAY "Server process has been stopped by Ctrl-C"
         WHEN -15575
            DISPLAY "Server process has been disconnected by the GAS"
         OTHERWISE
             DISPLAY "[ERROR] " || ret
             EXIT WHILE
         END CASE
         IF int_flag<>0 THEN
              LET int_flag = 0
              EXIT WHILE
         END IF
    END TRY
END WHILE
```

For more information see [com.WebServiceEngine methods](../15_library-reference/3786-webserviceengine-methods.md "Methods for the com.WebServiceEngine class.").

In the next step we extract the details of the incoming request, [Step 3: Parse the request](4872-step-3-parse-the-request.md "Using methods of the com.HttpServiceRequest class (among others), parse the details of the request.")
