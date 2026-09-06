---
title: "com.WebOperation.addOutputHeader"
source: "fgl-topics/c_gws_ComWebOperation_addOutputHeader.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class > WebOperation methods > com.WebOperation.addOutputHeader"
type: "concept"
---

# com.WebOperation.addOutputHeader

> Adds an output header for the current Web Operation definition.

## Syntax

```
addOutputHeader(
   header any-type  )
```

1. header defines the program variable that has details for
   the header. The any-type can be any Genero BDL variable
   type: `RECORD`, `ARRAY`, or simple data type.

## Usage

This method adds a header to the Web Operation object for input parameters.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
