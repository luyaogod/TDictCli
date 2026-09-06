---
title: "xml.Signature.setID"
source: "fgl-topics/c_gws_XmlSignature_setID.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.setID"
type: "concept"
---

# xml.Signature.setID

> Sets an ID value for the signature.

## Syntax

```
setID(
   id STRING )
```

1. id defines the ID value to be
   set.

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
