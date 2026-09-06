---
title: "xml.StaxReader.readFromText"
source: "fgl-topics/c_gws_XmlStaxReader_readFromText.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.readFromText"
type: "concept"
---

# xml.StaxReader.readFromText

> Sets the input stream of the StaxReader object to a TEXT large object and starts the streaming.

## Syntax

```
readFromText(
   _txt TEXT )
```

1. \_txt defines a TEXT lob located in memory and containing the XML to read.

## Usage

Use this method to set the input stream of the StaxReader object to a TEXT large object, where
\_txt is the text to read.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
