---
title: "xml.DomNode.setNodeValue"
source: "fgl-topics/c_gws_XmlDomNode_setNodeValue.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.setNodeValue"
type: "concept"
---

# xml.DomNode.setNodeValue

> Sets the node value for a DomNode object.

## Syntax

```
setNodeValue(
   value STRING )
```

1. value defines the node value.

## Usage

This method sets the node value for this DomNode object, where value is the
node value.

Use of this method is only recommended for nodes that are not parents of other nodes, which means
it can be used for a node of type:

- ATTRIBUTE\_NODE
- TEXT\_NODE
- CDATA\_SECTION\_NODE
- PROCESSING\_INSTRUCTION\_NODE
- COMMENT\_NODE

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
