---
title: "xml.DomNode.setAttributeNS"
source: "fgl-topics/c_gws_XmlDomNode_setAttributeNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.setAttributeNS"
type: "concept"
---

# xml.DomNode.setAttributeNS

> Sets (or resets) a namespace-qualified XML Attribute for an XML Element DomNode object.

## Syntax

```
setAttributeNS(
   prefix STRING,
   name STRING,
   ns STRING,
   value STRING )
```

1. prefix defines the prefix of the XMLAttribute.
2. name defines the name of the XML Attribute.
3. ns defines the namespace URI of the XML Attribute.
4. value defines the value of the XML Attribute.

## Usage

Use this method to set (or reset) a namespace-qualified XML Attribute for this XML Element [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object, where prefix is the prefix of
the XMLAttribute, name is the name of the XML Attribute, ns is
the namespace URI of the XML Attribute, and val is the value of the XML
Attribute.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
