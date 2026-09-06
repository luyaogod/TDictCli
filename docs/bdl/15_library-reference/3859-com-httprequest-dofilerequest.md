---
title: "com.HttpRequest.doFileRequest"
source: "fgl-topics/c_gws_ComHTTPRequest_doFileRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.doFileRequest"
type: "concept"
---

# com.HttpRequest.doFileRequest

> Performs the request by sending data contained in a file.

## Syntax

```
doFileRequest(
   filename STRING )
```

1. filename defines the file containing the data to be sent.

## Usage

Performs the request by sending data contained in the file passed as parameter. The data is sent
as is without any further conversion.

Supported methods are:
PUT, POST, PATCH, and DELETE.
> **Warning:**
>
> A message body is allowed in a
> DELETE request, but servers may ignore it if they do not support it.

If not defined by the programmer, the HTTP headers are automatically set as follows:

- `Content-Type` is defined based on the filename extension. If the file extension
  is not recognized, `Content-Type` defaults to
  `application/octet-stream`.
  > **Note:**
  >
  > File extensions to `Content-Type`
  > mapping can be customized in the file $FGLDIR/lib/wse/mime.cfg.
- `Content-Disposition` is set with the base name of the given
  filename as follows: `attachment;
  filename="basename"`.

For example, when calling the method as
follows:

```
CALL request.doFileRequest( "/opt/myapp/resources/logo.jpg" )
```

The resulting HTTP headers of the POST or PUT will look
like:

```
Content-Type: image/jpeg
Content-Disposition: attachment; filename="logo.jpg"
```

In HTTP 1.1, if the body size is greater than 32k, the request will be sent in several chunks of
the same size. You can disable chunk mode by setting [setBodyChunk](3872-com-httprequest-setbodychunk.md "Disable chunk mode in HTTP 1.1 if request body size is greater than 32 KB.") to `FALSE`.

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
