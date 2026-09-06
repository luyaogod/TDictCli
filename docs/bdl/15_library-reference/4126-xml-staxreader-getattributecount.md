---
title: "xml.StaxReader.getAttributeCount"
source: "fgl-topics/c_gws_XmlStaxReader_getAttributeCount.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.getAttributeCount"
type: "concept"
---

# xml.StaxReader.getAttributeCount

> Returns the number of XML attributes defined on the current XML node, or zero.

## Syntax

```
getAttributeCount()
  RETURNS INTEGER
```

## Usage

Use this method to return the number of XML attributes defined on the current XML node.

This method is only valid on a
`START_ELEMENT` node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
