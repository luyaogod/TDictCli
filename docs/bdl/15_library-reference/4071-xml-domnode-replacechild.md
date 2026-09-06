---
title: "xml.DomNode.replaceChild"
source: "fgl-topics/c_gws_XmlDomNode_replaceChild.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.replaceChild"
type: "concept"
---

# xml.DomNode.replaceChild

> Replaces an existing child DomNode with another child DomNode object.

## Syntax

```
replaceChild(
   newChild xml.DomNode,
   oldChild xml.DomNode )
```

1. newChild defines the replacement child.
2. oldChild defines the child to be replaced.

## Usage

Use this method to replace an existing child [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") with another child `DomNode` object, where
oldChild is the child to be replaced and newChild is the
replacement child.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
