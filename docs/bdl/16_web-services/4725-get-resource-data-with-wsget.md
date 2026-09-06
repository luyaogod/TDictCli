---
title: "Example: Get operation with WSGet"
source: "fgl-topics/c_gws_restful_high_level_get_operation_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Define your resource operations > Get resource data with WSGet"
type: "concept"
---

# Example: Get operation with WSGet

> Example of methods you can use to get data from a resource with the WSGet attribute.

## Example using WSGet

In this sample REST function the number of users from the users table of a database is
returned.

```
PUBLIC FUNCTION getNumberUsers()
  ATTRIBUTES(WSGet, 
             WSPath = "/users/count",
             WSDescription = "Returns a count of users")
  RETURNS INTEGER
    DEFINE cnt INTEGER
    SELECT COUNT(*) INTO cnt FROM users
    RETURN cnt
END FUNCTION
```

## Example using WSGet to return all users

In this sample REST function all users are returned. The `WSGet` attribute
is set to request data from the service. An example of the resource URL is:

http://myhost/gas/ws/r/myGroup/myXcf/Account/users

```
TYPE profileType RECORD
     id INTEGER,
     name VARCHAR(100),
     email VARCHAR(255),
     category VARCHAR(10),
     status INTEGER,
     ccode VARCHAR(3)
     # ...
   END RECORD

PUBLIC FUNCTION getAllUsers()
  ATTRIBUTES(WSGet,
             WSPath = "/users",
             WSDescription = "Returns all user profiles"
             )
  RETURNS (DYNAMIC ARRAY OF profileType )

    DEFINE arr DYNAMIC ARRAY OF profileType
    DEFINE i INTEGER = 1
    
    # code to get users
     DECLARE usersCurs CURSOR FOR SELECT id, name FROM users ORDER BY name
     FOREACH usersCurs INTO arr[i].*
       LET i = i+1
     END FOREACH
     CALL arr.deleteElement(arr.getLength())
     # Remove the empty element implied by reference in FOREACH loop
  RETURN arr
END FUNCTION
```

## Related links

**Related concepts**  

[HTTP operation attributes (Verbs)](4814-http-operation-attributes-verbs.md "Attributes that map an HTTP operation or verb action in a function to a REST resource.")

[REST resource URI naming practice](4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.")
