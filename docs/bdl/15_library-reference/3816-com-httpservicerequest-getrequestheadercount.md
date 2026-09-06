---
title: "com.HttpServiceRequest.getRequestHeaderCount"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_getRequestHeaderCount.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.getRequestHeaderCount"
type: "concept"
---

# com.HttpServiceRequest.getRequestHeaderCount

> Returns the number of request headers.

## Syntax

```
getRequestHeaderCount()
  RETURNS INTEGER
```

## Usage

The `getRequestHeaderCount()` method returns the entire URL request containing the
host, port, document and query string.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
