---
title: "com.WebServiceEngine.SetFaultDetail"
source: "fgl-topics/c_gws_ComWebServiceEngine_SetFaultDetail.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.SetFaultDetail"
type: "concept"
---

# com.WebServiceEngine.SetFaultDetail

> Defines the published SOAP Fault.

## Syntax

```
com.WebServiceEngine.SetFaultDetail(
   fault any-type )
```

1. fault defines the program variable that has details of the fault. The any-type can be any Genero BDL variable
   type: `RECORD`, `ARRAY`, or simple data type.

## Usage

The `com.WebServiceEngine.SetFaultDetail()` class method defines the published
SOAP Fault to be returned to the client when the operation has finished, where
fault is one of the published variables defined as Fault for that operation.

This method must be called inside a Web Service operation.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
