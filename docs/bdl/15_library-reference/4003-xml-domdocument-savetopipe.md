---
title: "xml.DomDocument.saveToPipe"
source: "fgl-topics/c_gws_XmlDomDocument_saveToPipe.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.saveToPipe"
type: "concept"
---

# xml.DomDocument.saveToPipe

> Saves this xml.DomDocument object as an XML Document via a PIPE.

## Syntax

```
saveToPipe(
   name STRING )
```

1. name defines the command to start the pipe.

## Usage

Saves a `xml.DomDocument` object as a XML Document to a PIPE, where
name is the command to start the pipe.

See [`setFeature()`](4006-xml-domdocument-setfeature.md "Sets a feature for this xml.DomDocument object.") to specify how the document can be saved.

See [`getErrorsCount()`](3986-xml-domdocument-geterrorscount.md "Returns the number of errors encountered during the loading, saving or validation of a XML document.") and [`getErrorDescription()`](3985-xml-domdocument-geterrordescription.md "Returns the error description at the given position.")
to retrieve error messages related to XML document.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
