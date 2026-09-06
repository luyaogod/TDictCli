---
title: "xml.DomDocument.getDocumentElement"
source: "fgl-topics/c_gws_XmlDomDocument_getDocumentElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getDocumentElement"
type: "concept"
---

# xml.DomDocument.getDocumentElement

> Returns the root XML Element xml.DomNode object for this xml.DomDocument object.

## Syntax

```
getDocumentElement()
  RETURNS xml.DomNode
```

## Usage

Returns the root XML Element `xml.DomNode` object for this
`xml.DomDocument` object.

Returns the XML element `xml.DomNode` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example 1 : Create a namespace qualified document with processing instructions](4018-example-1-create-a-namespace-qualified-document-with-process.md "Example 1 : Create a namespace qualified document with processing instructions")
