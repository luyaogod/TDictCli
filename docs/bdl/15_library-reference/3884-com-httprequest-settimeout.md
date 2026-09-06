---
title: "com.HttpRequest.setTimeOut"
source: "fgl-topics/c_gws_ComHTTPRequest_setTimeOut.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setTimeOut"
type: "concept"
---

# com.HttpRequest.setTimeOut

> Defines the timeout for a reading or writing operation.

## Syntax

```
setTimeOut(
   timeout INTEGER )
```

1. timeout defines the
   number of seconds.

## Usage

The `setTimeOut()` method defines a delay in seconds, to wait for an HTTP request
read or write operation. If the operation is not terminated after the timeout, it returns
immediately with an error.

Use the value of -1 to define an infinite timeout.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
