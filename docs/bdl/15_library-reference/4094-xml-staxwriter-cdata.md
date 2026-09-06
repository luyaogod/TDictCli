---
title: "xml.StaxWriter.cdata"
source: "fgl-topics/c_gws_XmlStaxWriter_cdata.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.cdata"
type: "concept"
---

# xml.StaxWriter.cdata

> Writes a XML CData to the StaxWriter stream.

## Syntax

```
cdata(
   cdata STRING )
```

1. cdata defines the data contained in the CData section, or NULL.

## Usage

This method writes XML character data passed as parameter as a CData.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
