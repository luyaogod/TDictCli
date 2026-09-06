---
title: "json.JSONReader.getSize"
source: "fgl-topics/c_gws_jsonJSONReader_getSize.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.getSize"
type: "concept"
---

# json.JSONReader.getSize

> Returns the number of bytes read by the JSONReader in the current stream.

## Syntax

```
getSize()
  RETURNS INTEGER
```

## Usage

This method returns the number of bytes read by the JSONReader since the start of the current
stream. It returns `-1` if no bytes have been written.

For example, this method can be used to give an indication of progress when processing a large
file.
