---
title: "xml.DomDocument.getXmlEncoding"
source: "fgl-topics/c_gws_XmlDomDocument_getXmlEncoding.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getXmlEncoding"
type: "concept"
---

# xml.DomDocument.getXmlEncoding

> Returns the document encoding as defined in the XML document declaration.

## Syntax

```
getXmlEncoding()
  RETURNS STRING
```

## Usage

This method returns the document encoding as defined in the XML document declaration, or
`NULL` if there is none.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
