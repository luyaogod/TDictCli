---
title: "com.WebServiceEngine.GetHTTPServiceRequest"
source: "fgl-topics/c_gws_ComWebServiceEngine_GetHTTPServiceRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.GetHTTPServiceRequest"
type: "concept"
---

# com.WebServiceEngine.GetHTTPServiceRequest

> Get a handle for an incoming HTTP service request.

## Syntax

```
com.WebServiceEngine.GetHTTPServiceRequest(
   timeout INTEGER)
  RETURNS com.HttpServiceRequest
```

1. timeout defines the
   timeout in seconds.

It returns a [`com.HttpServiceRequest`](3806-the-httpservicerequest-class.md "The com.HttpServiceRequest class provides an interface to process incoming XML and TEXT requests over HTTP on the server side, with an access to the HTTP layer and additional XML streaming possibilities.") object.

## Usage

The
`com.WebServiceEngine.GetHTTPServiceRequest()` class method returns a [`com.HttpServiceRequest`](3806-the-httpservicerequest-class.md "The com.HttpServiceRequest class provides an interface to process incoming XML and TEXT requests over HTTP on the server side, with an access to the HTTP layer and additional XML streaming possibilities.") object to
handle an incoming HTTP request, or `NULL` if no request arrives during the specified timeout.

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
- Any new call to this function will raise an error until the previous HTTP request is handled by
  sending a response back to the client, or destroyed.

**Error conditions (summary)**

- [-15565](4483-genero-bdl-errors.md): Invalid/unsupported
  HTTP request
  - Thrown for invalid or unsupported HTTP methods. Supported methods are GET, PUT, POST, HEAD, and
    DELETE.
- [-15575](4483-genero-bdl-errors.md): GAS disconnected
  the Web Services program
  - Thrown when the Genero Application Server disconnects the WS program.
- [-15552](4483-genero-bdl-errors.md): URL charset
  conversion failure
  - Thrown when the server cannot convert UTF-8 URLs to the runtime (`fglrun`) locale
    charset.
  - Recommendation: run the Web Services server program in UTF-8.

## Example

This example demonstrates using `com.WebServiceEngine.getHTTPServiceRequest(10)`
to handle incoming requests. In the `WHILE` loop, if no request arrives within 10
seconds the call returns `NULL` and the current time is output. When a request
arrives (for example, `curl localhost:8090/service_request`), it replies with 200 and
includes the request URL. The loop is wrapped in a `TRY/CATCH` to handle disconnects,
using the [status](../09_advanced-features/0939-status.md "status is a predefined variable that contains the execution status of the last instruction.") variable to report errors (`status ==
-15575`).

Key points:

- `getHTTPServiceRequest(10)`: waits up to 10 seconds (use -1 to wait
  indefinitely).
- `req IS NULL`: timeout heartbeat (prints time).
- `req.sendTextResponse(..., req.getUrl())`: sends 200 response including the
  request URL.

```
# service_request.4gl
IMPORT com

MAIN
  DEFINE req com.HttpServiceRequest
  DISPLAY "Server started"
  CALL com.WebServiceEngine.Start()
  TRY
    WHILE TRUE
      LET req = com.WebServiceEngine.getHTTPServiceRequest(10)
      IF req IS NULL THEN
        DISPLAY "HTTP request timeout...: ", CURRENT YEAR TO FRACTION
      ELSE
         DISPLAY "Request URL is :",req.getURL()
         CALL req.sendTextResponse(200,NULL,SFMT("It works (%1)", req.getUrl()))
      END IF
    END WHILE
    CATCH
      IF status == -15575 THEN
        DISPLAY "Disconnected : ",sqlca.sqlerrm
      ELSE
        DISPLAY SFMT("ERROR : %1 - %2",status,sqlca.sqlerrm)
      END IF
  END TRY
END MAIN
```

## Related links

**Related concepts**  

[com.WebServiceEngine.ProcessServices](3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services.")

[com.WebServiceEngine.HandleRequest](3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.")
