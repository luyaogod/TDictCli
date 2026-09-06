---
title: "com.WebServiceEngine.HandleRequest"
source: "fgl-topics/c_gws_ComWebServiceEngine_HandleRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.HandleRequest"
type: "concept"
---

# com.WebServiceEngine.HandleRequest

> Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.

## Syntax

```
com.WebServiceEngine.HandleRequest(
   timeout INTEGER,
   status INTEGER )
  RETURNS com.HttpServiceRequest
```

1. timeout defines the
   timeout in seconds.
2. status defines the variable receiving the method execution status. The
   variable must be `INTEGER`, `SMALLINT`, or `BIGINT`;
   otherwise, a runtime error is thrown.

The method returns a `com.HttpServiceRequest` object.

## Usage

`com.WebServiceEngine.HandleRequest()` waits for incoming HTTP requests and
behaves as follows:

- **Registered service**: If the request targets a registered SOAP or REST service, the engine
  processes it automatically (similar to `ProcessServices()`) and returns
  `NULL` with `status = 0`.
- **Unregistered service**: If the request targets a service that is not registered, the method
  returns a `status = 1` and a valid instance of an [`com.HttpServiceRequest`](3806-the-httpservicerequest-class.md "The com.HttpServiceRequest class provides an interface to process incoming XML and TEXT requests over HTTP on the server side, with an access to the HTTP layer and additional XML streaming possibilities.") object,
  immediately usable to handle the incoming request, is returned.
  > **Note:**
  >
  > This exception applies for query-only requests targeting the service root path (for example
  > `/service?abc`) which may fall back to low-level handling when no REST operation
  > matches.

`NULL` is returned if no request arrives during the specified timeout.

- timeout seconds to wait for a request. Use -1 for infinite wait. Returns
  `NULL` on timeout.
- In case of error, the method throws an exception and sets the
  `status` variable. Depending on the error, a human-readable description of the
  problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

  Even if the exception is trapped, the GWS HTTP server is not in a
  valid state anymore, and you must close the application properly before exiting the program.
  > **Note:**
  >
  > This is not an issue in production environments as the Genero Application Server (GAS) and
  > GWSProxy will detect the ended DVM, return an HTTP error code to the client app, and any new request
  > will start a new DVM in a clean state via GWS proxy and the pool configuration.

Before using this method, you register the service and start the engine the same as when using
[ProcessServices()](3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services."). For example,
call [RegisterRestService()](3795-com-webserviceengine-registerservice.md "Registers a SOAP service in the engine.") or
[RegisterService()](3795-com-webserviceengine-registerservice.md "Registers a SOAP service in the engine.")  before
starting the engine.

You typically call `com.WebServiceEngine.HandleRequest()` inside a loop that runs
until the server stops or a fatal error occurs. As long as the `int_flag` variable is `FALSE` or
the execution status is valid, the condition within the loop remains
`TRUE` and requests are executed.

The execution status is typically handled in a `CASE / END
CASE` block to manage all possible cases. The status returned provides
information about the execution of the last operation:

| Status value | Meaning | Notes |
| --- | --- | --- |
| 0 | OK | Normal successful result. |
| Negative values | Error | For example, [-8](3805-error-codes-of-com-webservicesengine.md) or [-10](3805-error-codes-of-com-webservicesengine.md) are returned for an invalid or unsupported HTTP method. Supported methods are `GET`, `PUT`, `POST`, `HEAD`, `DELETE`, `OPTIONS`, `TRACE`, and `PATCH`. For the full list of error codes, go to [Error codes of com.WebServicesEngine](3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods.") |
| 1 | Request for an unregistered service | The engine received a request for a service name your application has not registered. You may send a response manually using the [`com.HttpServiceRequest`](3806-the-httpservicerequest-class.md "The com.HttpServiceRequest class provides an interface to process incoming XML and TEXT requests over HTTP on the server side, with an access to the HTTP layer and additional XML streaming possibilities.") object. If you don’t respond, the next call to `HandleRequest()` will fail because the previous request object remains active. |

> **Important:**
>
> **Special case**: If the [com.WebServiceEngine.Flush()](3787-com-webserviceengine-flush.md "Forces the Web Service engine to immediately flush the response of the web service operation.") method is used, the return status handling must be done in the
> operation function, while `com.WebServiceEngine.HandleRequest()` will return the code
> [-31](3805-error-codes-of-com-webservicesengine.md).

| Error code | Meaning | Notes / Recommendation |
| --- | --- | --- |
| [-15552](4483-genero-bdl-errors.md) | URL charset conversion failure | Thrown when the server cannot convert UTF‑8 URLs to the fglrun locale charset. **Recommendation:** run the Web Services server program in UTF‑8. |
| [-15575](4483-genero-bdl-errors.md) | GAS disconnected the Web Services program | Thrown when the Genero Application Server disconnects the WS program. |

## Example: Handle HTTP requests

This example shows how to use `com.WebServiceEngine.HandleRequest()` to process
incoming requests and handle `status = 1` by responding to requests for unregistered
services.
> **Tip:**
>
> Use `com.WebServiceEngine.ProcessServices()` for servers that handle only registered web
> services. Use `com.WebServiceEngine.HandleRequest()` when you need to intercept and
> respond to unregistered services or implement custom endpoints.

```
IMPORT com
IMPORT FGL myservice

MAIN
  DEFINE status INTEGER
  DEFINE req com.HttpServiceRequest

  CALL com.WebServiceEngine.RegisterRestService("myservice", "MyService")
  DISPLAY "Server started"
  CALL com.WebServiceEngine.Start()

  WHILE TRUE
    LET req = com.WebServiceEngine.HandleRequest(-1, status)

    CASE status
      WHEN 0
        DISPLAY "Request processed."
      WHEN 1
        DISPLAY "Unregistered service request."
        CALL req.sendResponse(400, SFMT("Unknown service (%1)", req.getUrl()))
      WHEN -1
        DISPLAY "Timeout reached."
      WHEN -2
        DISPLAY "Disconnected from application server."
        EXIT PROGRAM
      WHEN -3
        DISPLAY "Client connection lost."
      WHEN -4
        DISPLAY "Server interrupted with Ctrl-C."
      WHEN -8
        DISPLAY "Invalid HTTP request."
      WHEN -10
        DISPLAY "Unsupported HTTP request or verb."
      WHEN -23
        DISPLAY "Deserialization error."
      WHEN -31
        DISPLAY "Flush detected."
      WHEN -35
        DISPLAY "No such REST operation found."
      WHEN -36
        DISPLAY "Missing REST parameter."
      WHEN -15552
        DISPLAY "UTF-8 URL conversion error."
      WHEN -15575
        DISPLAY "GAS disconnected the Web Services program."
      OTHERWISE
        DISPLAY "Unexpected server error ", status, " - ", sqlca.sqlerrm
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

[com.WebServiceEngine.ProcessServices](3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services.")

[com.WebServiceEngine.GetHTTPServiceRequest](3788-com-webserviceengine-gethttpservicerequest.md "Get a handle for an incoming HTTP service request.")
