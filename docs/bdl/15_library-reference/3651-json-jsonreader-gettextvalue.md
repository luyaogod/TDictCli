---
title: "json.JSONReader.getTextValue"
source: "fgl-topics/c_gws_jsonJSONReader_getTextValue.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.getTextValue"
type: "concept"
---

# json.JSONReader.getTextValue

> Gets a text value of the current JSON node.

## Syntax

```
getTextValue(
   txt TEXT )
```

1. txt is of [TEXT](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") type.

## Usage

Use this method to return a string containing text in the current JSON node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
