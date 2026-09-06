---
title: "xml.Signature.getReferenceURI"
source: "fgl-topics/c_gws_XmlSignature_getReferenceURI.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.getReferenceURI"
type: "concept"
---

# xml.Signature.getReferenceURI

> Returns the URI of the reference in this signature object.

## Syntax

```
getReferenceURI(
   index INTEGER )
  RETURNS STRING
```

1. index defines the index in this signature object.

## Usage

This method returns the URI of the reference specified in the index index in
this signature object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
