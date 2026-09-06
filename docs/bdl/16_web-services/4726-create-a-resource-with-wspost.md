---
title: "Example: create resource with WSPost"
source: "fgl-topics/c_gws_restful_high_level_create_resource_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Define your resource operations > Create a resource with WSPost"
type: "concept"
---

# Example: create resource with WSPost

> Create a new resource with the WSPost attribute.

## Example creating resource with WSPost

In this sample REST function a new user resource is created. The function's input parameter
thisUser variable of type `profileType` provides the new
user's details. The `thisUser` data is passed in the message body in either
JSON or XML format.

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
     status INTEGER
     # ...
   END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION createUser( thisUser profileType )
  ATTRIBUTES(WSPost,
             WSPath = "/users",
             WSDescription = "Create a user profile",
             WSThrows = "400:@userError")
  RETURNS STRING
    DEFINE ret STRING
    TRY
      INSERT INTO users VALUES (thisUser.*)
      LET ret = SFMT("Created user: %1",thisUser.name)
    CATCH
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

[Handling application level errors](4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service.")
