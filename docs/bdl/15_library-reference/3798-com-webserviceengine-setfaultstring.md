---
title: "com.WebServiceEngine.SetFaultString"
source: "fgl-topics/c_gws_ComWebServiceEngine_SetFaultString.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.SetFaultString"
type: "concept"
---

# com.WebServiceEngine.SetFaultString

> Defines the description of a SOAP Fault.

## Syntax

```
com.WebServiceEngine.SetFaultString(
   str STRING )
```

1. str defines the description of the
   fault.

## Usage

The `com.WebServiceEngine.SetFaultString()` class method defines a user SOAP Fault
description to be returned to the client, where str contains the description of
the fault.

This method must be called inside a Web Service function.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
