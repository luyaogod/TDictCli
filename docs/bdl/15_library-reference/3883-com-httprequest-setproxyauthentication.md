---
title: "com.HttpRequest.setProxyAuthentication"
source: "fgl-topics/c_gws_ComHTTPRequest_setProxyAuthentication.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setProxyAuthentication"
type: "concept"
---

# com.HttpRequest.setProxyAuthentication

> Define the login and password to use for proxy authentication.

## Syntax

```
setProxyAuthentication(
   login STRING,
   password STRING,
   scheme STRING,
   realm STRING )
```

1. login defines the login to use for
   authentication.
2. pass specifies the password to use for
   authentication.
3. scheme defines the method to be used
   during authentication. This is optional; it can be set to NULL.
4. realm defines the realm. This is
   optional; it can be set to NULL.

## Usage

Defines the login and password to use for proxy authentication for the current
`HttpRequest` request.

The scheme parameter defines the method to be used
during authentication. The supported values for the scheme parameter are
`Anonymous`, `Basic`, `Digest`, and
`NTLM`. The default is `Anonymous`.

If [setProxy()](3882-com-httprequest-setproxy.md "Configure the proxy URL.") is called, the login and
password will be used to authenticate against the proxy set by this API, regardless of whether a
proxy is configured in FGLPROFILE.

If `setProxy()` has not been called and a proxy is configured in FGLPROFILE, the
login and password set by this method will be used to authenticate against the proxy defined by
FGLPROFILE.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
