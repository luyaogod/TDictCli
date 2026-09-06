---
title: "xml.DomDocument.Create"
source: "fgl-topics/c_gws_XmlDomDocument_create.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.Create"
type: "concept"
---

# xml.DomDocument.Create

> Constructor of an empty xml.DomDocument object.

## Syntax

```
xml.DomDocument.Create()
  RETURNS xml.DomDocument
```

## Usage

Constructor of an empty `xml.DomDocument` object.

Returns an `xml.DomDocument` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

Create an `xml.DomDocument`:

```
DEFINE doc xml.DomDocument
LET doc = xml.DomDocument.Create()
```

## Related links

**Related concepts**  

[xml.DomDocument.CreateDocument](3968-xml-domdocument-createdocument.md "Constructor of an xml.DomDocument with an XML root element.")
