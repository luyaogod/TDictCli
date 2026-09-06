---
title: "xml.DomNode.getPreviousSibling"
source: "fgl-topics/c_gws_XmlDomNode_getPreviousSibling.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getPreviousSibling"
type: "concept"
---

# xml.DomNode.getPreviousSibling

> Returns the DomNode object immediately preceding a DomNode object.

## Syntax

```
getPreviousSibling()
  RETURNS xml.DomNode
```

## Usage

Use this method to return the `xml.DomNode` object immediately preceding this
`DomNode` object, or `NULL`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
