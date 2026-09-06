---
title: "Versioning with URI"
source: "fgl-topics/c_gws_restful_high_level_versioning_uri.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Hand-coding REST versioning > Versioning with URI"
type: "concept"
---

# Versioning with URI

> You can version a resource using the WSPath attribute. The version is included in the resource URI.

In effect it means that you have two versions of the API and are able to generate a different
openapi.json description per version based on the path of your REST API.

For example, before versioning a resource was at:

http://myhost/gas/ws/r/myGroup/myXcf/resource

After versioning the resource is now at:

http://myhost/gas/ws/r/myGroup/myXcf/myVersion/resource

Even though this in principle breaks the constraint that the URI should refer to a unique resource, it is the
preferred way.

## Example setting version in the path

This example assumes changes to the profile record type (`profileType2`) that are not compatible with
the previous version of the API. The `WSPath` to
the new function now contains a version number (`v2`) that provides the path to the resource.

```
IMPORT com

TYPE profileType2 RECORD
     id INTEGER,
     fname VARCHAR(100),
     lname VARCHAR(100),
     email VARCHAR(255)
     # ...
   END RECORD

PUBLIC FUNCTION FetchMyUser2Profile(id INTEGER ATTRIBUTES(WSParam) )
  ATTRIBUTES(
             WSGet,
             WSPath = "/users/v2/{id}",
             WSDescription = "Returns a user profile for API v2 ",
             WSThrows = "404:user not found"
             )
  RETURNS profileType2 ATTRIBUTES(WSName = "data",
                                 WSMedia = "application/json,application/xml")
    DEFINE p profileType2
    TRY
      SELECT * INTO p.* FROM users2
             WHERE @id = id
      IF sqlca.sqlcode == NOTFOUND THEN
        CALL com.WebServiceEngine.SetRestError(404,NULL)
      END IF
    CATCH
      LET userError.message = SFMT(" - SQL error:%1 [%2]",
                                   sqlca.sqlcode, SQLERRMESSAGE)
      CALL com.WebServiceEngine.SetRestError(400,userError)
    END TRY
    RETURN p
END FUNCTION
```

An example of the resource URL, using the new version to retrieve the data for user 1, is:

http://myhost/gas/ws/r/myGroup/myXcf/Accounts/users/v2/1

## Related links

**Related concepts**  

[REST resource URI naming practice](4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.")
