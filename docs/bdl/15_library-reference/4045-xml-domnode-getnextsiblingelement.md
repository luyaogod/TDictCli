---
title: "xml.DomNode.getNextSiblingElement"
source: "fgl-topics/c_gws_XmlDomNode_getNextSiblingElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getNextSiblingElement"
type: "concept"
---

# xml.DomNode.getNextSiblingElement

> Returns the XML Element DomNode object immediately following a DomNode object.

## Syntax

```
getNextSiblingElement()
  RETURNS xml.DomNode
```

## Usage

Returns the XML Element `xml.DomNode` object immediately following this DomNode
object, or NULL.

> **Important:**
>
> This method is not part of W3C standard
> API.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
