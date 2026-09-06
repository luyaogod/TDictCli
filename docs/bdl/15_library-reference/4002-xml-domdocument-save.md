---
title: "xml.DomDocument.save"
source: "fgl-topics/c_gws_XmlDomDocument_save.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.save"
type: "concept"
---

# xml.DomDocument.save

> Saves this xml.DomDocument object as a XML Document to a file or URL.

## Syntax

```
save(
   filename STRING )
```

1. filename defines a valid URL or the name of a file.

## Usage

Saves a `xml.DomDocument` object as a XML Document to a file or URL, where
filename is a valid URL or the name of the file.

Only the following kinds of URLs are supported:

- http://
- https://
- tcp://
- tcps://
- file:///
- alias://

See [fglprofile
Configuration](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") for more details about URL mapping with aliases, and for proxy and security
configuration.

See [`setFeature()`](4006-xml-domdocument-setfeature.md "Sets a feature for this xml.DomDocument object.") to
specify how the document can be saved.

See [`getErrorsCount()`](3986-xml-domdocument-geterrorscount.md "Returns the number of errors encountered during the loading, saving or validation of a XML document.") and [`getErrorDescription()`](3985-xml-domdocument-geterrordescription.md "Returns the error description at the given position.")
to retrieve error messages related to XML document.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
