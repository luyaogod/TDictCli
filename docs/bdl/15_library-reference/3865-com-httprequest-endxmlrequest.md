---
title: "com.HttpRequest.endXmlRequest"
source: "fgl-topics/c_gws_ComHTTPRequest_endXmlRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.endXmlRequest"
type: "concept"
---

# com.HttpRequest.endXmlRequest

> Terminates a streaming HTTP request.

## Syntax

```
endXmlRequest(
   stax xml.StaxWriter )
```

1. stax defines the `xml.StaxWriter` object used to write the
   HTTP request.

## Usage

The `endXmlRequest()` method terminates a streaming HTTP request by closing the
`xml.StaxWriter` object that was created with the [`beginXmlRequest()`](3851-com-httprequest-beginxmlrequest.md "Starts a streaming HTTP request.")
method.

This HTTP request method is non-blocking. It returns immediately
after the call. Use the [com.HttpRequest.getResponse](3867-com-httprequest-getresponse.md "Waits for and returns the response produced by one of request methods.") method, to perform a
synchronous HTTP request, suspending the program flow until the response returns from the server. If
the program must keep going, use the [com.HttpRequest.getAsyncResponse](3866-com-httprequest-getasyncresponse.md "Retrieves an asynchronous response produced by one of the request methods.")
method, to check if a response is available.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
