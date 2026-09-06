---
title: "com.WebOperation.setOutputAction"
source: "fgl-topics/c_gws_ComWebOperation_setOutputAction.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class > WebOperation methods > com.WebOperation.setOutputAction"
type: "concept"
---

# com.WebOperation.setOutputAction

> Sets the WS-Addressing action identifier of the output operation.

## Syntax

```
setOutputAction(
   action STRING )
```

1. action defines the WSA action identifier.

## Usage

When WS-Addressing is enabled, this method defines the WS-Addressing action identifier of the
output operation.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
