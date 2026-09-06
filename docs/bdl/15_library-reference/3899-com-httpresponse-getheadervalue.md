---
title: "com.HttpResponse.getHeaderValue"
source: "fgl-topics/c_gws_ComHTTPResponse_getHeaderValue.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getHeaderValue"
type: "concept"
---

# com.HttpResponse.getHeaderValue

> Returns the value of a header by position.

## Syntax

```
getHeaderValue(
   pos INTEGER )
  RETURNS STRING
```

1. pos
   specifies the ordinal position of the header.

## Usage

The `getHeaderValue()` method returns the value of the HTTP response header based
on the position passed as parameter.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
