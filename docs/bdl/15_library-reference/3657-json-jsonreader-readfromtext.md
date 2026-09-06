---
title: "json.JSONReader.readFromText"
source: "fgl-topics/c_gws_jsonJSONReader_readFromText.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.readFromText"
type: "concept"
---

# json.JSONReader.readFromText

> Sets the output stream of the JSONReader object to a TEXT large object.

## Syntax

```
readFromText(
   txt TEXT )
```

1. txt defines a TEXT large object located in memory.

## Usage

This method sets the input stream of the `JSONReader` object to a text large
object in memory, specified in txt.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
