---
title: "com.HttpRequest.setMaximumResponseLength"
source: "fgl-topics/c_gws_ComHTTPRequest_setMaximumResponseLength.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setMaximumResponseLength"
type: "concept"
---

# com.HttpRequest.setMaximumResponseLength

> Defines the maximum size in Kbytes of a response.

## Syntax

```
setMaximumResponseLength(
   length INTEGER )
```

1. length defines the maximum size in Kbytes.

## Usage

The `setMaximumResponseLength()` method sets the maximum authorized size in Kbytes
of the whole response (including headers, body and all control characters), before a break.

The value of -1 means no limit.
> **Note:**
>
> Setting the maximum response length is ignored for
> synchronous requests in a Genero Mobile for iOS (GMI) app. The iOS HTTP stack does not allow you to
> set a maximum response length when doing synchronous requests.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
