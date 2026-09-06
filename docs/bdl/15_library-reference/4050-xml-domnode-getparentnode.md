---
title: "xml.DomNode.getParentNode"
source: "fgl-topics/c_gws_XmlDomNode_getParentNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getParentNode"
type: "concept"
---

# xml.DomNode.getParentNode

> Returns the parent DomNode object for this DomNode object.

## Syntax

```
getParentNode()
  RETURNS xml.DomNode
```

## Usage

This method returns the parent `xml.DomNode` object for this
`DomNode` object, or `NULL`. In the case of a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.") node, this method will return
`NULL` (parent is not a `DomNode` object) but `isAttached()` will return
`TRUE`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
