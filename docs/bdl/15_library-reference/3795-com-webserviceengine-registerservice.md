---
title: "com.WebServiceEngine.RegisterService"
source: "fgl-topics/c_gws_ComWebServiceEngine_RegisterService.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.RegisterService"
type: "concept"
---

# com.WebServiceEngine.RegisterService

> Registers a SOAP service in the engine.

## Syntax

```
com.WebServiceEngine.RegisterService(
   ws com.WebService )
```

1. ws defines the name of the SOAP service object to register.

## Usage

You use the `com.WebServiceEngine.RegisterService()` method to register the [`com.WebService`](3754-the-webservice-class.md "The com.WebService class provides an interface to create and manage Genero Web Services.") object passed as
parameter.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
