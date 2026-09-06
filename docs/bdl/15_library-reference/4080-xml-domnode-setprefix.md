---
title: "xml.DomNode.setPrefix"
source: "fgl-topics/c_gws_XmlDomNode_setPrefix.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.setPrefix"
type: "concept"
---

# xml.DomNode.setPrefix

> Sets the prefix for a DomNode object.

## Syntax

```
setPrefix(
  prefix STRING )
```

1. prefix defines the prefix for this DomNode object.

## Usage

Use this method to set the prefix for this [xml.DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object.

This method is only valid on namespace qualified Element or Attribute nodes.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
