---
title: "xml.StaxWriter.writeToDocument"
source: "fgl-topics/c_gws_XmlStaxWriter_writeToDocument.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.writeToDocument"
type: "concept"
---

# xml.StaxWriter.writeToDocument

> Sets the output stream of the StaxWriter object to an xml.DomDocument object, and starts the streaming.

## Syntax

```
writeToDocument(
   doc xml.DomDocument )
```

1. doc defines the empty `DomDocument` object that will contain
   the resulting XML document.

## Usage

This method sets the output stream of the `StaxWriter` object to a
`DomDocument` specified in doc, and starts the streaming.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
