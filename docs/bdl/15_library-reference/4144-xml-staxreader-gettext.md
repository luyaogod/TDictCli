---
title: "xml.StaxReader.getText"
source: "fgl-topics/c_gws_XmlStaxReader_getText.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.getText"
type: "concept"
---

# xml.StaxReader.getText

> Returns as a string the value of the current XML node, or NULL.

## Syntax

```
getText()
  RETURNS STRING
```

## Usage

Use this method to return a string containing text in the current XML node, or
`NULL` if there is none.

This method is only valid on `CHARACTERS`, `CDATA`,
`SPACE`, `COMMENT`, `DTD`, and
`ENTITY_REFERENCE` nodes. For an `ENTITY_REFERENCE`, this method
returns the replacement value, or NULL if none. See [StaxReader Event Types](4167-staxreader-event-types.md "Event types of the xml.StaxReader class.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
