---
title: "xml.StaxWriter.setPrefix"
source: "fgl-topics/c_gws_XmlStaxWriter_setPrefix.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > xml.StaxWriter methods > xml.StaxWriter.setPrefix"
type: "concept"
---

# xml.StaxWriter.setPrefix

> Binds a namespace URI to a prefix.

## Syntax

```
setPrefix(
   prefix STRING,
   ns STRING )
```

1. prefix defines the prefix to be bind to the URI, cannot be NULL.
2. ns defines the namespace URI to be bind to the prefix, cannot be NULL.

## Usage

Use this method to bind a namespace URI, specified in `ns`, to a prefix defined in
prefix.

The prefix scope is the current `START_ELEMENT` / `END_ELEMENT` pair.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
