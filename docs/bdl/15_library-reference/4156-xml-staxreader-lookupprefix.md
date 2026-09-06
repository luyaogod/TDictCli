---
title: "xml.StaxReader.lookupPrefix"
source: "fgl-topics/c_gws_XmlStaxReader_lookupPrefix.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.lookupPrefix"
type: "concept"
---

# xml.StaxReader.lookupPrefix

> Looks up the prefix associated with a given namespace URI, starting from the current XML node the StaxReader cursor is pointing to.

## Syntax

```
lookupPrefix(
   ns STRING )
  RETURNS STRING
```

1. ns defines the namespace URI to look for. It cannot be NULL.

## Usage

Use this method to return the prefix associated with the namespace specified by
ns at the current XML node. It returns the prefix, or NULL if there is none.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
