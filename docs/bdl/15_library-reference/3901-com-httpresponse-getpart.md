---
title: "com.HttpResponse.getPart"
source: "fgl-topics/c_gws_ComHTTPResponse_getPart.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getPart"
type: "concept"
---

# com.HttpResponse.getPart

> Returns the HTTP part object at the specified index of the current HTTP response.

## Syntax

```
getPart(
   pos INTEGER )
  RETURNS com.HttpPart
```

1. pos is the input parameter that specifies the index number.

## Usage

Returns the HTTP part object at the specified index of the current HTTP response.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

Can raise error [-15554](4483-genero-bdl-errors.md)
(Index is out of bounds).

## Related links

**Related concepts**  

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
