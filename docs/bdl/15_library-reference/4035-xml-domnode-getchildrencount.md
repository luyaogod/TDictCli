---
title: "xml.DomNode.getChildrenCount"
source: "fgl-topics/c_gws_XmlDomNode_getChildrenCount.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getChildrenCount"
type: "concept"
---

# xml.DomNode.getChildrenCount

> Returns the number of child DomNode objects for a DomNode object.

## Syntax

```
getChildrenCount()
  RETURNS INTEGER
```

## Usage

Returns the number of child [xml.DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") objects for this DomNode object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
