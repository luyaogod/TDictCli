---
title: "xml.StaxReader.hasText"
source: "fgl-topics/c_gws_XmlStaxReader_hasText.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.hasText"
type: "concept"
---

# xml.StaxReader.hasText

> Checks whether the StaxReader cursor points to a node with a text value.

## Syntax

```
hasText()
  RETURNS INTEGER
```

## Usage

Use this method to check if the StaxReader cursor is pointing to a node with a text value. It
Returns `TRUE` if the current XML node has a text value, `FALSE`
otherwise. This method returns `TRUE` for `CHARACTERS`,
`SPACE`, `CDATA`, `COMMENT`,
`ENTITY_REFERENCE`, and `DTD` nodes, `FALSE` for all
other nodes.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
