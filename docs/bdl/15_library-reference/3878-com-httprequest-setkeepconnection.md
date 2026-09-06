---
title: "com.HttpRequest.setKeepConnection"
source: "fgl-topics/c_gws_ComHTTPRequest_setKeepConnection.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setKeepConnection"
type: "concept"
---

# com.HttpRequest.setKeepConnection

> Defines whether a connection is kept open if a new request occurs.

## Syntax

```
setKeepConnection(
   keep INTEGER )
```

1. keep defines if the connection is kept.

## Usage

The `setKeepConnection()` method defines whether the connection stays open when a
new HTTP request occurs.

The default is 0 ([`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.")).

> **Note:**
>
> If the server uses the Windows® NT (New Technology) LAN Manager (NTLM) protocol, authentication
> requires the request to keep the connection open via a call to
> `setKeepConnection(TRUE)`, otherwise authentication will fail.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
