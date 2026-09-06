---
title: "xml.StaxWriter.comment"
source: "fgl-topics/c_gws_XmlStaxWriter_comment.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.comment"
type: "concept"
---

# xml.StaxWriter.comment

> Writes a XML comment to the StaxWriter stream.

## Syntax

```
comment(
   comment STRING )
```

1. comment defines the data in the XML comment, or NULL.

## Usage

This method writes a XML comment to the stream.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
