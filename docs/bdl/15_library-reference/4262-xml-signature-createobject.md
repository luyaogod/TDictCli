---
title: "xml.Signature.createObject"
source: "fgl-topics/c_gws_XmlSignature_createObject.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.createObject"
type: "concept"
---

# xml.Signature.createObject

> Creates a new object that will embed additional XML nodes.

## Syntax

```
createObject()
  RETURNS INTEGER
```

## Usage

The returned value represents the index for any further manipulation of this signature
object.

> **Note:**
>
> An object is enveloping additional XML nodes, but is not necessarily signed unless
> there is a reference on it.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
