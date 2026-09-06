---
title: "com.WebOperation.setComment"
source: "fgl-topics/c_gws_ComWebOperation_setComment.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class > WebOperation methods > com.WebOperation.setComment"
type: "concept"
---

# com.WebOperation.setComment

> Sets the comment for the Web Operation object.

## Syntax

```
setComment(
   comment STRING )
```

1. comment defines the comment to be set.

## Usage

The `setComment()` method defines a comment for the current Web Operation
object.

The comment will appear in the WSDL of the service.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
