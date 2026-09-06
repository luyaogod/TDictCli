---
title: "xml.DomDocument.createAttribute"
source: "fgl-topics/c_gws_XmlDomDocument_createAttribute.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createAttribute"
type: "concept"
---

# xml.DomDocument.createAttribute

> Creates a XML Attribute xml.DomNode object for an xml.DomDocument object.

## Syntax

```
createAttribute(
   elt STRING )  
  RETURNS xml.DomNode
```

1. elt defines the name of the XML attribute.

Returns an `xml.DomNode` object.

## Usage

Creates an XML Attribute `xml.DomNode` object for an
`xml.DomDocument` object, where elt is the name of the XML
attribute. It cannot be `NULL`.

Returns the XML element `xml.DomNode` object.

To create a default namespace declaration attribute use `xmlns` as the name.
(Using [`declareNamespace`](3978-xml-domdocument-declarenamespace.md "Forces namespace declaration to a XML Element xml.DomNode for this xml.DomDocument object.") instead is recommended)

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")
