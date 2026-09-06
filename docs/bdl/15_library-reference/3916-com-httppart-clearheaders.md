---
title: "com.HttpPart.clearHeaders"
source: "fgl-topics/c_gws_ComHTTPPart_clearHeaders.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > HttpPart methods > com.HttpPart.clearHeaders"
type: "concept"
---

# com.HttpPart.clearHeaders

> Remove all headers from the HTTP part.

## Syntax

```
clearHeaders()
```

## Usage

This method removes all headers from an HTTP multipart object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
CALL req.clearHeaders()
```

## Related links

**Related concepts**  

[com.HttpPart.removeHeader](3925-com-httppart-removeheader.md "Remove the header of given name from the current HttpPart object.")
