---
title: "com.HttpRequest.clearAuthentication"
source: "fgl-topics/c_gws_ComHTTPRequest_clearAuthentication.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.clearAuthentication"
type: "concept"
---

# com.HttpRequest.clearAuthentication

> Removes user-defined authentication.

## Syntax

```
clearAuthentication()
```

## Usage

Removes user-defined authentication.

If an [authenticate](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") entry exists in the FGLPROFILE file, it will be used for authentication, even
if the user-defined authentication was removed.

> **Important:**
>
> The iOS HTTP stack doesn't provide a simple way to handle authentication. The
> GMI front-end uses the global iOS credential management system, that keeps credential values of
> previous request based on host and realm, until the keep-alive session is closed. Therefore, doing a
> `clearAuthentication()` on iOS devices does not take effect immediately.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
