---
title: "json.JSONReader.getIntervalValue"
source: "fgl-topics/c_gws_jsonJSONReader_getIntervalValue.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.getIntervalValue"
type: "concept"
---

# json.JSONReader.getIntervalValue

> Gets a datetime value of the current JSON node.

## Syntax

```
getIntervalValue(
   i INTERVAL)
```

1. dt is of [INTERVAL](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.")
   type.

## Usage

Use this method to get a string containing time interval value in the current JSON node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
