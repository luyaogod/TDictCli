---
title: "com.WebServiceEngine.GetOption"
source: "fgl-topics/c_gws_ComWebServiceEngine_GetOption.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.GetOption"
type: "concept"
---

# com.WebServiceEngine.GetOption

> Returns the value of a Web Service engine option.

## Syntax

```
com.WebServiceEngine.GetOption(
   option STRING )
  RETURNS STRING
```

1. option defines the [option](3804-webserviceengine-options.md) to be queried.

## Usage

The `com.WebServiceEngine.GetOption()` class method returns the current value of
the given Web Services engine option.

See [WebServiceEngine options](3804-webserviceengine-options.md) for the supported options.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
