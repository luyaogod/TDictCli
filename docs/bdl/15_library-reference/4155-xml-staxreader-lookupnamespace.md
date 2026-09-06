---
title: "xml.StaxReader.lookupNamespace"
source: "fgl-topics/c_gws_XmlStaxReader_lookupNamespace.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > xml.StaxReader methods > xml.StaxReader.lookupNamespace"
type: "concept"
---

# xml.StaxReader.lookupNamespace

> Looks up the namespace URI associated with a given prefix starting from the current XML node the StaxReader cursor is pointing to.

## Syntax

```
lookupNamespace(
   prefix STRING )
  RETURNS STRING
```

1. prefix defines the prefix to look
   for; if NULL the default namespace URI will be returned..

## Usage

Use this method to return the namespace URI associated with the prefix specified by
prefix at the current XML node. It returns a string with the namespace URI
associated with the prefix, or `NULL` if there is none.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
