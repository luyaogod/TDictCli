---
title: "xml.StaxWriter.dtd"
source: "fgl-topics/c_gws_XmlStaxWriter_dtd.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.dtd"
type: "concept"
---

# xml.StaxWriter.dtd

> Writes a DTD to the StaxWriter stream.

## Syntax

```
dtd(
   dtd STRING )
```

1. dtd defines a string representing a valid DTD, cannot be NULL.

## Usage

This method writes a document type definition (DTD) for the `StaxWriter` stream,
where `dtd` represents a valid DTD.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
