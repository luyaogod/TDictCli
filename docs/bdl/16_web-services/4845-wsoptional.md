---
title: "WSOptional"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSOptional.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes related to function parameters > WSOptional"
type: "concept"
---

# WSOptional

> Marks a parameter or return value as optional in the REST message.

## Syntax

```
WSOptional
```

`WSOptional` is an optional attribute.

## Usage

You can set `WSOptional` on the `ATTRIBUTES()`clause of a parameter
when a function does not require the parameter to be present when a client calls the
service. Setting parameters as optional is typically used when defining REST
functions in the following situations:

- If a query string (`WSQuery`)
  is not always required, set the input parameter with
  `WSOptional`. For an example, see [Example use of parameter attributes](4730-set-query-header-or-cookie-parameters.md).
- If a custom header (`WSHeader`), input or output, is not always required, set the parameter
  with `WSOptional`. For an example, see [Example use of parameter attributes](4730-set-query-header-or-cookie-parameters.md).
- If a cookie (`WSCookie`) is not always required, set the parameter with
  `WSOptional`. See Example WSCookie and WSOptional.
- If an input body parameter is not always required, set the parameter with
  `WSOptional`. See Example optional input body.

  Normally, a client
  request for a service with `WSPost`, `WSPut`, or `WSPatch` requires an **input** message body, and failing to
  specify a message body in a request will raise [error-9106](../15_library-reference/4483-genero-bdl-errors.md) or [error-9128](../15_library-reference/4483-genero-bdl-errors.md). With the body parameter defined as optional, an
  input body request is not required when a client calls the service. For
  more information, see [Set a request body](4731-set-a-request-body.md "Functions that create or update a resource need to set a request body for the incoming payload. You specify the request body in an input parameter.").
  > **Important:**
  >
  > If you have [multipart](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.") input bodies and any one input body is optional, then all input
  > parameters must set the `WSOptional` attribute; otherwise
  > fglcomp will raise [error-9138](../15_library-reference/4483-genero-bdl-errors.md).

## Example WSCookie and WSOptional

In this sample REST function a set of users is returned based on a value provided in a
cookie sent as a header. The function's `ccode` input parameter is set with
the `WSCookie`, `WSName`, and the `WSOptional`
attributes.

If the client application calling the function provides a value for the cookie in the
request, the name specified by `WSName` must be set in the query string of the URI,
such as
http://myhost/gas/ws/r/myGroup/myXcf/Accounts/usersbycountry?country=FRA

In the SQL query of the database, the `COALESCE` function is used to produce
a query that supports the optional parameter. If the cookie is not provided, all users are
returned.

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
     email VARCHAR(255),
     category VARCHAR(10),
     status INTEGER,
     country VARCHAR(3)
     # ...
   END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION getUsersByCountry(
     ccode STRING ATTRIBUTE(WSCookie, WSOptional, WSName = "country",
                            WSDescription = "Country code" )
                            )
  ATTRIBUTES (WSGet,
              WSPath = "/usersbycountry",
              WSDescription = "Gets users with the optional cookie value applied.",
              WSThrows = "400:Invalid,406:@userError" )
   RETURNS ( DYNAMIC ARRAY ATTRIBUTE(WSName = "Users",
             WSMedia = "application/json") OF profileType)
      
     DEFINE arr DYNAMIC ARRAY OF profileType
     DEFINE i INTEGER = 1
     TRY
       # code to get users
       DECLARE c4 CURSOR FOR SELECT * FROM users
                               WHERE users.country = COALESCE(ccode,users.country)
                               ORDER BY users.name ASC         
       # COALESCE function is used to produce a query that supports the optional parameter
       FOREACH c4 INTO arr[i].*
         LET i = i+1
       END FOREACH
       CALL arr.deleteElement(arr.getLength())
       # Remove the empty element implied by reference in FOREACH loop
     CATCH
       LET userError.message = SFMT("Error in SQL execution: %1 [%2]", sqlca.sqlcode, SQLERRMESSAGE )
       CALL com.WebServiceEngine.SetRestError(406,userError)
     END TRY
     RETURN arr
END FUNCTION
```

## Example: Optional input body

In this sample REST function the value returned is dependent on whether data has been
received in an input body. The function has an input parameter `rec` that is
set with the `WSOptional` attribute. Therefore, data may or may not be
received in the request body.

A user-defined variable `Context` set with the [WSContext](4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.") attribute is defined in the
module. This is used to detect if the client has sent a body at runtime. If the request does
not have an input body, the GWS sets the "`NoRequestBody`" entry in the
`Context` variable to true.

In the function the `Context` is checked for the status of the
"`NoRequestBody`" entry, and the value the function returns is based on
this.

```
PRIVATE DEFINE Context DICTIONARY ATTRIBUTES(WSContext) OF STRING

PUBLIC
FUNCTION testRecord(rec RECORD ATTRIBUTES(WSOptional) a INTEGER, b INTEGER END RECORD)
  ATTRIBUTES (WSPut, 
             WSPath = "/record",
             WSDescription = "Check for request body with WSContext variable")
  RETURNS (INTEGER)
  IF Context.contains("NoRequestBody") THEN
    DISPLAY Context["NoRequestBody"]
    RETURN 0
  ELSE
     RETURN rec.a + rec.b
  END IF
END FUNCTION
```

## Related links

**Related concepts**  

[HTTP verbs and attributes](4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.")
