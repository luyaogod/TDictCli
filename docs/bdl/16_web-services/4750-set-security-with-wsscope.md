---
title: "Example: set security with WSScope"
source: "fgl-topics/c_gws_restful_high_level_wsscope_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling security > Set security with WSScope"
type: "concept"
---

# Example: set security with WSScope

> You can set security using the WSScope attribute either at the function level or at the service level.

## Example 1: Setting security with WSScope at function level

In this sample REST function there is an example of a function that requires authentication
to access it. To execute this REST operation requires the request contains an access token
with a scope that matches what is in `WSScope`.

The `WSScope` attribute is set in the `ATTRIBUTES` clause of
the function. In this example the scope is set to "profile" or "profile.me".

Access token errors are automatically handled by the GWS engine. You
do not need to do anything in your code. If the client request does not have the correct
access token, the service will return HTTP 403.

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

PUBLIC FUNCTION FetchMyUserProfile( id INTEGER ATTRIBUTES(WSQuery) )
  ATTRIBUTES(
    WSGet,
    WSPath = "/users/profile",
    WSDescription = "Returns a user profile, requires authentication",
    WSThrows = "404:user not found",
    WSScope = "profile, profile.me")
  RETURNS profileType ATTRIBUTES(WSName = "data",
                                 WSMedia = "application/json,application/xml")
    DEFINE p profileType
    TRY
      SELECT * INTO p.* FROM users
             WHERE @id = id
      IF sqlca.sqlcode = NOTFOUND THEN
        CALL com.WebServiceEngine.SetRestError(404,NULL)
      END IF
    CATCH
       CALL com.WebServiceEngine.SetRestError(505,NULL)
    END TRY
    RETURN p
END FUNCTION
```

## Example 2: Setting security at Web service level via WSScope

This example sets the scope in the service information record of the module. The attributes set
are `WSInfo` and `WSScope`. All REST functions in the module require
the scope "users.fourjs" in order to execute.

```
PUBLIC DEFINE serviceInfo 
  RECORD ATTRIBUTES(WSInfo,
                    WSScope="users.fourjs")
    title STRING,
    version STRING,
    contact STRING
  END RECORD
```

## Related links

**Related concepts**  

[WSInfo](4804-wsinfo.md "Defines general metadata about the REST service.")
