---
title: "xml.DomNode.appendChild"
source: "fgl-topics/c_gws_XmlDomNode_appendChild.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > xml.DomNode methods > xml.DomNode.appendChild"
type: "concept"
---

# xml.DomNode.appendChild

> Adds a child DomNode object to the end of the child list for a DomNode object

## Syntax

```
appendChild(
  newChild xml.DomNode )
```

1. newChild defines the node to add

## Usage

Adds a child `DomNode` object to the end of the child list for this DomNode
object.

The DomNode object node must be the child
of an element or document node; otherwise the operation fails.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example 1 : Create a namespace qualified document with processing instructions](4018-example-1-create-a-namespace-qualified-document-with-process.md "Example 1 : Create a namespace qualified document with processing instructions")
