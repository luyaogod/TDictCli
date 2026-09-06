---
title: "xml.DomNode.addNextSibling"
source: "fgl-topics/c_gws_XmlDomNode_addNextSibling.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.addNextSibling"
type: "concept"
---

# xml.DomNode.addNextSibling

> Adds a DomNode object as the next sibling of a DomNode object.

## Syntax

```
addNextSibling(
  newNode xml.DomNode )
```

1. newNode defines the node to add.

## Usage

Adds a `DomNode` object as the next sibling of this DomNode object;
newNode is the node to add.

> **Important:**
>
> This method is not part of W3C standard
> API.

The DomNode object node must be the child
of an element or document node; otherwise the operation fails.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
