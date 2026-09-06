---
title: "WSDelete"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSDelete.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > HTTP operation attributes (Verbs) > WSDelete"
type: "concept"
---

# WSDelete

> In order to remove an existing resource, you set the WSDelete attribute.

## Syntax

```
WSDelete
```

## Usage

You use this attribute to specify the action of the HTTP verb DELETE to delete a resource. You
set the `WSDelete` attribute in the `ATTRIBUTES()` clause of the
function.

## Example deleting a resource with WSDelete

In this sample REST function a user resource is deleted. In the function's
`id` parameter the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") specifies the user to
delete.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors. In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. The `SQLERRMESSAGE` is set to the
`message` field of the `userError` variable, and a call to
`SetRestError()` returns the message defined in `WSThrows`
for the error.

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION deleteUsers( id INTEGER ATTRIBUTES(WSParam) )
  ATTRIBUTES(WSDelete,
             WSPath = "/users/{id}",
             WSDescription = "Delete a user profile",
             WSThrows = "400:@userError")
  RETURNS STRING
    DEFINE ret STRING
    TRY
      DELETE FROM users WHERE @id = id
      IF sqlca.sqlerrd[3] = 1 THEN # sqlerrd[3] = processed rows
         LET ret = SFMT("Deleted user with ID: %1",id)
      ELSE
         LET ret = SFMT("No user with ID: %1",id)
      END IF
    CATCH
      LET ret = SFMT("Error deleting user with ID: %1",id)     
      LET userError.message = SFMT("SQL error:%1 [%2]",
                                    sqlca.sqlcode, SQLERRMESSAGE)
      CALL com.WebServiceEngine.SetRestError(400,userError)
    END TRY
    RETURN ret
END FUNCTION
```
