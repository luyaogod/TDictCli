---
title: "com.HttpResponse.getServerCookies"
source: "fgl-topics/c_gws_ComHTTPResponse_getServerCookies.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getServerCookies"
type: "concept"
---

# com.HttpResponse.getServerCookies

> Returns all cookies set as response from a server.

## Syntax

```
getServerCookies(
   cookies RECORD  )
```

1. cookies defines the dynamic array of all cookies sent as a response from a
   server. See [WSHelper.WSServerCookiesType](../16_web-services/4926-wshelper-wsservercookiestype.md "The WSHelper.WSServerCookiesType defines a dynamic array for storing cookies.") for more information regarding
   `WSServerCookiesType`.

## Usage

This methods returns all cookies set as response from a server, in a dynamic array. All cookies
are returned, even those that have already expired.

If the method `setAutoCookies()` is enabled, cookies are automatically send back
if they have not expired. In other words, the `setAutoCookies(true)` handles all
cookies for you; however, you can still check them with the `getServerCookies()`
method if needed.

If `setAutoCookies(false)` (the default), you must handle the cookies and not send
the expired ones.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

This method may raise exception [-15535](4483-genero-bdl-errors.md) if the given dynamic array is not a dynamic array of RECORD with seven (7) members as
defined by `WSHelper.WSServerCookiesType`.

## Related links

**Related concepts**  

[com.HttpRequest.setAutoCookies](3870-com-httprequest-setautocookies.md "Enables automatic cookie management for a given request.")
