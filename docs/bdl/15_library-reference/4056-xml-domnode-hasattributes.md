---
title: "xml.DomNode.hasAttributes"
source: "fgl-topics/c_gws_XmlDomNode_hasAttributes.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.hasAttributes"
type: "concept"
---

# xml.DomNode.hasAttributes

> Identifies whether a node has XML Attribute nodes.

## Syntax

```
hasAttributes()
  RETURNS INTEGER
```

## Usage

This method returns `TRUE` if this node has XML Attribute nodes; otherwise returns
`FALSE`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
