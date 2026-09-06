---
title: "xml.DomNode.removeChild"
source: "fgl-topics/c_gws_XmlDomNode_removeChild.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.removeChild"
type: "concept"
---

# xml.DomNode.removeChild

> Removes a child DomNode object from the list of child DomNode objects.

## Syntax

```
removeChild(
   oldChild xml.DomNode )
```

1. oldChild defines the node to remove.

## Usage

Use this method to remove a child [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object from a list of child `DomNode` objects, where
oldchild is the node to remove.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
