---
title: "xml.DomDocument.createEntityReference"
source: "fgl-topics/c_gws_XmlDomDocument_createEntityReference.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createEntityReference"
type: "concept"
---

# xml.DomDocument.createEntityReference

> Creates a XML EntityReference xml.DomNode object for this xml.DomDocument object

## Syntax

```
createEntityReference(
   entity STRING )  
  RETURNS xml.DomNode
```

1. entity defines the name of the entity reference.

Returns an `xml.DomNode` object.

## Usage

Creates a XML EntityReference `xml.DomNode` object for this
`xml.DomDocument` object, where entity is the name of the entity
reference.

Returns the XML element `xml.DomNode` object.

An Entity Reference node is read-only and cannot
be modified.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")
