---
title: "xml.StaxWriter.close"
source: "fgl-topics/c_gws_XmlStaxWriter_close.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.close"
type: "concept"
---

# xml.StaxWriter.close

> Closes the StaxWriter streaming, and releases all associated resources.

## Syntax

```
close()
```

## Usage

This method closes the stream.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
