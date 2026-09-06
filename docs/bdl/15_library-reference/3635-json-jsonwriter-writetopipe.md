---
title: "json.JSONWriter.writeToPipe"
source: "fgl-topics/c_gws_jsonJSONWriter_writeToPipe.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods > json.JSONWriter.writeToPipe"
type: "concept"
---

# json.JSONWriter.writeToPipe

> Sets the output stream of the JSONWriter object to a PIPE, and starts the streaming.

## Syntax

```
writeToPipe(
   cmd STRING )
```

1. cmd defines the command to start the PIPE that will get the resulting JSON
   document.

## Usage

This method sets the output stream of the `JSONWriter` object to a PIPE command
specified in cmd, and starts the streaming.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
