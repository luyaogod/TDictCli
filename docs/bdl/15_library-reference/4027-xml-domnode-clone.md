---
title: "xml.DomNode.clone"
source: "fgl-topics/c_gws_XmlDomNode_clone.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.clone"
type: "concept"
---

# xml.DomNode.clone

> Returns a duplicate DomNode object of a node.

## Syntax

```
clone(
   deep INTEGER )
  RETURNS xml.DomNode
```

1. deep defines a boolean. If
   deep is TRUE, child DomNode objects are cloned too; otherwise only the
   DomNode itself is cloned.

## Usage

Returns a duplicate DomNode object of this node. If deep is TRUE,
child DomNode objects are cloned too; otherwise only the DomNode itself is cloned.

Returns a copy of this `xml.DomNode` object, or NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
