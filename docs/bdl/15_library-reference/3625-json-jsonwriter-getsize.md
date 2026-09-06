---
title: "json.JSONWriter.getSize"
source: "fgl-topics/c_gws_jsonJSONWriter_getSize.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods > json.JSONWriter.getSize"
type: "concept"
---

# json.JSONWriter.getSize

> Returns the number of bytes written by the JSONWriter in the current stream.

## Syntax

```
getSize()
  RETURNS INTEGER
```

## Usage

This method returns the number of bytes written by the JSONWriter since the start of the current
stream. It returns `-1` if no bytes have been written.

For example, this method can be used to give an indication of progress when processing a large
file.

## Related links

**Related concepts**  

[Example 1: Writing to a JSON file](3638-example-1-writing-to-a-json-file.md "This example shows two ways to write a JSON file.")
