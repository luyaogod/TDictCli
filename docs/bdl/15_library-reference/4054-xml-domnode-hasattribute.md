---
title: "xml.DomNode.hasAttribute"
source: "fgl-topics/c_gws_XmlDomNode_hasAttribute.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.hasAttribute"
type: "concept"
---

# xml.DomNode.hasAttribute

> Checks whether an XML Element DomNode object has the XML Attribute specified by a specified name.

## Syntax

```
hasAttribute(
   name STRING )
  RETURNS INTEGER
```

1. name defines the object name to check.

## Usage

Checks whether this XML Element [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object has the XML Attribute specified by name. Returns TRUE if
an XML Attribute with the given name is carried by this XML Element DomNode object, otherwise
returns FALSE.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
