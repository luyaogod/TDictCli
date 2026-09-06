---
title: "xml.DomDocument.removeDocumentNode"
source: "fgl-topics/c_gws_XmlDomDocument_removeDocumentNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.removeDocumentNode"
type: "concept"
---

# xml.DomDocument.removeDocumentNode

> Removes a child xml.DomNode object from the xml.DomNode children in this xml.DomDocument object.

## Syntax

```
removeDocumentNode(
   n xml.DomNode )
```

1. n defines the node to remove.

## Usage

Removes a child `xml.DomNode` object from the `xml.DomNode`
children in this `xml.DomDocument` object, where n is the node to
remove.

Only Text nodes, Processing Instruction nodes, Document
Fragment nodes, one Element node, and one Document Type node allowed.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
