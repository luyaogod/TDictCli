---
title: "xml.DomDocument.CreateDocument"
source: "fgl-topics/c_gws_XmlDomDocument_createDocument.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.CreateDocument"
type: "concept"
---

# xml.DomDocument.CreateDocument

> Constructor of an xml.DomDocument with an XML root element.

## Syntax

```
xml.DomDocument.CreateDocument(
   name STRING )
  RETURNS xml.DomDocument
```

1. name defines the XML element.

Returns an `xml.DomDocument` object.

## Usage

Constructor of an `xml.DomDocument` with an XML root element; where
name is the name of the XML Element.

Returns an `xml.DomDocument` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

Create an `xml.DomDocument` with an initial root node named ARoot:

```
DEFINE doc xml.DomDocument
LET doc = xml.DomDocument.CreateDocument("ARoot")
```

## Related links

**Related concepts**  

[xml.DomDocument.Create](3963-xml-domdocument-create.md "Constructor of an empty xml.DomDocument object.")
