---
title: "xml.StaxWriter.processingInstruction"
source: "fgl-topics/c_gws_XmlStaxWriter_processingInstruction.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.processingInstruction"
type: "concept"
---

# xml.StaxWriter.processingInstruction

> Writes a XML ProcessingInstruction to the StaxWriter stream

## Syntax

```
processingInstruction(
   target STRING,
   data STRING )
```

1. target defines the target of the Processing Instruction, cannot be NULL.
2. data defines the data of the Processing Instruction, or NULL.

## Usage

This method writes a XML ProcessingInstruction to the `StaxWriter` stream, where
target is the target of the Processing Instruction, which cannot be
`NULL`. data is the data of the Processing Instruction, or
`NULL`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
