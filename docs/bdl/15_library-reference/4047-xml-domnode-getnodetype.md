---
title: "xml.DomNode.getNodeType"
source: "fgl-topics/c_gws_XmlDomNode_getNodeType.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.getNodeType"
type: "concept"
---

# xml.DomNode.getNodeType

> Gets the XML type for this DomNode object.

## Syntax

```
getNodeType()
  RETURNS STRING
```

## Usage

This method returns the XML type for this DomNode object; it returns one of the XML [DomNode types](4082-domnode-types.md "List of types for the xml.DomNode class."), or NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
