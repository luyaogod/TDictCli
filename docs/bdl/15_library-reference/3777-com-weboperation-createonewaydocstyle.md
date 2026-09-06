---
title: "com.WebOperation.CreateOneWayDOCStyle"
source: "fgl-topics/c_gws_ComWebOperation_CreateOneWayDOCStyle.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class > WebOperation methods > com.WebOperation.CreateOneWayDOCStyle"
type: "concept"
---

# com.WebOperation.CreateOneWayDOCStyle

> Creates a new Web Operation object with One-Way Document style.

## Syntax

```
com.WebOperation.CreateOneWayDOCStyle(
   function STRING,
   operation STRING,
   inputVar any-type )
  RETURNS com.WebOperation
```

1. function defines the
   program function to be called to process the XML operation.
2. operation defines the
   XML operation.
3. inputVar defines the program variable for the input of
   the operation. The any-type can be any Genero BDL variable
   type: `RECORD`, `ARRAY`, or simple data type.
   It can return a `NULL` value.

## Usage

This method creates a One-Way DOC style `com.WebOperation` object, where
function is the name of the program function that is executed to process the XML
operation.

The function name must be a string literal, not a string variable, due to [operation publication restrictions](../05_upgrading/0273-web-services-changes.md).

There is no output parameter to be returned to the client.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
