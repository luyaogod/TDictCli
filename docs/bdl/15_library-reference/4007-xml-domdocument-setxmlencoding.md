---
title: "xml.DomDocument.setXmlEncoding"
source: "fgl-topics/c_gws_XmlDomDocument_setXmlEncoding.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.setXmlEncoding"
type: "concept"
---

# xml.DomDocument.setXmlEncoding

> Sets the XML document encoding in the XML declaration of this xml.DomDocument.

## Syntax

```
setXmlEncoding(
   encoding STRING )
```

1. encoding defines the XML document encoding.

## Usage

Sets the XML document encoding in the XML declaration, or `NULL`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
