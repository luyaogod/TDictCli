---
title: "json.JSONWriter.setProperty"
source: "fgl-topics/c_gws_jsonJSONWriter_setProperty.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods > json.JSONWriter.setProperty"
type: "concept"
---

# json.JSONWriter.setProperty

> Sets a property of a JSONWriter object.

## Syntax

```
setProperty(
   name STRING)
```

1. name is a string identifying the element in the name-value pair of the JSON
   string.

## Usage

Use this method to set a property of a [JSONWriter](3618-the-json-jsonwriter-class.md "The json.JSONWriter class provides an interface compatible with JSON streaming that writes data in a JSON format to an output source.")
object, where name identifies the element in the name-value pair of the JSON
string.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Example 1: Writing to a JSON file](3638-example-1-writing-to-a-json-file.md "This example shows two ways to write a JSON file.")
