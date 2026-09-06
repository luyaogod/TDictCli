---
title: "xml.Signature.setObjectID"
source: "fgl-topics/c_gws_XmlSignature_setObjectID.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.setObjectID"
type: "concept"
---

# xml.Signature.setObjectID

> Sets an ID for the signature object.

## Syntax

```
setObjectID(
   index INTEGER,
   id STRING )
```

1. index defines the index value.
2. id defines the value to be set.

## Usage

This method sets the ID defined by id value for the signature object specified
in the index (index).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
