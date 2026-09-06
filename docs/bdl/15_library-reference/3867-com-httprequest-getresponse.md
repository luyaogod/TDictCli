---
title: "com.HttpRequest.getResponse"
source: "fgl-topics/c_gws_ComHTTPRequest_getResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.getResponse"
type: "concept"
---

# com.HttpRequest.getResponse

> Waits for and returns the response produced by one of request methods.

## Syntax

```
getResponse()
  RETURNS com.HttpResponse
```

## Usage

The `getResponse()` method waits for a response from the server and returns a
[`com.HttpResponse`](3888-the-httpresponse-class.md "The com.HttpResponse class provides an interface to perform XML and TEXT responses over HTTP, with additional XML streaming possibilities, on the client side.") object
corresponding to the response that was produced by a call to one of the request methods: [`doRequest()`](3861-com-httprequest-dorequest.md "Performs the HTTP request."), [`doTextRequest()`](3862-com-httprequest-dotextrequest.md "Performs the request by sending an entire string at once."), [`doXmlRequest()`](3863-com-httprequest-doxmlrequest.md "Performs the request by sending an entire XML document at once."), [`doFormEncodedRequest()`](3860-com-httprequest-doformencodedrequest.md "Performs an \"application/x-www-form-urlencoded forms\" encoded query."), or [`beginXmlRequest()`](3851-com-httprequest-beginxmlrequest.md "Starts a streaming HTTP request.") and [`endXmlRequest()`](3865-com-httprequest-endxmlrequest.md "Terminates a streaming HTTP request.").

> **Note:**
>
> On iOS, a long running HTTP request will display a message box, to allow the user to cancel
> the request. If the user cancels the HTTP request, the error [-15578](4483-genero-bdl-errors.md) will be raised. This error
> can be trapped with `TRY/CATCH`.

Unlike [`getAsyncResponse()`](3866-com-httprequest-getasyncresponse.md "Retrieves an asynchronous response produced by one of the request methods."), the `getResponse()` method is blocking;
it stops program flow until an HTTP response is received from the server.

Define a response timeout with the [com.HttpRequest.setTimeOut](3884-com-httprequest-settimeout.md "Defines the timeout for a reading or writing operation.")
method.

> **Note:**
>
> On iOS devices, when using this method, it is not possible to distinguish different timeouts
> for the connection and for read/write operation, defined respectively by the [`setConnectionTimeOut()`](3876-com-httprequest-setconnectiontimeout.md "Defines the timeout for the establishment of the connection.")
> and [`setTimeOut()`](3884-com-httprequest-settimeout.md "Defines the timeout for a reading or writing operation.") methods.
> If both timeouts are defined, the longest timeout will be used for the connection and read/write
> operations.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[com.HttpRequest.getAsyncResponse](3866-com-httprequest-getasyncresponse.md "Retrieves an asynchronous response produced by one of the request methods.")

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")
