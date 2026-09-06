---
title: "xml.StaxReader.nextTag"
source: "fgl-topics/c_gws_XmlStaxReader_nextTag.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.nextTag"
type: "concept"
---

# xml.StaxReader.nextTag

> Moves the StaxReader cursor to the next XML open or end tag

## Syntax

```
nextTag()
```

## Usage

Use this method to move the StaxReader cursor to the next XML open or end tag. The cursor points
to the end of the document if there is no next XML open or end tag.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
