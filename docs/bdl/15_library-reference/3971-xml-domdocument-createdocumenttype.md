---
title: "xml.DomDocument.createDocumentType"
source: "fgl-topics/c_gws_XmlDomDocument_createDocumentType.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.createDocumentType"
type: "concept"
---

# xml.DomDocument.createDocumentType

> Creates a XML Document Type (DTD) xml.DomNode object for this xml.DomDocument object.

## Syntax

```
createDocumentType(
   name STRING,
   publicID STRING,
   systemID STRING,
   internalSubset STRING )  
  RETURNS xml.DomNode
```

1. name defines the name of the document type.
2. publicID defines the URI of the public identifier.
3. systemID defines the URL of the system identifier (Specifies the file
   location of the external DTD subset).
4. internalSubset defines the internal DTD subset.

Returns an `xml.DomNode` object.

## Usage

Creates a XML Document Type (DTD) `xml.DomNode` object for this
`xml.DomDocument` object; name is the name of the document type;
publicID is the URI of the public identifier or `NULL`;
systemID is the URL of the system identifier or `NULL` (Specifies
the file location of the external DTD subset); internalSubset is the internal DTD
subset or `NULL`.

Returns an `xml.DomNode` object, or an error if internalSubset
is malformed.

> **Important:**
>
> This method is not part of W3C standard
> API.

Only internal DTDs are supported.

The public identifier cannot be set without
the system identifier.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Node creation methods usage examples](4012-node-creation-methods-usage-examples.md "Node creation methods usage examples for the xml.DomDocument class.")
