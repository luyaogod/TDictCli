---
title: "Example: Update resource with WSPut"
source: "fgl-topics/c_gws_restful_high_level_update_resource_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Define your resource operations > Update a resource with WSPut"
type: "concept"
---

# Example: Update resource with WSPut

> Update a resource with the WSPut attribute.

For instance, when you need to update the full resource, you use [WSPut](4821-wsput.md "Update an existing resource with the WSPut attribute."). When you just want to update a single
field in a resource, you use [WSPatch](4819-wspatch.md "In order to partially update an existing resource, you define the WSPatch attribute.").

## Example updating a resource with WSPut

In this sample REST function a user resource is updated. In the function's
`id` parameter, the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") specifies the user to update and the
`thisUser` variable of type `profileType` contains the values to
update. The `thisUser` data is passed in the message body in either JSON or XML
format.

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
     ccode VARCHAR(3)
   END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION updateUsers(
    id INTEGER ATTRIBUTES(WSParam),
    thisUser profileType)
  ATTRIBUTES(WSPut,
             WSPath = "/users/{id}",
             WSDescription = "Update a user profile",
             WSThrows = "400:@userError")
  RETURNS STRING
    DEFINE ret STRING
    TRY
      UPDATE users
        SET name = thisUser.name,
            email = thisUser.email,
            category = thisUser.category,
            ccode = thisUser.ccode
        WHERE id  = thisUser.id
       IF sqlca.sqlerrd[3] = 1 THEN # sqlerrd[3] indicates processed rows
         LET ret = SFMT("Updated user with ID: %1",id)
       ELSE
         LET ret = SFMT("No user with ID: %1",id)
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

[HTTP operation attributes (Verbs)](4814-http-operation-attributes-verbs.md "Attributes that map an HTTP operation or verb action in a function to a REST resource.")
