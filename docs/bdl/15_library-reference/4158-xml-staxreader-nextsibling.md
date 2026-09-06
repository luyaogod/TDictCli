---
title: "xml.StaxReader.nextSibling"
source: "fgl-topics/c_gws_XmlStaxReader_nextSibling.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.nextSibling"
type: "concept"
---

# xml.StaxReader.nextSibling

> Moves the StaxReader cursor to the immediate next sibling XML Element of the current node, skipping all its child nodes.

## Syntax

```
nextSibling()
```

## Usage

Use this method to move the StaxReader cursor to the next sibling of the current XML Element node
in the stream, skipping all its child nodes. The cursor points to the parent end tag if there are no
siblings.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
