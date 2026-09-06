---
title: "xml.Signature.getReferenceDigest"
source: "fgl-topics/c_gws_XmlSignature_getReferenceDigest.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.getReferenceDigest"
type: "concept"
---

# xml.Signature.getReferenceDigest

> Returns the digest algorithm identifier of the reference.

## Syntax

```
getReferenceDigest(
   index INTEGER )
  RETURNS STRING
```

1. index defines the index in this signature object.

## Usage

Returns the digest algorithm [identifier](4295-digest-identifier.md) of the reference of index ( *index*) in this signature object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
