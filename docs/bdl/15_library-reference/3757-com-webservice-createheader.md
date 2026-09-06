---
title: "com.WebService.createHeader"
source: "fgl-topics/c_gws_ComWebService_createHeader.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.createHeader"
type: "concept"
---

# com.WebService.createHeader

> Defines the header for the Web Service object.

## Syntax

```
createHeader(
   header RECORD,
   encoded INTEGER )
```

1. header defines the header for the Web Service object.
2. encoded defines an integer value specifying the encoding
   mechanism.

## Usage

The `createHeader()` method creates a global header for the current Web Service
object.

The Web Service header is defined by the first parameter. This will define SOAP headers exchanged
by the client and server.

When the parameter encoded is [TRUE](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") (1), the [SOAP Section
5](http://www.w3.org/TR/2000/NOTE-SOAP-20000508/#_Toc478383512) encoding mechanism is used. [FALSE](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.") (0) indicates the XML
Schema mechanism.

> **Important:**
>
> Since Web Services headers are generally in Document Style, we recommend to
> set the encoded parameter to 0.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
