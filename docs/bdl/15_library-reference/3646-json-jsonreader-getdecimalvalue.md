---
title: "json.JSONReader.getDecimalValue"
source: "fgl-topics/c_gws_jsonJSONReader_getDecimalValue.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.getDecimalValue"
type: "concept"
---

# json.JSONReader.getDecimalValue

> Gets a decimal value of a JSONReader object.

## Syntax

```
getDecimalValue(
   de DECIMAL)
```

1. de is of [DECIMAL](../08_language-basics/0560-decimal-p-s.md "The DECIMAL data type is provided to handle large numeric values with exact decimal storage.") type.

## Usage

Use this method to get a decimal value of a [JSONReader](3639-the-json-jsonreader-class.md "The json.JSONReader class provides an interface compatible with JSON streaming that reads data in a JSON format from an input source.")
object, where de can have a value in a decimal form.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
