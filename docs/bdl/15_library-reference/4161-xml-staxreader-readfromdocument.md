---
title: "xml.StaxReader.readFromDocument"
source: "fgl-topics/c_gws_XmlStaxReader_readFromDocument.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.readFromDocument"
type: "concept"
---

# xml.StaxReader.readFromDocument

> Sets the input stream of the StaxReader object to a DomDocument object and starts the streaming.

## Syntax

```
readFromDocument(
   _doc xml.DomDocument )
```

1. \_doc defines an `XML.DomDocument` object containing an XML
   document.

## Usage

Use this method to set the input stream of the StaxReader object to a
`xml.DomDocument` object, where \_doc is a valid object containing
an XML document.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
