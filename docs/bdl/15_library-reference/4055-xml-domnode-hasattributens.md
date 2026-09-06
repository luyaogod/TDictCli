---
title: "xml.DomNode.hasAttributeNS"
source: "fgl-topics/c_gws_XmlDomNode_hasAttributeNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.hasAttributeNS"
type: "concept"
---

# xml.DomNode.hasAttributeNS

> Checks whether a namespace qualified XML Attribute of a given name is carried by an XML Element DomNode object.

## Syntax

```
hasAttributeNS(
   name STRING,
   ns STRING )
  RETURNS INTEGER
```

1. name defines the name of the XMLAttribute to check
2. ns defines the namespace URI of the XML Attribute to check.

## Usage

Use this method to check whether a namespace qualified XML Attribute of a given name is carried
by this XML Element [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object, where
name is the name of the XMLAttribute to check; ns the
namespace URI of the XML Attribute to check. Returns `TRUE` if an XML Attribute with
the given name and namespace URI is carried by this XML Element DomNode object, otherwise returns
`FALSE`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
