---
title: "com.WebOperation.setInputEncoded"
source: "fgl-topics/c_gws_ComWebOperation_setInputEncoded.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class > WebOperation methods > com.WebOperation.setInputEncoded"
type: "concept"
---

# com.WebOperation.setInputEncoded

> Defines the encoding mechanism for Web Operation input parameters.

## Syntax

```
setInputEncoded(
   val INTEGER )
```

1. val defines an integer value defining the encoding
   mechanism to be used.

## Usage

When the parameter encoded is [TRUE](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") (1), the [SOAP Section
5](http://www.w3.org/TR/2000/NOTE-SOAP-20000508/#_Toc478383512) encoding mechanism is used. [FALSE](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.") (0) indicates the XML
Schema mechanism.

The XML Schema mechanism (`FALSE`) is not recommended.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
