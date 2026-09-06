---
title: "com.WebService.setComment"
source: "fgl-topics/c_gws_ComWebService_setComment.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.setComment"
type: "concept"
---

# com.WebService.setComment

> Defines the comment for the Web Service object.

## Syntax

```
setComment(
   comment STRING )
```

1. comment defines the comment to be set.

## Usage

The `setComment()` method defines the comment associated with a
`com.WebService` object.

The comment will be used when generating the WSDL file, as defined by the standard.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
