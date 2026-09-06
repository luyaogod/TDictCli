---
title: "xml.DomNode.getAttributeNodeItem"
source: "fgl-topics/c_gws_XmlDomNode_getAttributeNodeItem.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getAttributeNodeItem"
type: "concept"
---

# xml.DomNode.getAttributeNodeItem

> Returns the XML Attribute DomNode object at a given position on this XML Element DomNode object.

## Syntax

```
getAttributeNodeItem(
   index INTEGER )
  RETURNS xml.DomNode
```

1. index defines the position of the node to return.

## Usage

Returns the XML Attribute `DomNode` object at a given position on this XML Element
DomNode object, where index is the position of the node to return (Index starts
at 1).

Returns the XML Attribute DomNode object at the given position, or NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
