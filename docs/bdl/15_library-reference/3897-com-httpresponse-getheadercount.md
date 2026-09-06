---
title: "com.HttpResponse.getHeaderCount"
source: "fgl-topics/c_gws_ComHTTPResponse_getHeaderCount.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getHeaderCount"
type: "concept"
---

# com.HttpResponse.getHeaderCount

> Returns the number of headers.

## Syntax

```
getHeaderCount()
  RETURNS INTEGER
```

## Usage

The `getHeaderCount()` method returns the number of headers of the HTTP
response.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
