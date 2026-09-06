---
title: "com.HttpRequest.setBodyChunk()"
source: "fgl-topics/c_gws_ComHTTPRequest_setBodyChunk.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setBodyChunk()"
type: "concept"
---

# com.HttpRequest.setBodyChunk()

> Disable chunk mode in HTTP 1.1 if request body size is greater than 32 KB.

## Syntax

```
setBodyChunk(
   val BOOLEAN )
```

1. val defines whether the sending of several chunks is
   enabled or disabled.

## Usage

The `setBodyChunk()` method allows you to disable chunk mode in the request in
HTTP 1.1. Use this method to disable sending the body in several chunks when the message body size
is greater than 32 KB. For example, you can disable chunk mode when sending data contained in a file
with [doFileRequest](3859-com-httprequest-dofilerequest.md "Performs the request by sending data contained in a file.").

Set val to `FALSE` to disable chunk mode. The default value is
`TRUE`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
