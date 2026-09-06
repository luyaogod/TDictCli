---
title: "json.JSONReader.readFromPipe"
source: "fgl-topics/c_gws_jsonJSONReader_readFromPipe.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods > json.JSONReader.readFromPipe"
type: "concept"
---

# json.JSONReader.readFromPipe

> Sets the input stream of the JSONReader object to a PIPE.

## Syntax

```
readFromPipe(
   cmd STRING )
```

1. cmd defines the command to read from the PIPE.

## Usage

This method sets the input stream of the `JSONReader` object to a PIPE, where
cmd is the command to read from the PIPE.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
