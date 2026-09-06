---
title: "json.JSONReader.getValue"
source: "fgl-topics/c_gws_jsonJSONReader_getValue.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.getValue"
type: "concept"
---

# json.JSONReader.getValue

> Gets the value of a JSONReader object.

## Syntax

```
getValue() RETURNS STRING
```

## Usage

Use this method to get the value of a [JSONReader](3639-the-json-jsonreader-class.md "The json.JSONReader class provides an interface compatible with JSON streaming that reads data in a JSON format from an input source.")
object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example: Reading a JSON file](3661-example-reading-a-json-file.md "This example shows two ways to read a JSON file.")
