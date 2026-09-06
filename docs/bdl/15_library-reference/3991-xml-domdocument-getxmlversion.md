---
title: "xml.DomDocument.getXmlVersion"
source: "fgl-topics/c_gws_XmlDomDocument_getXmlVersion.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.getXmlVersion"
type: "concept"
---

# xml.DomDocument.getXmlVersion

> Returns the document version as defined in the XML document declaration.

## Syntax

```
getXmlVersion()
  RETURNS STRING
```

## Usage

This method returns the document version as defined in the XML document declaration, which is
1.0. No other versions are supported.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
