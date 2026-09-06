---
title: "xml.Signature.getDocument"
source: "fgl-topics/c_gws_XmlSignature_getDocument.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.getDocument"
type: "concept"
---

# xml.Signature.getDocument

> Returns a new DomDocument object representing the signature in XML.

## Syntax

```
getDocument()
  RETURNS xml.DomDocument
```

## Usage

Returns a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.")
object.

If the signature type is `enveloped`, it's up to the user to add it at the right
place in the XML document it is intended to sign.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
