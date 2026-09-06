---
title: "WSGet"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSGet.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > HTTP operation attributes (Verbs) > WSGet"
type: "concept"
---

# WSGet

> In order to retrieve a resource, you set the WSGet attribute.

## Syntax

```
WSGet
```

## Usage

You use this attribute to specify the action of the HTTP verb GET to return data from a resource.
You set the `WSGet` attribute in the `ATTRIBUTES()` clause of the
function.

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

[Example: Get operation with WSGet](4725-get-resource-data-with-wsget.md "Example of methods you can use to get data from a resource with the WSGet attribute.")
