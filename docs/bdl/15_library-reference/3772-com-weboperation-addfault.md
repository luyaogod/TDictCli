---
title: "com.WebOperation.addFault"
source: "fgl-topics/c_gws_ComWebOperation_addFault.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class > WebOperation methods > com.WebOperation.addFault"
type: "concept"
---

# com.WebOperation.addFault

> Adds a fault to the current Web Operation definition.

## Syntax

```
addFault(
   fault any-type,
   wsaaction STRING )
```

1. fault defines the program variable that has details of the fault. The any-type can be any Genero BDL variable
   type: `RECORD`, `ARRAY`, or simple data type.
2. wsaaction defines the type of action.

## Usage

Adds a fault the Web Operation can throw during operation processing, where
fault is any variable previously created as fault of the
`com.WebService` object, and wsaaction the WS-Addressing action
identifier if WS-Addressing is supported. If WS-Addressing is not supported, pass NULL as second
parameter.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
