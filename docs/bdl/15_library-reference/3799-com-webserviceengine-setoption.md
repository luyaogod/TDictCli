---
title: "com.WebServiceEngine.SetOption"
source: "fgl-topics/c_gws_ComWebServiceEngine_SetOption.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.SetOption"
type: "concept"
---

# com.WebServiceEngine.SetOption

> Sets an option for the Web Service engine.

## Syntax

```
com.WebServiceEngine.SetOption(
   optionName STRING,
   optionValue STRING )
```

1. optionName defines the [option](3804-webserviceengine-options.md) to set.
2. optionValue defines the value of the
   option to set.

## Usage

The `com.WebServiceEngine.SetOption()` class method configures the Web Services
engine with options.

See [WebServiceEngine options](3804-webserviceengine-options.md) for the supported options.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
