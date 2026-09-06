---
title: "com.HttpServiceRequest.setResponseCookies"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_setResponseCookies.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.setResponseCookies"
type: "concept"
---

# com.HttpServiceRequest.setResponseCookies

> Allows the server to return cookies to be set on the client application sending the request.

## Syntax

```
setResponseCookies(
   cookies WSHelper.WSServerCookiesType )
```

1. cookies defines a `WSHelper.WSServerCookiesType` record for
   the cookies to set.

## Usage

Allows the server to return cookies to be set on the client application sending the request.

A runtime error ( for example, -15566) may be raised if the cookies
`sameSite` attribute is not set correctly. For instance, you can not set it to "None"
if secure is not set to `TRUE` as well. This raises the error
"`SameSite None value requires Secure to be set`".

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Setting cookies example

```
IMPORT FGL WSHelper

DEFINE cookies WSHelper.WSServerCookiesType

# Set first cookie named 'CookieName'
LET cookies[1].name = "CookieName"
LET cookies[1].value = "AnyValue"
LET cookies[1].expires = now + INTERVAL (5) MINUTE TO MINUTE
LET cookies[1].secure = TRUE
LET cookies[1].sameSite = "Lax" 

# Set second cookie named 'SecondCookie'
LET cookies[2].name = "SecondCookie
LET cookies[2].value = "AnotherValue"

# Set all cookies defined in the cookie array
CALL req.setResponseCookies(cookies)
CALL req.sendTextResponse(200, NULL, "Hello world")
```

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
