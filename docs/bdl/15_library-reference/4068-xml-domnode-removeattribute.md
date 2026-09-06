---
title: "xml.DomNode.removeAttribute"
source: "fgl-topics/c_gws_XmlDomNode_removeAttribute.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.removeAttribute"
type: "concept"
---

# xml.DomNode.removeAttribute

> Removes an XML Attribute for an XML Element DomNode object.

## Syntax

```
removeAttribute(
  name STRING )
```

1. name defines the name of the XML attribute to remove.

## Usage

Use this method to remove an XML Attribute for this XML Element [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object, where name
is the name of the XML attribute to remove. Status is updated with an [error code](4483-genero-bdl-errors.md "System error messages sorted by error number.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
