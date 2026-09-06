---
title: "xml.StaxWriter.attribute"
source: "fgl-topics/c_gws_XmlStaxWriter_attribute.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.attribute"
type: "concept"
---

# xml.StaxWriter.attribute

> Writes a XML attribute to the StaxWriter stream.

## Syntax

```
attribute(
   name STRING,
   value STRING )
```

1. name defines the local name of the XML attribute. It cannot be NULL.
2. value defines the value of the XML attribute. It cannot be NULL.

## Usage

Attributes can only be written on the StaxWriter stream if it points to a
`START_ELEMENT` or an `EMPTY_ELEMENT`, otherwise the operation fails
with an exception. This method can only be called after a `startElement`,
`startElementNS`, `emptyElement`, `emptyElementNS`, or
`attribute` and `attributeNS`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
