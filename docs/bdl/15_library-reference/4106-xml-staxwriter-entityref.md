---
title: "xml.StaxWriter.entityRef"
source: "fgl-topics/c_gws_XmlStaxWriter_entityRef.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.entityRef"
type: "concept"
---

# xml.StaxWriter.entityRef

> Writes a XML EntityReference to the StaxWriter stream.

## Syntax

```
entityRef(
   name STRING )
```

1. name defines the name of the entity, cannot be NULL.

## Usage

This method writes a XML EntityReference, specified by name, to the
`StaxWriter` stream.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
