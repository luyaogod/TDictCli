---
title: "xml.DomDocument.prependDocumentNode"
source: "fgl-topics/c_gws_XmlDomDocument_prependDocumentNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.prependDocumentNode"
type: "concept"
---

# xml.DomDocument.prependDocumentNode

> Adds a child xml.DomNode object to the beginning of the xml.DomNode children of this xml.DomDocument object.

## Syntax

```
prependDocumentNode(
   n xml.DomNode )
```

1. n defines the node to add.

## Usage

Adds a child `xml.DomNode` object to the beginning of the
`xml.DomNode` children in this `xml.DomDocument`object;
n is the node to add.

> **Important:**
>
> This method is not part of W3C standard
> API.

See [`getErrorsCount()`](3986-xml-domdocument-geterrorscount.md "Returns the number of errors encountered during the loading, saving or validation of a XML document.") and [`getErrorDescription()`](3985-xml-domdocument-geterrordescription.md "Returns the error description at the given position.")
to retrieve error messages related to XML document.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example 1 : Create a namespace qualified document with processing instructions](4018-example-1-create-a-namespace-qualified-document-with-process.md "Example 1 : Create a namespace qualified document with processing instructions")
