---
title: "xml.StaxWriter.writeToPipe"
source: "fgl-topics/c_gws_XmlStaxWriter_writeToPipe.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.writeToPipe"
type: "concept"
---

# xml.StaxWriter.writeToPipe

> Sets the output stream of the StaxWriter object to a PIPE, and starts the streaming.

## Syntax

```
writeToPipe(
   name STRING )
```

1. name defines the command to start the
   PIPE that will get the resulting XML document.

## Usage

This method sets the output stream of the `StaxWriter` object to a PIPE command
specified in name, and starts the streaming.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
