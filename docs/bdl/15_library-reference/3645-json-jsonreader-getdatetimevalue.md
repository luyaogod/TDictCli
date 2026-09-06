---
title: "json.JSONReader.getDateTimeValue"
source: "fgl-topics/c_gws_jsonJSONReader_getDateTimeValue.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.getDateTimeValue"
type: "concept"
---

# json.JSONReader.getDateTimeValue

> Gets a datetime value of the current JSON node.

## Syntax

```
getDateTimeValue(
   dt DATETIME )
```

1. dt is of [DATETIME](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")
   type.

## Usage

Use this method to get a string containing a datetime value in the current JSON node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
