---
title: "xml.DomNode.setAttributeNodeNS"
source: "fgl-topics/c_gws_XmlDomNode_setAttributeNodeNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.setAttributeNodeNS"
type: "concept"
---

# xml.DomNode.setAttributeNodeNS

> Sets (or resets) a namespace-qualified XML Attribute DomNode object to an XML Element DomNode object.

## Syntax

```
setAttributeNodeNS(
   attr xml.DomNode )
```

1. attr defines the XML Attribute DomNode object to set.

## Usage

Use this method to set (or reset) a namespace-qualified XML Attribute `DomNode`
object to an XML Element `DomNode` object, where attr is the XML
Attribute `DomNode` object to set.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
