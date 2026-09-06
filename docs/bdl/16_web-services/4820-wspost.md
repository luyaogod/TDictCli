---
title: "WSPost"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSPost.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > HTTP operation attributes (Verbs) > WSPost"
type: "concept"
---

# WSPost

> In order to create a new resource, you set the WSPost attribute.

## Syntax

```
WSPost
```

## Usage

You use this attribute to specify the action of the HTTP verb POST to create a new resource. You
set the `WSPost` attribute in the `ATTRIBUTES()` clause of
your function.

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
