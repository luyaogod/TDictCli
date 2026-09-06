---
title: "xml.DomDocument.normalize"
source: "fgl-topics/c_gws_XmlDomDocument_normalize.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.normalize"
type: "concept"
---

# xml.DomDocument.normalize

> Normalizes the entire Document.

## Syntax

```
normalize()
```

## Usage

Normalizes the entire Document. This method merges adjacent text nodes, removes empty text nodes,
and sets namespace declarations as if the document had been saved.

See [`getErrorsCount()`](3986-xml-domdocument-geterrorscount.md "Returns the number of errors encountered during the loading, saving or validation of a XML document.") and [`getErrorDescription()`](3985-xml-domdocument-geterrordescription.md "Returns the error description at the given position.")
to retrieve error messages related to XML document.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
