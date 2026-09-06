---
title: "xml.DomNode.addPreviousSibling"
source: "fgl-topics/c_gws_XmlDomNode_addPreviousSibling.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.addPreviousSibling"
type: "concept"
---

# xml.DomNode.addPreviousSibling

> Adds a DomNode object as the previous sibling of a DomNode object.

## Syntax

```
addPreviousSibling(
   newNode xml.DomNode )
```

1. newNode defines the node to add.

## Usage

Adds a `DomNode` object as the previous sibling of this [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object; newNode is
the node to add.

> **Important:**
>
> This method is not part of W3C standard
> API.

The DomNode object node must be the child
of an element or document node; otherwise the operation fails.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
