---
title: "xml.DomDocument.createTextNode"
source: "fgl-topics/c_gws_XmlDomDocument_createTextNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createTextNode"
type: "concept"
---

# xml.DomDocument.createTextNode

> Creates a XML Text xml.DomNode object for this xml.DomDocument object.

## Syntax

```
createTextNode(
   txt STRING )  
  RETURNS xml.DomNode
```

1. txt defines the data of the XML Text
   node.

Returns an `xml.DomNode` object.

## Usage

Creates a XML Text `xml.DomNode` object for this `xml.DomDocument`
object, where txt is the data of the XML Text node, or `NULL`.

Returns the XML element `xml.DomNode` object.

Only the characters `#x9`, `#xA`, `#xD`,
[`#x20`-`#xD7FF`], [`#xE000`-`#xFFFD`]
and [`#x10000`-`#x10FFFF`] are allowed in the content of an XML Text
node. The [`save()`](4002-xml-domdocument-save.md "Saves this xml.DomDocument object as a XML Document to a file or URL.") and [`normalize()`](3999-xml-domdocument-normalize.md "Normalizes the entire Document.") methods will fail if characters other than those allowed exist
in a Text node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")
