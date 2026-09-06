---
title: "com.HttpPart.CreateFromData"
source: "fgl-topics/c_gws_ComHTTPPart_CreateFromData.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.CreateFromData"
type: "concept"
---

# com.HttpPart.CreateFromData

> Creates a new HttpPart object based on given BYTE located in memory.

## Syntax

```
com.HttpPart.CreateFromData(
   b BYTE )
  RETURNS com.HttpPart
```

1. b defines a BYTE object located in
   memory.

## Usage

Creates a new HttpPart object based on given BYTE located in memory. To be
used via the [addPart()](3849-com-httprequest-addpart.md "Adds a new part to the HTTP root part request.") method.

Defaults HTTP headers:

- Content-Type: application/octet-stream
- Content-Transfer-Encoding: binary

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
