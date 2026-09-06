---
title: "com.HttpRequest.setProxy"
source: "fgl-topics/c_gws_ComHTTPRequest_setProxy.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setProxy"
type: "concept"
---

# com.HttpRequest.setProxy

> Configure the proxy URL.

## Syntax

```
setProxy(
   host STRING,
   port INTEGER )
```

1. host defines the host of the proxy.
2. port defines the port number of the proxy.

## Usage

Defines the proxy URL to be used for the current `HttpRequest` request. Even if a
proxy URL is configured in FGLPROFILE, the proxy set by the `setProxy()` method will
be used.

If proxy\_host is NULL or proxy\_port is < 0, error [-15535](4483-genero-bdl-errors.md) (Invalid parameter) is
raised.

If proxy\_host is unreachable, error [-15579](4483-genero-bdl-errors.md) (COM\_PROXY\_ERROR) is raised.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
