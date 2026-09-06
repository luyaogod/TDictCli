---
title: "Create the client application"
source: "fgl-topics/c_gws_rest_high_level_client_call_service.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Create the client application"
type: "concept"
---

# Create the client application

> Prepare the client application that uses the generated stub to call the REST service.

After generating the client stub for the REST service, the next step is to build the client application that calls the service functions. The client uses the stub to perform requests and process the responses returned by the server.

This section introduces the steps involved in creating the client application, such as writing
the main program block, accessing secure services, calling the stub functions, and handling errors
returned by the REST server.

## Related links

1. [Write the MAIN program code block](4786-write-the-main-program-code-block.md)

   Create the MAIN program block that imports the client stub and calls its functions.
2. [Call the stub functions](4789-call-the-stub-functions.md)

   Call the functions defined in the client stub to interact with the REST service from your application code.
3. [Handle GWS REST server errors](4790-handle-rest-server-errors.md)

   Handle the status code returned by a REST service call and process expected and unexpected errors in the client application.
4. [Troubleshooting no matching REST operation](4791-troubleshooting.md)

   Understanding how the GWS finds an operation to use for a request can help in troubleshooting unexpected errors.
