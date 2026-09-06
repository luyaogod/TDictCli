---
title: "xml.DomNode.appendChildElement"
source: "fgl-topics/c_gws_XmlDomNode_appendChildElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.appendChildElement"
type: "concept"
---

# xml.DomNode.appendChildElement

> Creates and adds a child XML Element node to the end of the list of child nodes for an XML Element DomNode object.

## Syntax

```
appendChildElement(
   name STRING )
  RETURNS xml.DomNode
```

1. name defines the XML Element name.

## Usage

Creates and adds a child XML Element node to the end of the list of child nodes for this XML
Element `DomNode` object.

> **Important:**
>
> This method is not part of W3C standard
> API.

Returns the XML element `xml.DomNode` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
