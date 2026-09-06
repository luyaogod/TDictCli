---
title: "com.HttpRequest.doDataRequest"
source: "fgl-topics/c_gws_ComHTTPRequest_doDataRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.doDataRequest"
type: "concept"
---

# com.HttpRequest.doDataRequest

> Performs the request by sending binary data.

## Syntax

```
doDataRequest(
   b BYTE )
```

1. b defines the binary data.

## Usage

Performs the request by sending binary data contained in the `b` variable.

Supported methods are:
PUT, POST, PATCH, and DELETE.
> **Warning:**
>
> A message body is allowed in a
> DELETE request, but servers may ignore it if they do not support it.

The `b` must be located in memory and not `NULL` otherwise the
operation fails.

The default Content-Type header is `application/octet-stream`, but it can be
changed to any other mime type. For example: `image/jpeg`.

In HTTP 1.1, if the body size is greater than 32k, the request will be sent in several chunks of
the same size.

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
