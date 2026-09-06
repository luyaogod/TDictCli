---
title: "xml.DomDocument.createAttributeNS"
source: "fgl-topics/c_gws_XmlDomDocument_createAttributeNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createAttributeNS"
type: "concept"
---

# xml.DomDocument.createAttributeNS

> Creates a XML namespace-qualified Attribute xml.DomNode object for this xml.DomDocument object.

## Syntax

```
createAttributeNS(
   prefix STRING,
   elt STRING,
   ns STRING )
  RETURNS xml.DomNode
```

1. prefix defines the prefix of the XML attribute.
2. elt defines the name of the XML attribute.
3. ns defines the namespace URI of the XML attribute.

Returns an `xml.DomNode` object.

## Usage

Creates a XML namespace-qualified Attribute `xml.DomNode` object for this
`xml.DomDocument` object where prefix is the prefix of the XML
attribute, it cannot be `NULL`; elt is the name of the XML
attribute; it cannot be `NULL`; ns is the namespace URI of the XML
attribute, it cannot be `NULL`.

Returns the XML element `xml.DomNode` object.

To create a namespace declaration attribute use `xmlns` as the prefix and
`http://www.w3.org/XML/1998/namespace` as the namespace. Using [`declareNamespace`](3978-xml-domdocument-declarenamespace.md "Forces namespace declaration to a XML Element xml.DomNode for this xml.DomDocument object.") instead
is recommended.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")
