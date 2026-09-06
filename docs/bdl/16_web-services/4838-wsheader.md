---
title: "WSHeader"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSHeader.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes set on both input and output > WSHeader"
type: "concept"
---

# WSHeader

> Defines a custom HTTP header for a parameter or return value.

## Syntax

```
WSHeader
```

Where:

`WSHeader` in an `ATTRIBUTES()` clause of an input parameter
or return value of the function defines a HTTP header.

`WSHeader` is an optional attribute.

## Usage

Use the `WSHeader` attribute to define a custom header in an input parameter
or in a return value. Custom headers are commonly used to provide information; however, you
can also use them to pass data that implements logic on the server or client side.

Typically, you use `WSHeader` to pass data related to the service in the
form of an HTTP header. For example, an input parameter defined with
`ATTRIBUTES(WSHeader, WSName="Authorization")` sets the
standard HTTP authorization header.

You can set the `WSHeader` attribute on parameters defined as primitive
types, records, arrays, and simple dictionary type.

`WSHeader` supports the OpenAPI default styles: `style:
simple` and `explode: false`. Table 1
shows how the header (`X-MyHeader`) is serialized when the parameter
(`id`) is a primitive type, an array, record, and a dictionary type. The
dictionary and record types are defined as objects according to the OpenAPI specification.
For further information on OpenAPI serialization, see the [Parameter serialization](http://swagger.io/docs/specification/serialization/) page.

| style | explode | URI template | Primitive value X-MyHeader = 5 | Array X-MyHeader = [3, 4, 5] | Object X-MyHeader = {"role": "admin", "firstName": "Alex"} |
| --- | --- | --- | --- | --- | --- |
| simple | false | {id} | X-MyHeader:5 | X-MyHeader: 3,4,5 | X-MyHeader: role,admin,firstName,Alex |

The GWS and fglrestful supports
serialization and deserialization of parameters according to defaults set by the OpenAPI
specification. If a value contains special characters as defined by [RFC3986](http://tools.ietf.org/html/rfc3986#section-2.2) (for example, `:/?#[]@!$&'()*+,;=`), the GWS engine
serializes them with percent-encoding in the parameter. For
example:

```
DEFINE b DYNAMIC ARRAY OF STRING = ["O,ne", "Two", "Three", NULL, "Five"]
```

The parameter contains:
`O%2Cne,Two,Three,,Five`. The server and the client must deserialize
"`O%2Cne`" to "`O,ne`" for insertion in Genero
BDL.

Use `WSName` with
`WSHeader` to name headers; otherwise the GWS gives the headers default
names ("rv0", "rv1", etc) at runtime.

## Example: WSHeader and WSName in input parameter

In the sample REST function, the Genero Application Server (GAS) environment variable
"`X-FourJs-Environment-Variable-REMOTE_ADDR`" that contains the IP address
of the remote client is passed in a header.

As the GAS variable has hyphens ("-"), which are not allowed in Genero BDL variables, the
`WSName` attribute is set on the `ip_addr` parameter to
identify "`X-FourJs-Environment-Variable-REMOTE_ADDR`" in the HTTP request.

The `ip_addr` parameter is BDL friendly and this is used in the function.
The parameter is also optional as defined by the attribute `WSOptional`.
> **Tip:**
>
> The [WSContext](4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.") attribute, if
> set, can provide your service with access to environment variables set by the GAS.

```
PUBLIC FUNCTION getRemoteAddress (ip_addr STRING
   ATTRIBUTES(WSHeader, 
              WSOptional, 
              WSName = "X-FourJs-Environment-Variable-REMOTE_ADDR") )
   ATTRIBUTES (WSGet, 
               WSPath = "/users/ip",
               WSDescription = "Get remote address of the client")
   RETURNS (INTEGER ATTRIBUTES(WSHeader), STRING)
   DEFINE ip STRING
     LET ip = ip_addr
     IF ip IS NULL THEN
        LET ip = "Got no remote address."
     ELSE
        LET ip = SFMT("Hello there, you're at %1",ip )
     END if
     RETURN 3, ip
END FUNCTION
```

## Example responses in header and body

In this sample REST function data is returned in a header and in
the message body. The function's `RETURNS` clause has two return values:

- An integer is returned in a header. It is specified with the `WSHeader`
  attribute.
- A string is returned in the body. It is specified without an attribute.

```
PUBLIC FUNCTION help()
  ATTRIBUTES (WSGet,
              WSPath = "/help")
  RETURNS (INTEGER ATTRIBUTE(WSHeader, WSDescription = "Reference number"),
    STRING)
    RETURN 3, "Hello world"
END FUNCTION
```

![Sample output of the HTTP response](../_images/rest_response_head_body_output.png)

*Output of the HTTP response*

In the output the header is given a default name, "rv0",
at runtime. You can change default header naming via the [WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") attribute.
> **Warning:**
>
> **Standard response headers**
>
> Setting
> a standard [HTTP header](http://www.iana.org/assignments/message-headers/message-headers.xhtml) on a response must be handled with care,
> especially for those that define the response body such as `Content-Type`, or
> `Content-Encoding`. Make sure what you define with `WSName` does not
> conflict with what is specified in the OpenAPI documentation for the service.

## Example: WSHeader returns a record

In this sample REST function data is sent in a custom header in the
`RETURNS` clause of the function on a user-defined record type
("`T_RECORD`"). The string ("Return record in header: ..." ) is returned in
the response body.

```
PUBLIC TYPE t_user RECORD
         user_id INTEGER,
         user_name VARCHAR(50)
     END RECORD

PUBLIC FUNCTION getUserInfo(
    p_user_id INTEGER ATTRIBUTES(WSHeader, WSOptional))
    ATTRIBUTES(WSGet,
        WSPath = "/v1/ok",
        WSDescription = "Return user record in header",
        WSRetCode = "202:Accepted",
        WSThrows = "400:Bad Request,401:Unauthorized,500:Internal Server Error")
    RETURNS(STRING ATTRIBUTES(WSMedia = "application/json"),
        t_user ATTRIBUTES(WSHeader))

    DEFINE u t_user

    SELECT * INTO u.* FROM user WHERE user_id = p_user_id

    RETURN "User info is in header...", u

END FUNCTION
```

## Example: WSHeader returns data in a dictionary

In this sample REST function a list of available bikes is returned in a `DICTIONARY
OF STRING` set in the custom header `dict`.

The function's input parameter thisBike passes a value representing a
model of bike in a header. The `WSGet` attribute is set to request details
about thisBike from the service. The dictionary `dict` is
queried for details of thisBike.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors.
The `message` field of the `userError` variable is set if the
bike is not found, and a call to `SetRestError()`
returns the message defined in `WSThrows` for the error.

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION getBikeDetails(
    thisBike STRING ATTRIBUTES(WSHeader))
    ATTRIBUTES(WSGet,
        WSPath = "/bikes",
        WSDescription = "Get bike details. Return list of available bikes",
        WSRetCode = "202:Accepted",
        WSThrows = "404:@userError")
    RETURNS(STRING ATTRIBUTES(WSMedia = "application/xml, application/json"),
        DICTIONARY ATTRIBUTES(WSHeader, WSName = "dict") OF STRING)

    DEFINE b STRING
    DEFINE dict DICTIONARY OF STRING
    DEFINE keys DYNAMIC ARRAY OF STRING
    DEFINE i INT

    LET dict["YZF"] = "250cc"
    LET dict["Z650"] = "652cc"
    LET dict["MT-07"] = "689cc"
    
    LET b = "Bike details"
    LET keys = dict.getKeys()
    DISPLAY "Number bikes: ", dict.getLength()
    FOR i = 1 TO keys.getLength()
        DISPLAY i, " ", keys[i], " ", dict[keys[i]]
    END FOR

    IF dict.contains(thisBike) THEN
        LET b = SFMT("This bike has: %1 ", dict[thisBike])
    ELSE
        LET userError.message = SFMT("Could not find '%1' ", thisBike)
        CALL com.WebServiceEngine.SetRestError(404, userError)
    END IF
    RETURN b, dict

END FUNCTION
```

## Related links

**Related concepts**  

[WSErrorHeader](4831-wserrorheader.md "Defines a custom HTTP header to include in an error response.")
