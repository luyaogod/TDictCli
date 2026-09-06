---
title: "WSRetCode"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSRetCode.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > Error handling attributes > WSRetCode"
type: "concept"
---

# WSRetCode

> Sets the HTTP success status code returned by the REST function.

## Syntax

```
WSRetCode= "{ code | code:description } "
```

Where `WSRetCode` is a string enclosed in quotations and where:

1. code can be:
   - an HTTP code representing successful status in the range 200 to 399
   - or the value "2XX", which is a specific code supported by the Swagger and OpenAPI specification
     to represent all response codes between 200 and 299, allowing the REST web service to return a code
     in this range dynamically at runtime with a call to the [SetRestStatus](../15_library-reference/3801-com-webserviceengine-setreststatus.md "Manages HTTP response status codes (200 - 299) for a REST high-level web service function.") method.
2. description provides a description.

If you use the code option without a description, the default description
defined in [RFC
2616](http://tools.ietf.org/html/rfc2616) is returned to the client. If you use the code option with a
description, there is no space before or after the colon (`:`).

`WSRetCode` is an optional attribute.

## Usage

You use `WSRetCode` in the `ATTRIBUTES()` clause of the REST
function to specify the HTTP protocol status in the response. The HTTP code and HTTP description you
provide is sent in the HTTP protocol response.

For examples setting the `code` with the "2XX" value, go to [com.WebServiceEngine.SetRestStatus](../15_library-reference/3801-com-webserviceengine-setreststatus.md "Manages HTTP response status codes (200 - 299) for a REST high-level web service function.").

## Example using WSRetCode in REST function

In this trivial
example, instead of "200 OK", the HTTP response is changed to "215 correct".

```
PUBLIC FUNCTION Add(
    a INTEGER ATTRIBUTES(WSQuery),
    b INTEGER ATTRIBUTES(WSQuery) )
  ATTRIBUTES(WSGet,
             WSRetCode="215:correct" )
  RETURNS INTEGER
    RETURN a + b
END FUNCTION
```

![In the image the HTTP response is changed to 215 correct instead of 200 OK](../_images/rest_response_with_wsretcode.png)

*HTTP response using WSRetCode*
