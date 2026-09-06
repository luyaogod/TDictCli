---
title: "xml.DomNode.insertAfterChild"
source: "fgl-topics/c_gws_XmlDomNode_insertAfterChild.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.insertAfterChild"
type: "concept"
---

# xml.DomNode.insertAfterChild

> Inserts a DomNode object after an existing child DomNode object.

## Syntax

```
insertAfterChild(
   newChild xml.DomNode,
   refChild xml.DomNode )
```

1. newChild defines the node to insert.
2. refChild defines the reference node (the
   node before which the new node must be inserted).

## Usage

Use this method to inserts a `DomNode` object after an existing child DomNode
object; newChild is the node to insert, refChild is the
reference node (the node before which the new node must be inserted).

> **Important:**
>
> This method is not part of W3C standard
> API.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
