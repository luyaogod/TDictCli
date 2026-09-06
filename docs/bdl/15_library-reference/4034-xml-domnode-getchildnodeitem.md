---
title: "xml.DomNode.getChildNodeItem"
source: "fgl-topics/c_gws_XmlDomNode_getChildNodeItem.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getChildNodeItem"
type: "concept"
---

# xml.DomNode.getChildNodeItem

> Returns the child DomNode object at a given position for a DomNode object.

## Syntax

```
getChildNodeItem(
   index INTEGER )
  RETURNS xml.DomNode
```

1. index defines the position of the child node in the collection.

## Usage

Returns the child `DomNode` object at a given position for this DomNode
object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
