---
title: "xml.StaxWriter.startElement"
source: "fgl-topics/c_gws_XmlStaxWriter_startElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.startElement"
type: "concept"
---

# xml.StaxWriter.startElement

> Writes a XML start element to the StaxWriter stream.

## Syntax

```
startElement(
   name STRING )
```

1. name defines the local name of the XML start element, cannot be NULL.

## Usage

This method writes a XML start element to the `StaxWriter` stream. All
`startElement` methods open a new scope and set the stream to a
`START_ELEMENT`. Writing the corresponding `endElement()` causes the scope to be
closed.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
