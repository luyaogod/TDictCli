---
title: "com.HttpPart.removeHeader"
source: "fgl-topics/c_gws_ComHTTPPart_removeHeader.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.removeHeader"
type: "concept"
---

# com.HttpPart.removeHeader

> Remove the header of given name from the current HttpPart object.

## Syntax

```
removeHeader(
   name STRING )
```

1. name is the name of the header to remove.

## Usage

With an HTTP multipart object, the `"Content-Type"` header may be
optional for each part:

```
https://tools.abc.org/html/rfc1234#section-1.1
```

A request might be rejected by some providers if the "Content-Type" header is set for a
part. This method allows you to remove headers from an HttpPart object based on the
header name.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
CALL req.removeHeader("Content-Type")
```

## Related links

**Related concepts**  

[com.HttpPart.clearHeaders](3916-com-httppart-clearheaders.md "Remove all headers from the HTTP part.")
