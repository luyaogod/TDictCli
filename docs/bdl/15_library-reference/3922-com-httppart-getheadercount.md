---
title: "com.HttpPart.getHeaderCount"
source: "fgl-topics/c_gws_ComHTTPPart_getHeaderCount.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.getHeaderCount"
type: "concept"
---

# com.HttpPart.getHeaderCount

> Retrieve the number of headers for the current HTTP part.

## Syntax

```
getHeaderCount()
  RETURNS INTEGER
```

1. pos specifies the position of the multipart header.

## Usage

Use this method to retrieve the number of HTTP headers for current HTTP part.

The method returns the number of HTTP headers for current part.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
LET count = httppart.getHeaderCount(3)
```

## Related links

**Related concepts**  

[com.HttpPart.setHeader](3926-com-httppart-setheader.md "Sets a named HTTP multipart header using a string value.")
