---
title: "xml.StaxReader.getAttributeNamespace"
source: "fgl-topics/c_gws_XmlStaxReader_getAttributeNamespace.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.getAttributeNamespace"
type: "concept"
---

# xml.StaxReader.getAttributeNamespace

> Returns the namespace URI of a XML attribute defined at a given position on the current XML node, or NULL.

## Syntax

```
getAttributeNamespace(
   index INTEGER )
  RETURNS STRING
```

1. index defines the position of the attribute to return (index starts at
   1).

## Usage

This method returns the value of a namespace URI attribute at index position
on the current XML node, or NULL.

This method is only valid on a
`START_ELEMENT` node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
