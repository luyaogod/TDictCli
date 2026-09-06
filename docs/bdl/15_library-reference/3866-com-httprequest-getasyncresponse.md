---
title: "com.HttpRequest.getAsyncResponse"
source: "fgl-topics/c_gws_ComHTTPRequest_getAsyncResponse.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.getAsyncResponse"
type: "concept"
---

# com.HttpRequest.getAsyncResponse

> Retrieves an asynchronous response produced by one of the request methods.

## Syntax

```
getAsyncResponse()
  RETURNS com.HttpResponse
```

## Usage

If a response is available, the `getAsyncResponse()` method returns a [`com.HttpResponse`](3888-the-httpresponse-class.md "The com.HttpResponse class provides an interface to perform XML and TEXT responses over HTTP, with additional XML streaming possibilities, on the client side.") object corresponding to
the response that was produced by a call to one of the request methods: [`doRequest()`](3861-com-httprequest-dorequest.md "Performs the HTTP request."), [doTextRequest()](3862-com-httprequest-dotextrequest.md "Performs the request by sending an entire string at once."), [`doXmlRequest()`](3863-com-httprequest-doxmlrequest.md "Performs the request by sending an entire XML document at once."), [`doFormEncodedRequest()`](3860-com-httprequest-doformencodedrequest.md "Performs an \"application/x-www-form-urlencoded forms\" encoded query."), or [`beginXmlRequest()`](3851-com-httprequest-beginxmlrequest.md "Starts a streaming HTTP request.") and [`endXmlRequest()`](3865-com-httprequest-endxmlrequest.md "Terminates a streaming HTTP request.").

Unlike [`getResponse()`](3867-com-httprequest-getresponse.md "Waits for and returns the response produced by one of request methods."),
the `getAsyncResponse()` method retrieves an asynchronous response because it is
non-blocking and returns immediately without stopping the program flow while waiting for a
response.

The method returns `NULL` if the HTTP response has not yet been received from the
server, indicating that the request is still in process.

This method is typically called just after a `do*Request()` call, and if the
returned value is `NULL`, it is called again after a short period of time, to check
for a response. For example, within a dialog, use an `ON IDLE` block to issue a
`getAsyncRequest()` every second.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[com.HttpRequest.getResponse](3867-com-httprequest-getresponse.md "Waits for and returns the response produced by one of request methods.")

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")
