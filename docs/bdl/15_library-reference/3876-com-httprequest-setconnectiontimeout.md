---
title: "com.HttpRequest.setConnectionTimeOut"
source: "fgl-topics/c_gws_ComHTTPRequest_setConnectionTimeOut.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setConnectionTimeOut"
type: "concept"
---

# com.HttpRequest.setConnectionTimeOut

> Defines the timeout for the establishment of the connection.

## Syntax

```
setConnectionTimeOut(
   timeout INTEGER )
```

1. timeout defines the
   number of seconds.

## Usage

The `setConnectionTimeOut()` method sets the time value in seconds to wait for the
establishment of the connection, before a break.

The value of -1 means infinite wait.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
