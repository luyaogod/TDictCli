---
title: "xml.Signature.appendObjectData"
source: "fgl-topics/c_gws_XmlSignature_appendObjectData.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.appendObjectData"
type: "concept"
---

# xml.Signature.appendObjectData

> Appends a copy of a XML DomNode to the signature object index.

## Syntax

```
appendObjectData(
   index INTEGER,
   node xml.DomNode )
```

1. index defines the index in this
   Signature object.
2. `node` defines the [xml.DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
