---
title: "com.WebService.publishOperation"
source: "fgl-topics/c_gws_ComWebService_publishOperation.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.publishOperation"
type: "concept"
---

# com.WebService.publishOperation

> Publishes a Web Operation.

## Syntax

```
publishOperation(
   op com.WebOperation,
   role STRING )
```

1. op defines the Web Operation object.
2. role identifies uniquely the Web Operation.

## Usage

The `publishOperation()` method publishes the Web Operation specified by the
`com.WebOperation` object passed as parameter.

The role identifies the operation, if several operations have the same name,
by setting the SOAPAction HTTP header. Usually this parameter is set to NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
