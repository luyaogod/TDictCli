---
title: "xml.DomDocument.loadFromPipe"
source: "fgl-topics/c_gws_XmlDomDocument_loadFromPipe.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.loadFromPipe"
type: "concept"
---

# xml.DomDocument.loadFromPipe

> Loads a XML Document into this xml.DomDocument object from a PIPE.

## Syntax

```
loadFromPipe(
   name STRING )
```

1. name defines the command to read from the PIPE.

## Usage

Use this method to load a XML Document into a `xml.DomDocument` object from a
PIPE, where name is the command to read from the PIPE.

> **Important:**
>
> This method is not part of W3C standard
> API.

Use [`setFeature()`](4006-xml-domdocument-setfeature.md "Sets a feature for this xml.DomDocument object.") to
specify how the document can be loaded. HTML parsing is possible when [enable-html-compliancy](4016-domdocument-features.md "A list of features for the xml.DomDocument class.") is
enabled.

See [`getErrorsCount()`](3986-xml-domdocument-geterrorscount.md "Returns the number of errors encountered during the loading, saving or validation of a XML document.") and [`getErrorDescription()`](3985-xml-domdocument-geterrordescription.md "Returns the error description at the given position.")
to retrieve error messages related to XML document.

When loading a document, if [xml.DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") objects are still referenced in other variables of the program, the
entire document is kept in memory. Otherwise, the DOM nodes of the document are deleted
before loading the new document. For more details about object references and garbage
collection in BDL, see [Working with objects](../09_advanced-features/0946-working-with-objects.md "This topic describes basic object usage in Genero BDL.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
