---
title: "xml.DomDocument.isXmlStandalone"
source: "fgl-topics/c_gws_XmlDomDocument_isXmlStandalone.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.isXmlStandalone"
type: "concept"
---

# xml.DomDocument.isXmlStandalone

> Checks whether the XML standalone attribute is set in the XML declaration.

## Syntax

```
isXmlStandalone()
  RETURNS INTEGER
```

## Usage

Use this method to check if the XML standalone attribute is set in the XML declaration or not. It
returns `TRUE` if set to yes, `FALSE` if not.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[xml.DomDocument.setXmlEncoding](4007-xml-domdocument-setxmlencoding.md "Sets the XML document encoding in the XML declaration of this xml.DomDocument.")
