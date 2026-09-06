---
title: "com.HttpPart.getContentAsDomDocument"
source: "fgl-topics/c_gws_ComHTTPPart_getContentAsDomDocument.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.getContentAsDomDocument"
type: "concept"
---

# com.HttpPart.getContentAsDomDocument

> Returns the HTTP part as an XML document.

## Syntax

```
getContentAsDomDocument()
  RETURNS xml.DomDocument
```

## Usage

Returns the HTTP part as an XML document.

To be used via methods: [com.HttpResponse.getPart](3901-com-httpresponse-getpart.md "Returns the HTTP part object at the specified index of the current HTTP response."),
[com.HttpResponse.getPartCount](3902-com-httpresponse-getpartcount.md "Returns the number of additional parts in the HTTP response."), and [com.HttpResponse.getPartFromID](3903-com-httpresponse-getpartfromid.md "Returns the HTTP part object marked with the given Content-ID value as identifier, or NULL if none.")

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The [error
-15573](4483-genero-bdl-errors.md) is raised if the part cannot be converted to a XML DomDocument.
