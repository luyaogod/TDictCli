---
title: "WSName"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSName.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes related to function parameters > WSName"
type: "concept"
---

# WSName

> Specifies an alternative name for a parameter or return value in the REST message.

## Syntax

```
WSName=" description "
```

Where:

1. description provides an alternative name for the input parameter or
   return value it is applied to.

`WSName` is an optional attribute.

## Usage

Use this attribute to specify names for parameters in a REST function. In general, you do not
need to use `WSName`. Consider using it to improve the readability of input
parameters or to format the data returned in the response.

In a HTTP request when you reference a parameter set with `WSName`, the
`WSName` value is used in the call to the function. See the examples of
setting this attribute with different parameter types in Table 1.

| Parameter type | Function parameters with WSName | Description |
| --- | --- | --- |
| Header | `PUBLIC FUNCTION myfunc(p1 INTEGER ATTRIBUTES(WSHeader, WSName = "myCustomHeader")` | The parameter `p1` holds a custom header. The REST function expects a header named "myCustomHeader", not "p1". |
| Query | `PUBLIC FUNCTION myfunc(p2 STRING ATTRIBUTES(WSQuery, WSName = "country")` | On the URL to call the REST function, you must use "country" in the query string, not "p2", for example `/users?country=france` |
| Cookie | `PUBLIC FUNCTION myfunc(p3 INTEGER ATTRIBUTES(WSCookie, WSName = "token")` | The parameter `p3` holds a cookie. The REST function expects a cookie named "token", not "p3". |

You use `WSName` to format the output data. It allows you to override the
default names ("rv0", "rv1", and so on) that is otherwise output at runtime. By specifying a
more recognizable name for the data, you improve the readability of the output.

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

## Example 2: Providing a user-friendly parameter name

You would consider setting the `WSName` attribute clause of a parameter if the
parameter name is not obvious or user-friendly.

For example, if the parameter references a field name in a database table that is abbreviated,
or could be ambiguous in your function, `WSName` provides you with an option to
provide a clearer name.

```
PUBLIC FUNCTION getUsers(reg STRING ATTRIBUTES(WSQuery, WSOptional, WSName = "region"),
     ccode STRING ATTRIBUTES(WSHeader, WSName = "country") )
 # ...
END FUNCTION
```

In the URL that calls the function, the value specified by `WSName` must be
used: http://myhost/gas/ws/r/myGroup/myXcf/MyService/users?region=alsace

## Example 3: Formatting output data

In this sample REST function, a set of users is returned. In the returns clause of the function
the `WSName` attribute is set on the array returning the user records.
`"Users"` will appear as the root path for the output in JSON, or the root
tag for the output in XML.

```
TYPE profileType RECORD
     id INT,
     name VARCHAR(100),
     email VARCHAR(100)
     # ...
   END RECORD

PUBLIC FUNCTION getUsers()
  ATTRIBUTES (WSGet,
              WSPath = "/users",
              WSDescription = "Get users resource"
              )
  RETURNS (DYNAMIC ARRAY ATTRIBUTES(WSName = "Users") OF
        profileType ATTRIBUTES(WSMedia = "application/json, application/xml"))

    DEFINE regionList DYNAMIC ARRAY OF profileType
    DEFINE i INTEGER = 1

     DECLARE usersCurs CURSOR FOR SELECT id, name FROM users ORDER BY name
     FOREACH usersCurs INTO regionList[i].*
       LET i = i+1
     END FOREACH
     CALL regionList.deleteElement(regionList.getLength())
     # Remove the empty element implied by reference in FOREACH loop

    RETURN regionList
END FUNCTION
```

## Related links

**Related concepts**  

[Customize XML serialization with WSName and XMLName](4742-customize-xml-data-output.md "Use XML serialization attributes to customize serialization at runtime and improve the readability of the XML output.")

[XMLName](5037-xmlname.md "XMLName")
