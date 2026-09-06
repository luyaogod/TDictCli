---
title: "com.HttpRequest.setAuthentication"
source: "fgl-topics/c_gws_ComHTTPRequest_setAuthentication.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setAuthentication"
type: "concept"
---

# com.HttpRequest.setAuthentication

> Defines the user login and password to authenticate to the server.

## Syntax

```
setAuthentication(
   login STRING,
   pass STRING,
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

The `setAuthentication()` method defines the mandatory user login and password to
authenticate to the server.

> **Note:**
>
> If the server uses the Windows® NT (New Technology) LAN Manager (NTLM) protocol, authentication
> requires the request to keep the connection open via a call to
> `setKeepConnection(TRUE)`, otherwise authentication will fail.

The scheme parameter defines the method to be used
during authentication. The supported values for the scheme parameter are
`Anonymous`, `Basic`, `Digest`, and
`NTLM`. The default is `Anonymous`.

An optional realm can be specified.

With `Anonymous`, `Digest`, or `NTLM`
authentication, you must re-send the request if you get a 401 or 407 HTTP return code (authorization
required). In the case of `NTLM`, the request must be re-sent one more time.

If a user-defined authentication is set and there is an [authenticate](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") entry for
this URL in the FGLPROFILE file, the user-defined authentication has priority.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
