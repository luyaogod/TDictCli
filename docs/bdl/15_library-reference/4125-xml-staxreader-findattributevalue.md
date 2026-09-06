---
title: "xml.StaxReader.findAttributeValue"
source: "fgl-topics/c_gws_XmlStaxReader_findAttributeValue.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.findAttributeValue"
type: "concept"
---

# xml.StaxReader.findAttributeValue

> Returns the value of an XML attribute of a given name and/or namespace.

## Syntax

```
findAttributeValue(
   name STRING,
   ns STRING )
  RETURNS STRING
```

1. name defines the name of the
   attribute to retrieve. It cannot be NULL.
2. ns defines the namespace URI of the
   attribute to retrieve, or NULL if the attribute is not namespace-qualified.

## Usage

This method returns the value of an XML attribute of a given name and/or namespace on the current
XML node, where name is the name of the attribute and ns is
the namespace or is NULL if not namespace-qualified.

This method is only valid on a
`START_ELEMENT` node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
