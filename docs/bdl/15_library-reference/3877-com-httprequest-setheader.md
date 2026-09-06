---
title: "com.HttpRequest.setHeader"
source: "fgl-topics/c_gws_ComHTTPRequest_setHeader.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setHeader"
type: "concept"
---

# com.HttpRequest.setHeader

> Sets an HTTP header for the request.

## Syntax

```
setHeader(
   name STRING,
   value STRING )
```

1. name defines the HTTP
   header name.
2. value defines the HTTP header value.

## Usage

The `setHeader()` method defines an HTTP header with a name and
value for the request.

If a header exists with the same name, it is replaced with the new value.

Setting a header after the body has been sent, or if a streaming operation has been started, will
not take effect. It will only be set when a new request is reissued.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
