---
title: "com.HttpResponse.getStatusCode"
source: "fgl-topics/c_gws_ComHTTPResponse_getStatusCode.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpResponse class > HttpResponse methods > com.HttpResponse.getStatusCode"
type: "concept"
---

# com.HttpResponse.getStatusCode

> Returns the HTTP status code.

## Syntax

```
getStatusCode()
  RETURNS INTEGER
```

## Usage

The `getStatusCode()` method returns the status code for the HTTP response.

When the returned HTTP status code is 401 or 407, authorization is required.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[com.HttpRequest.setAuthentication](3869-com-httprequest-setauthentication.md "Defines the user login and password to authenticate to the server.")
