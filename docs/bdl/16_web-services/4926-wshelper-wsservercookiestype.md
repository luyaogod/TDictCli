---
title: "WSHelper.WSServerCookiesType"
source: "fgl-topics/c_gws_wshelper_WSServerCookiesType.html"
breadcrumb: "Web services > Reference > WSHelper library > WSHelper: library > WSHelper.WSServerCookiesType"
type: "concept"
---

# WSHelper.WSServerCookiesType

> The WSHelper.WSServerCookiesType defines a dynamic array for storing cookies.

## WSServerCookiesType

```
TYPE WSServerCookiesType DYNAMIC ARRAY OF RECORD
    name      STRING, # Cookie name
    value     STRING, # Cookie value
    path      STRING, # Cookie path (or null)
    domain    STRING, # Cookie domain (or null)
    expires   DATETIME YEAR TO SECOND, # Cookie expiration date (or null)
    httpOnly  BOOLEAN,
    secure    BOOLEAN
    sameSite  STRING # Lax (default), Strict, or None (requires secure)
  END RECORD
```

1. `name` is the name of the cookie to be set. This field is mandatory. It will be
   URL-encoded on the wire.
2. `value` is the value of the cookie to be set. This field is mandatory. It will be
   URL-encoded on the wire.
3. `path` is the main path the cookie has to be set for. Any path containing that
   main path will then return the cookie. If no name is set, the cookie will be set by the path the
   client has provided.
4. `domain` is the domain (hostname) or sub-domain (for example:
   .strasbourg.4js.com) the cookie will be set on client side. If not set, the domain will be the
   hostname provided by the client.
5. `expires` is a `DATETIME YEAR TO SECOND` (on local time) from when
   that cookie will expire, and thus not be sent by the client anymore. If `NULL`, the
   cookie is called session cookie and will be sent as long as the client keeps the session open.
6. `http_only` is set to `TRUE` if the cookie is only for the HTTP
   layer, `FALSE` if cookie can be accessible in JavaScript.
7. `secure` is set to `TRUE` if cookie must only be sent in HTTPS,
   `FALSE` for HTTP and HTTPS.
8. `sameSite` is the value that defines security from cross-site attacks when
   cookies are used. Three values are possibles for sending cookies to other sites: "Strict", "Lax", or
   "None". The default is "Lax". If `NULL`, the default "Lax" is used.

## Usage

The `WSHelper.WSServerCookiesType` defines a dynamic array for storing cookies
that the server sends to the client.

For examples using this record, see [com.HttpServiceRequest.setResponseCookies](../15_library-reference/3841-com-httpservicerequest-setresponsecookies.md "Allows the server to return cookies to be set on the client application sending the request.") and [Examples using com.HttpServiceRequest methods](../15_library-reference/3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.").
