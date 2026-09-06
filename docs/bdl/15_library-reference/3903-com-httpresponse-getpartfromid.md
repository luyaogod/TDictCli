---
title: "com.HttpResponse.getPartFromID"
source: "fgl-topics/c_gws_ComHTTPResponse_getPartFromID.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getPartFromID"
type: "concept"
---

# com.HttpResponse.getPartFromID

> Returns the HTTP part object marked with the given Content-ID value as identifier, or NULL if none.

## Syntax

```
getPartFromID(
   id STRING )
  RETURNS com.HttpPart
```

1. id defines the HTTP header.

## Usage

Returns the HTTP part object marked with the given Content-ID value as identifier, or NULL if
none.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
