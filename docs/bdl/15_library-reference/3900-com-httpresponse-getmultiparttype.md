---
title: "com.HttpResponse.getMultipartType"
source: "fgl-topics/c_gws_ComHTTPResponse_getMultipartType.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getMultipartType"
type: "concept"
---

# com.HttpResponse.getMultipartType

> Returns whether a response is multipart or not, and the kind of multipart if any.

## Syntax

```
getMultipartType()
  RETURNS STRING
```

## Usage

Returns whether a response is multipart or not, and the kind of multipart if any.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
