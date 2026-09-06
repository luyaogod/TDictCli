---
title: "xml.DomDocument.getDocumentNodeItem"
source: "fgl-topics/c_gws_XmlDomDocument_getDocumentNodeItem.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getDocumentNodeItem"
type: "concept"
---

# xml.DomDocument.getDocumentNodeItem

> Returns the child xml.DomNode object at a given position for this xml.DomDocument object.

## Syntax

```
getDocumentNodeItem(
   index INTEGER )
  RETURNS xml.DomNode
```

1. index defines the position of the node to return (index starts at 1).

Returns an `xml.DomNode` object.

## Usage

Returns the child `xml.DomNode` object at a given position for this
`xml.DomDocument` object where index is the position of the node
to return (index starts at 1), or `NULL`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
