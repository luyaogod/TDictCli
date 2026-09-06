---
title: "xml.StaxReader.Create"
source: "fgl-topics/c_gws_XmlStaxReader_Create.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.Create"
type: "concept"
---

# xml.StaxReader.Create

> Constructor of a StaxReader object.

## Syntax

```
xml.StaxReader.Create()
  RETURNS xml.StaxReader
```

## Usage

Use this method to create and return a [StaxReader](4121-the-staxreader-class.md "The StaxReader class provides methods compatible with Streaming API for XML(StAX) for reading XML documents.") object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
