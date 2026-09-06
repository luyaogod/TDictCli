---
title: "json.JSONWriter.endObject"
source: "fgl-topics/c_gws_jsonJSONWriter_endObject.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods > json.JSONWriter.endObject"
type: "concept"
---

# json.JSONWriter.endObject

> Writes a JSON end object token to the JSONWriter stream.

## Syntax

```
endObject()
```

## Usage

This method writes an end object token to the `JSONWriter` corresponding to the
last [startObject](3633-json-jsonwriter-startobject.md "Writes a JSON start object token to the JSONWriter stream.") token.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example 1: Writing to a JSON file](3638-example-1-writing-to-a-json-file.md "This example shows two ways to write a JSON file.")
