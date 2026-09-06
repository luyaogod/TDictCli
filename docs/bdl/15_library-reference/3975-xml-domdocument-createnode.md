---
title: "xml.DomDocument.createNode"
source: "fgl-topics/c_gws_XmlDomDocument_createNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createNode"
type: "concept"
---

# xml.DomDocument.createNode

> Creates an xml.DomNode object from a string for this xml.DomDocument object.

## Syntax

```
createNode(
   str STRING ) 
  RETURNS xml.DomNode
```

1. str defines the string representation of the `xml.DomNode` to
   be created.

Returns an `xml.DomNode` object.

## Usage

Creates an `xml.DomNode` object from a string for this
`xml.DomDocument` object; str is the string representation of the
`xml.DomNode` to be created.

Returns the XML element `xml.DomNode` object.

> **Important:**
>
> This method is not part of W3C standard
> API.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")
