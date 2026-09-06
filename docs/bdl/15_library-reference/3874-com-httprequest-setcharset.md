---
title: "com.HttpRequest.setCharset"
source: "fgl-topics/c_gws_ComHTTPRequest_setCharset.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setCharset"
type: "concept"
---

# com.HttpRequest.setCharset

> Defines the charset used when sending text or XML.

## Syntax

```
setCharset(
   charset STRING )
```

1. charset defines the character set to use.

## Usage

Defines the character set used when sending an HTTP request.

By default, no character set information will be transmitted in the HTTP header. This is also the
case when specifying `NULL` as parameter for this method.

If no character set is specified in HTTP headers, UTF-8 will implicitly be used
as defined by the HTTP standards.
> **Note:**
>
> **Default charset for web services:** The default character encoding for web service
> communication is UTF-8. If the request payload is in JSON format, you do not need to explicitly set
> the charset, as UTF-8 is assumed by default. For other content types (such as XML or plain text),
> you may need to explicitly set the charset to ensure correct encoding. For example: `CALL
> request.setCharset("UTF-8")`

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")
