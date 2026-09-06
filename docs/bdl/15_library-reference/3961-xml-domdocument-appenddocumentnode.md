---
title: "xml.DomDocument.appendDocumentNode"
source: "fgl-topics/c_gws_XmlDomDocument_appendDocumentNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.appendDocumentNode"
type: "concept"
---

# xml.DomDocument.appendDocumentNode

> Adds a child DomNode object to the end of the DomNode children for this DomDocument object.

## Syntax

```
appendDocumentNode(
   n xml.DomNode )
```

1. n defines the node to add.

## Usage

Adds a child `xml.DomNode` object to the end of the `xml.DomNode`
children for this `xml.DomDocument` object, where n is the node to
add.

Only Text nodes, Processing Instruction nodes, Document
Fragment nodes, one Element node, and one Document Type node allowed.

> **Note:**
>
> A fragment is a structure created to receive XML nodes that are not always valid.
> Once a fragment is added to a valid node, the fragment becomes empty as all nodes
> are moved from the fragment as a child to a valid node. So developers can work on
> the fragment until it is added to another node. At that time developers should no
> longer work on the fragment but rather on the valid node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[HTML document usage example](4013-html-document-usage-example.md "The HTML language provides tags that allow the user to provide an embedded style sheet (the \"style\" tag) and to write embedded client side script (the \"script\" tag). According to the HTML 4.0 specification, the content of these tags must be managed as CDATA section.")

[Example 1 : Create a namespace qualified document with processing instructions](4018-example-1-create-a-namespace-qualified-document-with-process.md "Example 1 : Create a namespace qualified document with processing instructions")
