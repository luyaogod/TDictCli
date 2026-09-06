---
title: "xml.XsltTransformer.CreateFromDocument"
source: "fgl-topics/c_gws_XmlXsltTransformer_CreateFromDocument.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML transformation classes > The XsltTransformer class > XsltTransformer methods > xml.XsltTransformer.CreateFromDocument"
type: "concept"
---

# xml.XsltTransformer.CreateFromDocument

> Returns a new instance of a XsltTransformer object to be used to transform a XML document based on a given style sheet.

## Syntax

```
xml.XsltTransformer.CreateFromDocument(
   doc xml.DomDocument )
  RETURNS xml.XsltTransformer
```

1. doc defines the given style sheet.

## Usage

Returns a new instance of an `XsltTransformer` object to be used to transform a
XML document based on the given style sheet.

No exception is raised in case of error. It is recommended to check if errors have been
encountered loading the style sheet by using the [xml.XsltTransformer.getErrorsCount](4347-xml-xslttransformer-geterrorscount.md "Returns the number of errors.") and [xml.XsltTransformer.getErrorDescription](4346-xml-xslttransformer-geterrordescription.md "Returns the exact description of the error referenced by the index.") methods.

To see these methods in a working example, see [Example: Using xml.XsltTransformer methods](4350-example.md "This Genero application provides a working example using methods from the XsltTransformer class.").
