---
title: "com.WebOperation.initiateSession"
source: "fgl-topics/c_gws_ComWebOperation_initiateSession.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class > WebOperation methods > com.WebOperation.initiateSession"
type: "concept"
---

# com.WebOperation.initiateSession

> Defines the Web Operation as session initiator.

## Syntax

```
initiateSession(
   ok INTEGER  )
```

1. ok must have an integer value of 1
   to define a session initiator.

## Usage

Pass the parameter with the value of 1 to `initiateSession()` in order to define
the current Web Operation as a session initiator.

A new session must be instantiated in this operation, and must be returned to the client via the
state variable defined at service creation.

This method works only for [stateful](3759-com-webservice-createstatefulwebservice.md "Creates a new object to implement a stateful Web service.") web services.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
