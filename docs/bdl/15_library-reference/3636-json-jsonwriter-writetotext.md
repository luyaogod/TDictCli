---
title: "json.JSONWriter.writeToText"
source: "fgl-topics/c_gws_jsonJSONWriter_writeToText.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods > json.JSONWriter.writeToText"
type: "concept"
---

# json.JSONWriter.writeToText

> Sets the output stream of the JSONWriter object to a TEXT large object, and starts the streaming.

## Syntax

```
writeToText(
   txt TEXT )
```

1. txt defines a TEXT lob located in memory for streaming.

## Usage

This method sets the output stream of the `JSONWriter` object to a text large
object in memory, specified in txt, and starts the streaming.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
