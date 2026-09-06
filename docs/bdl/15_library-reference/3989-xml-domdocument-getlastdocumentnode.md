---
title: "xml.DomDocument.getLastDocumentNode"
source: "fgl-topics/c_gws_XmlDomDocument_getLastDocumentNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getLastDocumentNode"
type: "concept"
---

# xml.DomDocument.getLastDocumentNode

> Returns the last child xml.DomNode object in this xml.DomDocument object.

## Syntax

```
getLastDocumentNode()
  RETURNS xml.DomNode
```

## Usage

Use this method to return the last child `xml.DomNode` object in this
`xml.DomDocument` object, or `NULL`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
