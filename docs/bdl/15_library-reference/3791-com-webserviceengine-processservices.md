---
title: "com.WebServiceEngine.ProcessServices"
source: "fgl-topics/c_gws_ComWebServiceEngine_ProcessServices.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.ProcessServices"
type: "concept"
---

# com.WebServiceEngine.ProcessServices

> Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services.

## Syntax

```
com.WebServiceEngine.ProcessServices(
   timeout INTEGER )
  RETURNS INTEGER
```

1. timeout defines the
   timeout in seconds.

## Usage

The `com.WebServiceEngine.ProcessServices()` method
processes requests made by client apps for operations of a SOAP or REST web service registered by
the Genero Web Service engine. Typically, you call the method inside a `WHILE` loop
to process the various requests and control the workflow for all requests made by clients. As long
as the `int_flag` variable is
`FALSE` or the execution status is valid, the condition within the
loop remains `TRUE` and requests are executed.

For a complete example of Web
service execution and status handling, see [Process requests](../16_web-services/4660-step-5-start-the-gws-server-and-process-requests.md).

The
timeout parameter defines the wait period (in seconds) the method may wait to
process a request. The value -1 specifies an infinite waiting
time.

The
status returned by the method provides information about the execution of the
last web operation. A return status of zero means OK. For a complete list of error codes, see [Error codes of com.WebServicesEngine](3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods.")

The execution status is typically handled in a
`CASE / END CASE` block, to treat all possible execution cases.

A
status of [-8](3805-error-codes-of-com-webservicesengine.md) or
[-10](3805-error-codes-of-com-webservicesengine.md) is returned if an invalid or unsupported HTTP request is sent.

Even if the exception is trapped, the GWS HTTP server is not in a
valid state anymore, and you must close the application properly before exiting the program.
> **Note:**
>
> This is not an issue in production environments as the Genero Application Server (GAS) and
> GWSProxy will detect the ended DVM, return an HTTP error code to the client app, and any new request
> will start a new DVM in a clean state via GWS proxy and the pool configuration.

> **Note:**
>
> If the [`com.WebServiceEngine.Flush()`](3787-com-webserviceengine-flush.md "Forces the Web Service engine to immediately flush the response of the web service operation.") method is used, the return status handling
> must be done in the web operation function, while
> `com.WebServiceEngine.ProcessServices()` will return the code [-31](3805-error-codes-of-com-webservicesengine.md "Error codes returned by com.WebServiceEngine methods."), to indicated that a flush was done.

## Example: Handle HTTP requests

For a complete example of how to use `com.WebServiceEngine.ProcessServices()` to
handle web service execution and status handling, go to [Process requests](../16_web-services/4660-step-5-start-the-gws-server-and-process-requests.md).
> **Tip:**
>
> When you need to intercept and respond to unregistered services or implement custom endpoints,
> use `com.WebServiceEngine.HandleRequest` instead of
> `com.WebServiceEngine.ProcessServices()`

## Related links

**Related concepts**  

[com.WebServiceEngine.HandleRequest](3790-com-webserviceengine-handlerequest.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services, or return an HttpServiceRequest object to handle a low-level request not registered at all.")

[com.WebServiceEngine.GetHTTPServiceRequest](3788-com-webserviceengine-gethttpservicerequest.md "Get a handle for an incoming HTTP service request.")
