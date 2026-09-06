---
title: "com.HttpRequest.setAutoReply"
source: "fgl-topics/c_gws_ComHTTPRequest_setAutoReply.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setAutoReply"
type: "concept"
---

# com.HttpRequest.setAutoReply

> Defines the auto reply option for response methods.

## Syntax

```
setAutoReply(
   val INTEGER )
```

1. val defines auto-reply when 1 ([`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.")).

## Usage

The `setAutoReply()` method defines whether `getResponse()` or
`getAsyncResponse()` will automatically perform another HTTP GET request if the
response contains HTTP Authentication, Proxy Authentication or HTTP redirect data.

Available for GET method and the HTTP HEAD method.

The default is 1 ([`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.")).

> **Important:**
>
> On iOS devices, `setAutoReply()` is ignored for redirection in
> synchronous requests. The iOS HTTP stack does not allow you to set an auto reply option when doing
> synchronous requests.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[com.HttpRequest.getResponse](3867-com-httprequest-getresponse.md "Waits for and returns the response produced by one of request methods.")

[com.HttpRequest.getAsyncResponse](3866-com-httprequest-getasyncresponse.md "Retrieves an asynchronous response produced by one of the request methods.")

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")
