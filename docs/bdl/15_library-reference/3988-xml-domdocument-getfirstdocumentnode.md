---
title: "xml.DomDocument.getFirstDocumentNode"
source: "fgl-topics/c_gws_XmlDomDocument_getFirstDocumentNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getFirstDocumentNode"
type: "concept"
---

# xml.DomDocument.getFirstDocumentNode

> Returns the first child xml.DomNode object for this xml.DomDocument object.

## Syntax

```
getFirstDocumentNode()
  RETURNS xml.DomNode
```

## Usage

Use this method to return the first child `xml.DomNode` object for this
`DomDocument` object, or `NULL`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
