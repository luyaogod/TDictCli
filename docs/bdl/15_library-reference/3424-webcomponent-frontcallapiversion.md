---
title: "webcomponent.frontCallAPIVersion"
source: "fgl-topics/c_fgl_frontcall_webcomponent_frontcallapiversion.html"
breadcrumb: "Library reference > Built-in front calls > Web component front calls > webcomponent.frontCallAPIVersion"
type: "concept"
---

# webcomponent.frontCallAPIVersion

> Returns the API version of web component front-end calls.

## Syntax

```
ui.Interface.frontCall("webcomponent", "frontCallAPIVersion",
  [],[result])
```

1. result - Holds the API version for web component front calls.

## Usage

This front call can be used to check the API version for the web component front calls.

If the API version changes, you must adapt the code to the expected front call API
implemented for the web components.

The value returned by this front call is typically a version number, such as
`1.0`, `1.1`, etc.

## Example

```
FUNCTION wc_api_version()
    DEFINE vers STRING
    TRY
       CALL ui.Interface.frontCall("webcomponent","frontCallAPIVersion",[],[vers])
       -- we can safely call "webcomponent" "call" in the code
       RETURN vers
    CATCH
       -- we cannot call the "webcomponent" functions...
       RETURN 0
    END TRY
END FUNCTION
```
