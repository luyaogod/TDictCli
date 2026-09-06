---
title: "xml.StaxWriter.emptyElementNS"
source: "fgl-topics/c_gws_XmlStaxWriter_emptyElementNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.emptyElementNS"
type: "concept"
---

# xml.StaxWriter.emptyElementNS

> Writes an empty namespace qualified XML element to the StaxWriter stream.

## Syntax

```
emptyElementNS(
   name STRING,
   ns STRING )
```

1. name defines the local name of the XML empty element, cannot be NULL.
2. ns defines the namespace URI of the XML empty element, cannot be NULL.

## Usage

This method writes an empty namespace qualified XML element to the `StaxWriter`
stream.

If namespace URI has not been bound to a prefix with one of
the functions `setPrefix`,
`setDefaultNamespace`,
or `declareDefaultNamespace`, the operation fails with an exception.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
