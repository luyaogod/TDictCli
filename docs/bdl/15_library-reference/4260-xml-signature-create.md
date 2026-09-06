---
title: "xml.Signature.Create"
source: "fgl-topics/c_gws_XmlSignature_Create.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.Create"
type: "concept"
---

# xml.Signature.Create

> Constructor of a blank Signature object.

## Syntax

```
xml.Signature.Create()
  RETURNS xml.Signature
```

## Usage

Returns a [XML
Signature object](4255-the-signature-class.md "The xml.Signature class provides methods to create detached, enveloped or enveloping XML signatures of one or more references of XML documents or document fragments, and to determine whether a signed referenced document has been modified afterwards.") or
NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
