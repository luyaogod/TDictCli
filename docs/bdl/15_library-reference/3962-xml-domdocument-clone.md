---
title: "xml.DomDocument.clone"
source: "fgl-topics/c_gws_XmlDomDocument_clone.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.clone"
type: "concept"
---

# xml.DomDocument.clone

> Returns a copy of this xml.DomDocument object.

## Syntax

```
clone()
  RETURNS xml.DomDocument
```

## Usage

Returns a copy of this `xml.DomDocument` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
