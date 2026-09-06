---
title: "json.JSONWriter.setBooleanValue"
source: "fgl-topics/c_gws_jsonJSONWriter_setBooleanValue.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods > json.JSONWriter.setBooleanValue"
type: "concept"
---

# json.JSONWriter.setBooleanValue

> Sets a boolean value of a JSONWriter object.

## Syntax

```
setBooleanValue(
   value BOOLEAN)
```

1. value is of boolean type.

## Usage

Use this method to set a boolean value of a [JSONWriter](3618-the-json-jsonwriter-class.md "The json.JSONWriter class provides an interface compatible with JSON streaming that writes data in a JSON format to an output source.")
object, where value can be `TRUE` (integer 1) or
`FALSE` (integer 0).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example 1: Writing to a JSON file](3638-example-1-writing-to-a-json-file.md "This example shows two ways to write a JSON file.")
