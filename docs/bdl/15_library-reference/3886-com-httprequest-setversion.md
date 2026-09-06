---
title: "com.HttpRequest.setVersion"
source: "fgl-topics/c_gws_ComHTTPRequest_setVersion.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setVersion"
type: "concept"
---

# com.HttpRequest.setVersion

> Sets the HTTP version of the request.

## Syntax

```
setVersion(
   version STRING )
```

1. version defines the HTTP version of the request.

## Usage

The `setVersion()` method defines the HTTP version of the request.

Accepted versions are 1.0 and 1.1 (only these two versions are supported).

The default is 1.1.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
