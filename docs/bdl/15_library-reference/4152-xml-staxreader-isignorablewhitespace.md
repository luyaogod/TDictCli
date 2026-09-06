---
title: "xml.StaxReader.isIgnorableWhitespace"
source: "fgl-topics/c_gws_XmlStaxReader_isIgnorableWhitespace.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.isIgnorableWhitespace"
type: "concept"
---

# xml.StaxReader.isIgnorableWhitespace

> Checks whether the StaxReader cursor points to ignorable whitespace.

## Syntax

```
isIgnorableWhitespace()
  RETURNS INTEGER
```

## Usage

Use this method to check if the StaxReader cursor is pointing to a text node with ignorable
whitespace. It returns `TRUE` if there is ignorable whitespace,
`FALSE` otherwise.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
