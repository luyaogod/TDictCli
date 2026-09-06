---
title: "xml.XsltTransformer.getErrorDescription"
source: "fgl-topics/c_gws_XmlXsltTransformer_getErrorDescription.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML transformation classes > The XsltTransformer class > XsltTransformer methods > xml.XsltTransformer.getErrorDescription"
type: "concept"
---

# xml.XsltTransformer.getErrorDescription

> Returns the exact description of the error referenced by the index.

## Syntax

```
getErrorDescription(
   index INTEGER )
  RETURNS STRING
```

1. index defines the index of the error returned by `xml.XSLTtransfomer.getErrorsCount()`.

## Usage

This method returns the exact description of the error referenced by the index during the
`doTransform()` call or during the `CreateFromDocument()` when the
style sheet contains some errors.

This method is typically used in conjunction with `getErrorsCount()`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

To see these methods in a working example, see [Example: Using xml.XsltTransformer methods](4350-example.md "This Genero application provides a working example using methods from the XsltTransformer class.").

## Related links

**Related concepts**  

[xml.XsltTransformer.CreateFromDocument](4344-xml-xslttransformer-createfromdocument.md "Returns a new instance of a XsltTransformer object to be used to transform a XML document based on a given style sheet.")

[xml.XsltTransformer.doTransform](4345-xml-xslttransformer-dotransform.md "Transforms the input source based on the XSLT style sheet used to create the XsltTransformer instance and returns the transformed document.")
