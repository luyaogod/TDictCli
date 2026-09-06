---
title: "xml.StaxReader.readFromPipe"
source: "fgl-topics/c_gws_XmlStaxReader_readFromPipe.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.readFromPipe"
type: "concept"
---

# xml.StaxReader.readFromPipe

> Sets the input stream of the StaxReader object to a PIPE and starts the streaming.

## Syntax

```
readFromPipe(
   name STRING )
```

1. name defines the command to start the
   PIPE and where the reader will get the XML from.

## Usage

Use this method to set the input stream of the StaxReader object to a PIPE, where
name is the command to read from the PIPE.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
