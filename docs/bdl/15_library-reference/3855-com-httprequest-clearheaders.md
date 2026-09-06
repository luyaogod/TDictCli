---
title: "com.HttpRequest.clearHeaders"
source: "fgl-topics/c_gws_ComHTTPRequest_clearHeaders.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.clearHeaders"
type: "concept"
---

# com.HttpRequest.clearHeaders

> Removes all user-defined HTTP request headers.

## Syntax

```
clearHeaders()
```

## Usage

Removes all user-defined HTTP request headers defined with the [`setHeader()`](3877-com-httprequest-setheader.md "Sets an HTTP header for the request.") method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
