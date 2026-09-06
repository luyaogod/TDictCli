---
title: "com.HttpPart.CreateFromString"
source: "fgl-topics/c_gws_ComHTTPPart_CreateFromString.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.CreateFromString"
type: "concept"
---

# com.HttpPart.CreateFromString

> Creates a new HttpPart object based on given string.

## Syntax

```
com.HttpPart.CreateFromString(
   str STRING )
  RETURNS com.HttpPart
```

1. str specifies a string value.

## Usage

Creates a new HttpPart object based on given string. To be
used via the [addPart()](3849-com-httprequest-addpart.md "Adds a new part to the HTTP root part request.") method.

Defaults HTTP multipart headers:

- Content-Type: text/plain
- Content-Transfer-Encoding: 8bits

Notice that the string will be converted during request sending into ISO-8859-1 by default,
unless a different charset has been set via setHeader("Content-Type","text/plain; charset=UTF-8")
for instance.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
