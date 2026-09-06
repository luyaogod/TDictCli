---
title: "xml.DomDocument.createElement"
source: "fgl-topics/c_gws_XmlDomDocument_createElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createElement"
type: "concept"
---

# xml.DomDocument.createElement

> Creates a XML Element xml.DomNode object for an xml.DomDocument object

## Syntax

```
createElement(
   elt STRING ) 
  RETURNS xml.DomNode
```

1. elt defines the name of the XML element.

Returns an `xml.DomNode` object.

## Usage

Creates a XML Element `xml.DomNode` object for this
`xml.DomDocument` object, where elt is the name of the XML
element, which cannot be `NULL`.

Returns the XML element `xml.DomNode` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")
