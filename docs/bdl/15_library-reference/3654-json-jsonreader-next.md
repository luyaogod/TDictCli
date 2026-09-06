---
title: "json.JSONReader.next"
source: "fgl-topics/c_gws_jsonJSONReader_next.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.next"
type: "concept"
---

# json.JSONReader.next

> Moves the JSONReader cursor to the next node.

## Syntax

```
next()
```

## Usage

Use this method to move the JSONReader cursor to the next node in the stream.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example: Reading a JSON file](3661-example-reading-a-json-file.md "This example shows two ways to read a JSON file.")
