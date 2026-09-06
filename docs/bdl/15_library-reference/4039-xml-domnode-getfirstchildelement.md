---
title: "xml.DomNode.getFirstChildElement"
source: "fgl-topics/c_gws_XmlDomNode_getFirstChildElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getFirstChildElement"
type: "concept"
---

# xml.DomNode.getFirstChildElement

> Returns the first XML Element child DomNode object for this DomNode object.

## Syntax

```
getFirstChildElement()
  RETURNS xml.DomNode
```

## Usage

Returns the first XML Element child `DomNode` object for this DomNode object, or
NULL.

> **Important:**
>
> This method is not part of W3C standard
> API.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
