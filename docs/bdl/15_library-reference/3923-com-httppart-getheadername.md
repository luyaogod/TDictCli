---
title: "com.HttpPart.getHeaderName"
source: "fgl-topics/c_gws_ComHTTPPart_getHeaderName.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.getHeaderName"
type: "concept"
---

# com.HttpPart.getHeaderName

> Retrieve the name of an HTTP multipart header as a string, where the multipart header is specified by its position.

## Syntax

```
getHeaderName(
   pos INTEGER )
  RETURNS STRING
```

1. pos specifies the position of the multipart header.

## Usage

Use this method to retrieve the name of an HTTP multipart header as a string.

The method returns the name for the header part specified by its position.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
LET str = httppart.getHeaderName(3)
```

## Related links

**Related concepts**  

[com.HttpPart.getHeader](3921-com-httppart-getheader.md "Returns a named HTTP multipart header as a string.")
