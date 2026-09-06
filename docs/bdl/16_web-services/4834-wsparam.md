---
title: "WSParam"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSParam.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes set on input parameters only > WSParam"
type: "concept"
---

# WSParam

> Maps a parameter to a value in the resource path template.

## Syntax

```
WSParam
```

`WSParam` is an optional attribute.

## Usage

You use this attribute for path templating to specify values that are provided at runtime in
function parameters. You set the `WSParam` attribute in the
`ATTRIBUTES()` clause of the parameter.

For each Genero BDL function parameter with a `WSParam` attribute, there must be a
matching template path in the `WSPath`. fglcomp checks to ensure that there is one template value per
parameter, otherwise compilation [error-9111](../15_library-reference/4483-genero-bdl-errors.md) is thrown.

You can set the `WSParam` attribute on parameters defined as records and arrays.

`WSParam` supports the default `style: simple`
and `explode: false`. [Table 1](4834-wsparam.md) shows how the path (`users`) is serialized when the parameter
(`id`) is a primitive type, an array, or a record. For further information on OpenAPI
serialization, see the [Parameter serialization](http://swagger.io/docs/specification/serialization/) page.

| style | explode | Primitive value id = 5 | Array id = [3, 4, 5] | Record = {"role": "admin", "lastName": "Smith"} |
| --- | --- | --- | --- | --- |
| simple | false | /users/5 | /users/3,4,5 | /users/role,admin,lastName,Smith |

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

## Example path templating with WSParam

In this sample REST function a user resource is returned. In the function's
`p_user_id` parameter the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") specifies the user to
return.

The `WSPath` attribute sets the resource identifier or endpoint of the URI
to the resource. It contains the function's `p_user_id` parameter enclosed in
curly brackets `{}` as the path template.

The client application needs to provide an appropriate parameter value when making a call
to the function. For example, the variable part of the path (`p_user_id`) is
replaced by the integer value `/4` in the URL:

http://myhost/gas/ws/r/myGroup/myXcf/MyService/accounts/4

```
TYPE accountType RECORD
         user_id INTEGER,
         user_name VARCHAR(50)
END RECORD

PUBLIC FUNCTION getAccountById(p_user_id INTEGER ATTRIBUTES(WSParam))
  ATTRIBUTES(WSGet,
             WSPath = "/accounts/{p_user_id}",
             WSDescription = "Returns an account record",
             WSThrows = "404:not found"
            )
  RETURNS accountType ATTRIBUTES(WSName = "Account",WSMedia = "application/json") 
    DEFINE p_accountRec accountType
    
    SELECT * INTO p_accountRec.* FROM users WHERE id = p_user_id

    RETURN p_accountRec
END FUNCTION
```

## Example WSParam set on parameter defined as array

In this sample REST function user records are returned. In the function the attribute
`WSParam` is set on the parameter `idlist`. This parameter
is an array type containing a list of user ids. The data is passed by the client in the URL
(`3,4,5` in the example) for insertion into the Genero BDL parameter:

http://myhost/gas/ws/r/myGroup/myXcf/MyService/users/3,4,5

The `WSPath` attribute contains the `idlist` parameter
enclosed in curly brackets `{}` as the path template.

The client application needs to provide appropriate parameter values when making a call to
the function. The GWS deserializes the user ids sent by the client in the URL.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors. In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. The `SQLERRMESSAGE` is set to the
`message` field of the `userError` variable, and a call to
`SetRestError()` returns the message defined in `WSThrows`
for the error.

```
IMPORT com

TYPE profileType RECORD
     id INTEGER,
     name VARCHAR(100),
     email VARCHAR(255)
     # ...
   END RECORD

PUBLIC FUNCTION getUserById(idlist DYNAMIC ARRAY ATTRIBUTES(WSParam) OF INTEGER)
  ATTRIBUTES(WSGet, WSPath = "/users/{idlist}",
            WSDescription = "Returns one or more user records",
            WSThrows = "404:Not Found,406:@userError,500:Internal Server Error"
  )
  RETURNS DYNAMIC ARRAY ATTRIBUTES(WSName="Users") OF profileType
  
  DEFINE p_userRec profileType
  DEFINE p_recList DYNAMIC ARRAY OF profileType
  DEFINE i INTEGER
  DEFINE j INTEGER = 1 
     TRY
       FOR i = 1 TO idlist.getLength()
          INITIALIZE p_userRec TO NULL
          SELECT * INTO p_userRec.* FROM users WHERE users.id = idlist[i]
          IF p_userRec.id IS NOT NULL THEN
            LET p_recList[j] = p_userRec
            LET j = j + 1
          END IF
       END FOR
     CATCH
       LET userError.message = SFMT("Error in SQL execution: %1 [%2]", sqlca.sqlcode, SQLERRMESSAGE )
       CALL com.WebServiceEngine.SetRestError(406,userError)
     END TRY
     RETURN p_recList
END FUNCTION
```

## Related links

**Related concepts**  

[Set resource path with WSParam and WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL.")

[Set query, header, or cookie parameters](4730-set-query-header-or-cookie-parameters.md "Define the WSQuery, WSHeader, or WSCookie parameters in your function if the resource needs data passed as a query, cookie, or header.")
