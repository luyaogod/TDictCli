---
title: "xml.DomDocument.validateOneElement"
source: "fgl-topics/c_gws_XmlDomDocument_validateOneElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.validateOneElement"
type: "concept"
---

# xml.DomDocument.validateOneElement

> Performs a DTD or XML Schema validation of a XML Element xml.DomNode object.

## Syntax

```
validateOneElement(
   elt xml.DomNode )
  RETURNS INTEGER
```

1. elt defines the XML element `xml.DomNode` to validate.

## Usage

Performs a DTD or XML Schema validation of a XML Element `xml.DomNode` object;
elt is the XML Element `xml.DomNode` to validate.

Returns the number of validation errors, or zero if there are none.

> **Important:**
>
> This method is not part of W3C standard
> API.

See [`setFeature()`](4006-xml-domdocument-setfeature.md "Sets a feature for this xml.DomDocument object.")
to specify what kind of validation to do.

See [`getErrorsCount()`](3986-xml-domdocument-geterrorscount.md "Returns the number of errors encountered during the loading, saving or validation of a XML document.") and [`getErrorDescription()`](3985-xml-domdocument-geterrordescription.md "Returns the error description at the given position.")
to retrieve error messages related to XML document.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
