---
title: "com.WebService.createFault"
source: "fgl-topics/c_gws_ComWebService_createFault.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.createFault"
type: "concept"
---

# com.WebService.createFault

> Creates a global fault for a Web Service object.

## Syntax

```
createFault(
   fault RECORD ,
   encoded INTEGER  )
```

1. fault defines the Web Service fault.
2. encoded defines an integer value specifying the encoding
   mechanism.

## Usage

The `createFault()` method creates a global fault for this Web Service object.

The fault parameter can be of any type that defines the SOAP fault in a SOAP
response. In case of SOAP fault, the client for this Web Service will receive a variable
with the same structure.

When the parameter encoded is [TRUE](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") (1), the [SOAP Section
5](http://www.w3.org/TR/2000/NOTE-SOAP-20000508/#_Toc478383512) encoding mechanism is used. [FALSE](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.") (0) indicates the XML
Schema mechanism.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
