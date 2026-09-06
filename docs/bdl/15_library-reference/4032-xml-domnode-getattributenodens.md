---
title: "xml.DomNode.getAttributeNodeNS"
source: "fgl-topics/c_gws_XmlDomNode_getAttributeNodeNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getAttributeNodeNS"
type: "concept"
---

# xml.DomNode.getAttributeNodeNS

> Returns a namespace-qualified XML Attribute DomNode object for an XML Element DomNode object

## Syntax

```
getAttributeNodeNS(
   name STRING,
   ns STRING )
  RETURNS xml.DomNode
```

1. name defines the name of the XMLAttribute to retrieve.
2. ns defines the namespace URI of the XML Attribute to retrieve.

## Usage

Returns a namespace-qualified XML Attribute `DomNode` object for this XML Element
DomNode object, or NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
