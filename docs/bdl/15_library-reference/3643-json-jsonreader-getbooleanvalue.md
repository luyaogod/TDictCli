---
title: "json.JSONReader.getBooleanValue"
source: "fgl-topics/c_gws_jsonJSONReader_getBooleanValue.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.getBooleanValue"
type: "concept"
---

# json.JSONReader.getBooleanValue

> Gets a boolean value of a JSONReader object.

## Syntax

```
getBooleanValue(
   b BOOLEAN)
```

1. b is of boolean type.

## Usage

Use this method to get a boolean value of a [JSONReader](3639-the-json-jsonreader-class.md "The json.JSONReader class provides an interface compatible with JSON streaming that reads data in a JSON format from an input source.")
object, where b can be `TRUE` (integer 1) or
`FALSE` (integer 0).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
