---
title: "com.WebServiceEngine.SetFaultCode"
source: "fgl-topics/c_gws_ComWebServiceEngine_SetFaultCode.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.SetFaultCode"
type: "concept"
---

# com.WebServiceEngine.SetFaultCode

> Get a handle for an incoming HTTP service request.

## Syntax

```
com.WebServiceEngine.SetFaultCode(
   code STRING,
   ns STRING )
```

1. code defines the fault code.
2. ns defines the namespace of the fault
   code.

## Usage

The `com.WebServiceEngine.SetFaultCode()` class method defines a user SOAP Fault
code to be returned to the client, where code is the mandatory SOAP
Fault code and ns is the mandatory code namespace.

This method must be called inside a Web Service operation.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
