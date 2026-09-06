---
title: "xml.DomDocument.insertAfterDocumentNode"
source: "fgl-topics/c_gws_XmlDomDocument_insertAfterDocumentNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.insertAfterDocumentNode"
type: "concept"
---

# xml.DomDocument.insertAfterDocumentNode

> Inserts a child xml.DomNode object after another child DomNode in this xml.DomDocument object.

## Syntax

```
insertAfterDocumentNode(
   newNode xml.DomNode,
   ref xml.DomNode )
```

1. newNode defines the node to insert.
2. ref defines the reference node (the node after which the new node must be inserted).

## Usage

Inserts a child `xml.DomNode` object after another child
`xml.DomNode` in this `xml.DomDocument` object ;
newNode is the node to insert; ref is the reference node (the
node after which the new node must be inserted).

> **Important:**
>
> This method is not part of W3C standard
> API.

Only Text nodes, Processing Instruction nodes, Document
Fragment nodes, one Element node, and one Document Type node allowed.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
