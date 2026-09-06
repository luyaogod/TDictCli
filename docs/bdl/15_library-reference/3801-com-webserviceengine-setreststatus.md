---
title: "com.WebServiceEngine.SetRestStatus"
source: "fgl-topics/c_gws_ComWebServiceEngine_SetRestStatus.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.SetRestStatus"
type: "concept"
---

# com.WebServiceEngine.SetRestStatus

> Manages HTTP response status codes (200 - 299) for a REST high-level web service function.

## Syntax

```
com.WebServiceEngine.SetRestStatus(
   code INTEGER )
```

1. code defines a standard HTTP response status code (200 - 299) as defined in
   [RFC 2616
   Section 10.2](http://tools.ietf.org/html/rfc2616#section-10.2)

## Usage

You use this method in a REST web service function to dynamically change the HTTP return code. By
default, the REST operation returns 200 (success), 204 (if there is no body), and so on, but this
method can be called to set the return code to any value in the range from 200 to 299.

To support the Swagger and OpenAPI specification, Genero Web Services supports the "2XX" value,
which is a specific code that can be set in the [WSRetCode](../16_web-services/4830-wsretcode.md "Sets the HTTP success status code returned by the REST function.") attribute to define a range of response
codes between 200 and 299.

In the call to the `SetRestStatus()` method, the status code must be in the range
200 - 299; otherwise, the method throws an error code of
[-15582](4483-genero-bdl-errors.md).
> **Warning:**
>
> You cannot call the `SetRestStatus()` method outside a REST operation.

In the examples Example 1: managing user input and Example 2: managing zero different uses of the method are shown.

## Example 1: managing user input

In this example, the call to `SetRestStatus(205)` will send the code 205 to "Reset
Content" in the response:

```
IMPORT com
PUBLIC FUNCTION TestOne(x_one STRING)
  ATTRIBUTE (WSPOST,WSPath="/one",WSRetCode='2XX:HELLO')
  RETURNS (STRING)
  IF NOT x_one.equalsIgnoreCase("REGULAR") THEN
    CALL com.WebServiceEngine.SetRestStatus(205)
  END IF
  RETURN x_one
END FUNCTION
```

## Example 2: managing zero

In this example you code to trap when a parameter value is zero. In the call to the
`SetRestStatus()`method the status code 288 is returned:

```
IMPORT com
PUBLIC FUNCTION TestTwo(one INTEGER, two INTEGER )
  ATTRIBUTE (WSPOST,WSPath="/two",WSRetCode='2XX')
  RETURNS (INTEGER)
  IF one == 0 THEN
    CALL com.WebServiceEngine.SetRestStatus(288)
  END IF
  RETURN one + two
END FUNCTION
```

## Related links

**Related concepts**  

[Define functions in a module](../16_web-services/4723-define-functions-in-a-module.md "A GWS REST service is defined in a module.")

**Related reference**  

[High-level RESTful Web service attributes](../16_web-services/4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")
