---
title: "WSPatch"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSPatch.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > HTTP operation attributes (Verbs) > WSPatch"
type: "concept"
---

# WSPatch

> In order to partially update an existing resource, you define the WSPatch attribute.

## Syntax

```
WSPatch
```

## Usage

You use this attribute to specify the action of the HTTP verb PATCH to partially update an
existing resource. For instance, when you need to update the full resource, you use [WSPut](4821-wsput.md "Update an existing resource with the WSPut attribute."). When you just want to update a single
field in a resource, you use [WSPatch](4819-wspatch.md "In order to partially update an existing resource, you define the WSPatch attribute."). The
PATCH method is described in [rfc5789](https://datatracker.ietf.org/doc/html/rfc5789) (external link).

You set the `WSPatch` attribute in the `ATTRIBUTES()` clause of the
function.

An output message body is not allowed in the response, so returns
must be specified as headers with the [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") attribute.

## Example using WSPatch to push partial resource change

In this sample REST function a user email address is updated. In the function's
`id` parameter the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") specifies the user to update,
and the `newEmail` parameter contains the value to update. The
`newEmail` data is passed in the message body in either JSON or XML format.

A string is returned as a header. It is specified with a [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") attribute.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors. In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. The `SQLERRMESSAGE` is set to the
`message` field of the `userError` variable, and a call to
`SetRestError()` returns the message defined in `WSThrows`
for the error.

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION UpdateUserProfile(
   id STRING ATTRIBUTE(WSParam), newEmail STRING)
   ATTRIBUTES(WSPatch,
              WSPath = "/users/{id}",
              WSDescription = "Update user email address",
              WSThrows = "400:@UserError,404:no user found")
   RETURNS STRING ATTRIBUTE (WSHeader)
   DEFINE ret STRING
    TRY
      UPDATE users SET email = newEmail
        WHERE @id  = id
       IF sqlca.sqlerrd[3] = 1 THEN # sqlerrd[3] indicates processed rows
         LET ret = SFMT("Updated user with ID: %1",id)
       ELSE
         CALL com.WebServiceEngine.SetRestError(404,NULL)
       END IF
    CATCH
       LET ret=SFMT("Error updating user with ID: %1",id)     
       LET userError.message = SFMT("SQL error:%1 [%2]",
                                    sqlca.sqlcode, SQLERRMESSAGE)
       CALL com.WebServiceEngine.SetRestError(400,userError)
    END TRY
    RETURN ret

 END FUNCTION
```

## Related links

**Related concepts**  

[HTTP verbs and attributes](4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.")
