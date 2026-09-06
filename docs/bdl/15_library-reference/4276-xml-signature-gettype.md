---
title: "xml.Signature.getType"
source: "fgl-topics/c_gws_XmlSignature_getType.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.getType"
type: "concept"
---

# xml.Signature.getType

> Returns a string with the type of the signature object.

## Syntax

```
getType()
  RETURNS STRING
```

## Usage

The string can be [`Detached`](http://www.w3.org/TR/xmldsig-core/#def-SignatureDetached), [`Enveloped`](http://www.w3.org/TR/xmldsig-core/#def-SignatureEnveloped), [`Enveloping`](http://www.w3.org/TR/xmldsig-core/#def-SignatureEnveloping) or `Invalid` based on the XML-Signature
specification.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
