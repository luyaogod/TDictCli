---
title: "xml.Signature.setReferenceID"
source: "fgl-topics/c_gws_XmlSignature_setReferenceID.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.setReferenceID"
type: "concept"
---

# xml.Signature.setReferenceID

> Sets an ID for the signature reference in the specified signature object.

## Syntax

```
setReferenceID(
   index INTEGER,
   id STRING )
```

1. index defines the index value.
2. id defines the value to be set.

## Usage

This method sets the ID for the signature reference defined by the ID value
for the signature object specified in the index (index).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
