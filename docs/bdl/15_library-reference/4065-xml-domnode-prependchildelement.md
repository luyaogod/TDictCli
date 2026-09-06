---
title: "xml.DomNode.prependChildElement"
source: "fgl-topics/c_gws_XmlDomNode_prependChildElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.prependChildElement"
type: "concept"
---

# xml.DomNode.prependChildElement

> Creates and adds a child XML Element node to the beginning of the list of child nodes for this XML Element DomNode object.

## Syntax

```
prependChildElement(
   name STRING )
  RETURNS xml.DomNode
```

1. name defines the name of the XML element to add.

## Usage

Use this method to create and add a child XML Element node to the beginning of the list of child
nodes for this XML Element `DomNode` object; name is the name of
the XML element to add.

It returns the XML Element `DomNode` object, or NULL.

> **Important:**
>
> This method is not part of W3C standard
> API.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
