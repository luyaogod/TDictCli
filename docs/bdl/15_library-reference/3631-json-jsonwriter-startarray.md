---
title: "json.JSONWriter.startArray"
source: "fgl-topics/c_gws_jsonJSONWriter_startArray.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods > json.JSONWriter.startArray"
type: "concept"
---

# json.JSONWriter.startArray

> Writes a JSON start array token to the JSONWriter stream.

## Syntax

```
startArray()
```

## Usage

This method writes a JSON start array token to the `JSONWriter` stream. The
`startArray` method opens a new scope and sets the stream to a
`startArray`. Writing the corresponding [endArray](3622-json-jsonwriter-endarray.md "Writes a JSON end array tag to the JSONWriter stream.") causes the scope to be closed.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example 1: Writing to a JSON file](3638-example-1-writing-to-a-json-file.md "This example shows two ways to write a JSON file.")
