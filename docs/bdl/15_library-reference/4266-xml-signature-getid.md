---
title: "xml.Signature.getID"
source: "fgl-topics/c_gws_XmlSignature_getID.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.getID"
type: "concept"
---

# xml.Signature.getID

> Returns the ID value of the signature.

## Syntax

```
getID()
  RETURNS STRING
```

## Usage

This method returns the ID value of the signature.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
