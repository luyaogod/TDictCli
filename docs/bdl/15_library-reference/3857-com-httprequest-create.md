---
title: "com.HttpRequest.Create"
source: "fgl-topics/c_gws_ComHTTPRequest_Create.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.Create"
type: "concept"
---

# com.HttpRequest.Create

> Creates a new HttpRequest object from a URL.

## Syntax

```
com.HttpRequest.Create(
   url STRING )
  RETURNS com.HttpRequest
```

1. url defines the URL for the HTTP
   request.

## Usage

Creates a `com.HttpRequest` object by providing a mandatory URL with HTTP or HTTPS
as protocol.

The url parameter can be an identifier of a URL mapping with an optional
`alias://` prefix. See [FGLPROFILE
Configuration](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") for more details about URL mapping with aliases, and for proxy and security
configuration.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
