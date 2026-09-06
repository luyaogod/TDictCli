---
title: "com.HttpServiceRequest.readFormEncodedRequest"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_readFormEncodedRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > HttpServiceRequest methods > com.HttpServiceRequest.readFormEncodedRequest"
type: "concept"
---

# com.HttpServiceRequest.readFormEncodedRequest

> Returns the string of a GET request with UTF-8 conversion option.

## Syntax

```
readFormEncodedRequest(
   utf8 INTEGER )
  RETURNS STRING
```

1. utf8 specifies whether the returned string is converted from UTF‑8 to the
   current character set. The parameter is an `INTEGER` flag: `1` ([TRUE](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.")) enables conversion, and `0` ([FALSE](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.")) disables it.

## Usage

The `readFormEncodedRequest()` method returns the query of a `POST`
"`application/x-www-form-urlencoded`" request or the query string of a
`GET` request. The method decodes the query string using HTML4 or XFORM rules. If
utf8 is set to `1` ([TRUE](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.")), the
decoded string is then converted from UTF‑8 to the current character set.

The method handles UTF-8 encoded special characters (`&` or
`=`) in parameter values. When the method decodes a query string, any encoded special
characters (`&` or `=`) that appear within parameter values or
names (not as separators) are decoded and then doubled in the output. The separators
(`&` and `=`) that delimit parameters are preserved as single
characters.

**Example:**

If the request received is:

```
name1=value1&na%26me2=val%3Due2
```

where:

- `%26` is the URL-encoded value for `&` (ampersand)
- `%3D` is the URL-encoded value for `=` (equals sign)

The method will return:

```
name1=value1&na&&me=val==ue2
```

Breaking this down:

- `name1=value1` – first parameter (name and value separated by `=`)
- `&` – separator between parameters
- `na&&me2` – second parameter name where `%26` (encoded `&`) is decoded and doubled to `&&`
- `=` – separator between parameter name and value
- `val==ue2` – second parameter value where `%3D` (encoded `=`) is decoded and doubled to `==`

If the utf8 parameter is set to `1` ([TRUE](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.")), the decoded query string is translated from UTF-8 to the current
character set. This may lead to a conversion error.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Related links

**Related concepts**  

[Examples using com.HttpServiceRequest methods](3845-examples-httpservicerequest.md "These examples use methods of the com.HttpServiceRequest class.")

[Examples: Using the com.HttpPart class](3927-examples-httppart.md "Examples using methods of the com.HttpPart class.")
