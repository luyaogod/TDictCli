---
title: "xml.XsltTransformer.doTransform"
source: "fgl-topics/c_gws_XmlXsltTransformer_doTransform.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML transformation classes > The XsltTransformer class > XsltTransformer methods > xml.XsltTransformer.doTransform"
type: "concept"
---

# xml.XsltTransformer.doTransform

> Transforms the input source based on the XSLT style sheet used to create the XsltTransformer instance and returns the transformed document.

## Syntax

```
doTransform(
   doc xml.DomDocument )
  RETURNS xml.DomDocument
```

1. doc defines the input source.

## Usage

This method transforms the input source based on the XSLT style sheet used to create the
`XSLTTtransformer` instance and returns the transformed document. It raises a
`XML_ERROR_FAILED` exception in case of critical error.

The operation may return a new document that may be incomplete. You must check for errors that
may occur during the transformation with the [xml.XSLTtransfomer.getErrorsCount()](4347-xml-xslttransformer-geterrorscount.md "Returns the number of errors.") and [xml.XsltTransformer.getErrorDescription](4346-xml-xslttransformer-geterrordescription.md "Returns the exact description of the error referenced by the index.") methods.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

To see these methods in a working example, see [Example: Using xml.XsltTransformer methods](4350-example.md "This Genero application provides a working example using methods from the XsltTransformer class.").
