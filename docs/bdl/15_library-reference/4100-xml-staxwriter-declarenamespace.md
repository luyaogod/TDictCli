---
title: "xml.StaxWriter.declareNamespace"
source: "fgl-topics/c_gws_XmlStaxWriter_declareNamespace.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.declareNamespace"
type: "concept"
---

# xml.StaxWriter.declareNamespace

> Binds a namespace URI to a prefix, and forces the output of the XML namespace definition to the StaxWriter stream.

## Syntax

```
declareNamespace(
   prefix STRING,
   ns STRING )
```

1. prefix defines the prefix to bind to the URI, cannot be NULL.
2. ns defines the URI to bind to the default namespace, cannot be NULL.

## Usage

This method binds a namespace URI to a prefix, and forces the output of the XML namespace
definition to the `StaxWriter` stream.

The stream must point to a `START_ELEMENT`, and the prefix scope is the
current `START_ELEMENT` / `END_ELEMENT` pair.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
