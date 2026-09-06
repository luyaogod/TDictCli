---
title: "com.HttpRequest.doFormEncodedRequest"
source: "fgl-topics/c_gws_ComHTTPRequest_doFormEncodedRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.doFormEncodedRequest"
type: "concept"
---

# com.HttpRequest.doFormEncodedRequest

> Performs an "application/x-www-form-urlencoded forms" encoded query.

## Syntax

```
doFormEncodedRequest(
   query STRING,
   utf8 INTEGER )
```

1. query defines a list of name/value pairs separated by
   an `&`.
2. utf8 defines if the query string is UTF-8 encoded. The parameter is an
   `INTEGER` flag: `1` ([TRUE](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.")) enables UTF-8 encoding, and `0` ([FALSE](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.")) disables it.

## Usage

The `doFormEncodedRequest()` method performs a request with an
"application/x-www-form-urlencoded forms" encoded query.

Supported methods are GET and POST.

The query string is a list of name/value pairs separated by an ampersand
(`&`). For
example:

```
name1=value1&name2=value2&name3=value3
```

> **Note:**
>
> If you need to URL-encode the separator characters `&` and
> `=`, double them as following:
> `na&&me=va==lue`.

If the utf8 parameter is 1 (`TRUE`), the query string is encoded in UTF-8 as specified in [XForms1.0](http://www.w3.org/TR/xforms/#serialize-urlencode), otherwise
in ASCII as specified in [HTML4](http://www.w3.org/TR/1999/REC-html401-19991224/interact/forms.html#h-17.13.4.1).

This HTTP request method is non-blocking. It returns immediately
after the call. Use the [com.HttpRequest.getResponse](3867-com-httprequest-getresponse.md "Waits for and returns the response produced by one of request methods.") method, to perform a
synchronous HTTP request, suspending the program flow until the response returns from the server. If
the program must keep going, use the [com.HttpRequest.getAsyncResponse](3866-com-httprequest-getasyncresponse.md "Retrieves an asynchronous response produced by one of the request methods.")
method, to check if a response is available.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
