---
title: "xml.DomDocument.createElementNS"
source: "fgl-topics/c_gws_XmlDomDocument_createElementNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createElementNS"
type: "concept"
---

# xml.DomDocument.createElementNS

> Creates a XML namespace-qualified Element xml.DomNode object for an xml.DomDocument object.

## Syntax

```
createElementNS(
   prefix STRING,
   elt STRING,
   ns STRING )
  RETURNS xml.DomNode
```

1. prefix defines the prefix of the XML element, or `NULL` to use
   the default namespace.
2. elt defines the name of the XML element.
3. ns defines the namespace URI of the XML element.

Returns an `xml.DomNode` object.

## Usage

Creates a XML namespace-qualified Element `xml.DomNode` object for this
`xml.DomDocument` object, where prefix is the prefix of the XML
element, or `NULL` to use the default namespace. elt is the name
of the XML element, this cannot be `NULL`. ns is the namespace URI
of the XML element, this cannot be `NULL`.

Returns the XML element `xml.DomNode` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")

[Example 1 : Create a namespace qualified document with processing instructions](4018-example-1-create-a-namespace-qualified-document-with-process.md "Example 1 : Create a namespace qualified document with processing instructions")
