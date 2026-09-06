---
title: "com.HttpPart.CreateFromDomDocument"
source: "fgl-topics/c_gws_ComHTTPPart_CreateFromDomDocument.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.CreateFromDomDocument"
type: "concept"
---

# com.HttpPart.CreateFromDomDocument

> Creates a new HttpPart object based on given XML document.

## Syntax

```
com.HttpPart.CreateFromDomDocument(
   doc xml.DomDocument )
  RETURNS com.HttpPart
```

1. doc specifies an XML document.

## Usage

Creates a new HttpPart object based on given XML document. To be
used via the [addPart()](3849-com-httprequest-addpart.md "Adds a new part to the HTTP root part request.") method.

Defaults HTTP multipart headers:

- Content-Type: text/xml; charset=UTF-8
- Content-Transfer-Encoding: 8bits

A different charset can be set with the setHeader method. For example,
`setHeader("Content-Type","text/plain; charset=ISO-8859-1")` sets the charset to
ISO-8859-1.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[com.HttpPart.setHeader](3926-com-httppart-setheader.md "Sets a named HTTP multipart header using a string value.")
