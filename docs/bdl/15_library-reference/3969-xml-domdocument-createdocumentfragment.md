---
title: "xml.DomDocument.createDocumentFragment"
source: "fgl-topics/c_gws_XmlDomDocument_createDocumentFragment.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createDocumentFragment"
type: "concept"
---

# xml.DomDocument.createDocumentFragment

> Creates a XML Document Fragment xml.DomNode object for this xml.DomDocument object.

## Syntax

```
createDocumentFragment()
  RETURNS xml.DomNode
```

Returns an `xml.DomNode` object.

## Usage

Creates a XML Document Fragment `xml.DomNode` object for this
`xml.DomDocument` object.

Returns the XML element `xml.DomNode` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")
