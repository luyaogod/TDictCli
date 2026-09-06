---
title: "xml.StaxReader.getVersion"
source: "fgl-topics/c_gws_XmlStaxReader_getVersion.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.getVersion"
type: "concept"
---

# xml.StaxReader.getVersion

> Returns the document version defined in the XML Document declaration, or NULL.

## Syntax

```
getVersion()
  RETURNS STRING
```

## Usage

This method returns the document version as defined in the XML document declaration, or NULL if
not defined.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
