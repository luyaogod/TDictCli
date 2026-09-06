---
title: "com.WebService.setFeature"
source: "fgl-topics/c_gws_ComWebService_setFeature.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.setFeature"
type: "concept"
---

# com.WebService.setFeature

> Defines a feature for the current Web Service object.

## Syntax

```
setFeature(
   feature STRING,
   value STRING )
```

1. feature defines the name of the Web
   Service feature.
2. value defines the value of the feature.

## Usage

The `setFeature()` method defines a feature for the current Web Service object by
specifying a feature name and a value.

The features names are predefined. The second parameter must have a valid value for the specified
feature.

| Name | Description |
| --- | --- |
| `Soap1.1` | Defines whether the Web Service supports the SOAP 1.1 protocol. Default value is `FALSE`. |
| `Soap1.2` | Defines whether the Web Service supports the SOAP 1.2 protocol. Default value is `FALSE`. |
| `WS-Addressing1.0` | Defines whether the Web Service supports WS-Addressing 1.0. Valid values include:`TRUE` - The service supports WS-Addressing 1.0 and accepts requests without WS-Addressing.`REQUIRED` - The service supports WS-Addressing 1.0 and accepts only requests with WS-Addressing.`FALSE` - WS-Addressing 1.0 is disabled (Default). |

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
