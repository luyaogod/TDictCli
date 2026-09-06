---
title: "xml.Signature.getReferenceCount"
source: "fgl-topics/c_gws_XmlSignature_getReferenceCount.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.getReferenceCount"
type: "concept"
---

# xml.Signature.getReferenceCount

> Returns the number of references in this Signature object.

## Syntax

```
getReferenceCount()
  RETURNS INTEGER
```

## Usage

This method returns the number of references in the specified signature object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
