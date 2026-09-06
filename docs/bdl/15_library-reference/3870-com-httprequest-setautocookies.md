---
title: "com.HttpRequest.setAutoCookies"
source: "fgl-topics/c_gws_ComHTTPRequest_setAutoCookies.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setAutoCookies"
type: "concept"
---

# com.HttpRequest.setAutoCookies

> Enables automatic cookie management for a given request.

## Syntax

```
setAutoCookies( val INTEGER )
```

1. val defines the cookie management flag.

## Usage

Set to 1 ([`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.")) to activate
automatic cookie management.

If set to 1 ([`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.")), it activates
the automatic cookies management for that request. For the `com.HttpRequest`
instance:

- If the server returns session cookies (with no expiration date), they will be automatically sent
  again for the next request.
- If the server returns persistent cookies (with an expiration date), those cookies will be
  registered globally for the current fglrun process, and any other HttpRequest
  (including the current one) will automatically send those cookies according to the request path and
  domain, as long as the expiration date has not expired.

The [autocookiesmanagement](3804-webserviceengine-options.md) option of the [com.WebServiceEngine.SetOption](3799-com-webserviceengine-setoption.md "Sets an option for the Web Service engine.") method
activates the automatic cookies management for any HttpRequest.

The [maximumpersistentcookies](3804-webserviceengine-options.md) option of the [com.WebServiceEngine.SetOption](3799-com-webserviceengine-setoption.md "Sets an option for the Web Service engine.") method sets
the maximum number of cookies that can be handled by an fglrun process.

Default value is 0 ([`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.")).

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
