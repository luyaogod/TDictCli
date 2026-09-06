---
title: "xml.Signature.getCanonicalization"
source: "fgl-topics/c_gws_XmlSignature_getCanonicalization.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.getCanonicalization"
type: "concept"
---

# xml.Signature.getCanonicalization

> Returns the canonicalization identifier of the signature.

## Syntax

```
getCanonicalization()
  RETURNS STRING
```

## Usage

Returns one of the four canonicalization [identifiers](4296-transformation-identifier.md) of the signature.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
