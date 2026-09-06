---
title: "com.HttpResponse.getHeaderName"
source: "fgl-topics/c_gws_ComHTTPResponse_getHeaderName.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getHeaderName"
type: "concept"
---

# com.HttpResponse.getHeaderName

> Returns the name of a header by position.

## Syntax

```
getHeaderName(
   pos INTEGER )
  RETURNS STRING
```

1. pos
   specifies the ordinal position of the header.

## Usage

The `getHeaderName()` method returns the name of the HTTP response header
depending on the position passed as parameter.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
