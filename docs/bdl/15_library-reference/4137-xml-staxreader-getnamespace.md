---
title: "xml.StaxReader.getNamespace"
source: "fgl-topics/c_gws_XmlStaxReader_getNamespace.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.getNamespace"
type: "concept"
---

# xml.StaxReader.getNamespace

> Returns the namespace URI of the current XML node, or NULL.

## Syntax

```
getNamespace()
  RETURNS STRING
```

## Usage

This method returns the namespace URI of the current XML node, or NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
