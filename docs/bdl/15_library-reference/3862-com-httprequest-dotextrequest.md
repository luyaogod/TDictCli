---
title: "com.HttpRequest.doTextRequest"
source: "fgl-topics/c_gws_ComHTTPRequest_doTextRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.doTextRequest"
type: "concept"
---

# com.HttpRequest.doTextRequest

> Performs the request by sending an entire string at once.

## Syntax

```
doTextRequest(
   str STRING )
```

1. str defines a string containing the request. The request fails if
   `NULL` or an empty string ( `''` ) is provided.
   > **Tip:**
   >
   > Use [doRequest()](3861-com-httprequest-dorequest.md "Performs the HTTP request.")  for POST, PUT,
   > and PATCH requests with empty string.

## Usage

Performs the request by sending an entire string at once.

Supported methods are:
PUT, POST, PATCH, and DELETE.
> **Warning:**
>
> A message body is allowed in a
> DELETE request, but servers may ignore it if they do not support it.

The default Content-Type header is `text/plain`, but it can be changed if of the
form `*/*`. For example: `application/json`.

Automatic character set conversion from the [application
locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.") to the [user-defined charset](3874-com-httprequest-setcharset.md "Defines the charset used when sending text or XML.")
is performed. In case of conversion error, the method throws an exception.

To avoid character conversion problems when sending text over HTTP, consider setting the
same user-defined character set as the program defined by the application locale
(assuming that the server understands the client application character set).

In HTTP 1.1, if the body size is greater than 32 KB, the request will be sent in several chunks
of the same size.

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
