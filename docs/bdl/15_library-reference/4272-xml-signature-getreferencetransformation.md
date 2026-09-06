---
title: "xml.Signature.getReferenceTransformation"
source: "fgl-topics/c_gws_XmlSignature_getReferenceTransformation.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.getReferenceTransformation"
type: "concept"
---

# xml.Signature.getReferenceTransformation

> Gets the transformation identifier related to the reference of index referenceIndex.

## Syntax

```
getReferenceTransformation(
   referenceIndex INTEGER,
   index INTEGER )
  RETURNS STRING
```

1. referenceIndex is the index in this Signature object.
2. index defines the position in the
   list of transformation.

## Usage

Returns the transformation [identifier](4296-transformation-identifier.md) related to the reference of index *referenceIndex*, at the position
specified by *index* in the list of transformations.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
