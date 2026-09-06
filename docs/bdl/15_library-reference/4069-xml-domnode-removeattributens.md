---
title: "xml.DomNode.removeAttributeNS"
source: "fgl-topics/c_gws_XmlDomNode_removeAttributeNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.removeAttributeNS"
type: "concept"
---

# xml.DomNode.removeAttributeNS

> Removes a namespace qualified XML Attribute for an XML Element DomNode object

## Syntax

```
removeAttributeNS(
   name STRING,
   ns STRING )
```

1. name defines the name of the XML Attribute to remove.
2. ns defines the namespace URI of the XML Attribute to remove.

## Usage

Use this method to remove a namespace qualified XML Attribute for this XML Element [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object, where name
is the name and ns is the namespace URI of the XML Attribute to remove.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
