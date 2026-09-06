---
title: "xml.XsltTransformer methods"
source: "fgl-topics/c_gws_XmlXsltTransformer_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML transformation classes > The XsltTransformer class > XsltTransformer methods"
type: "concept"
---

# xml.XsltTransformer methods

> Methods for the xml.XSLTtransfomer class.

| Name | Description |
| --- | --- |
| xml.XsltTransformer.CreateFromDocument( doc xml.DomDocument ) RETURNS xml.XsltTransformer | Returns a new instance of a XsltTransformer object to be used to transform a XML document based on a given style sheet. |

| Name | Description |
| --- | --- |
| setParameter( param STRING, value STRING ) | Set parameters in an instance of an `XSLTTransformer` corresponding to the XSLT entries named by `param`. |
| getParameter( param STRING) RETURNS STRING | Get the value of a given parameter in an instance of an `XSLTTransformer` corresponding to the XSLT entry named by `param`. |
| clearParameters() | Clear all parameters in an instance of an `XSLTTransformer` object. |
| doTransform( doc xml.DomDocument ) RETURNS xml.DomDocument | Transforms the input source based on the XSLT style sheet used to create the `XsltTransformer` instance and returns the transformed document. |
| getErrorDescription( index INTEGER ) RETURNS STRING | Returns the exact description of the error referenced by the index. |
| getErrorsCount() RETURNS INTEGER | Returns the number of errors. |
