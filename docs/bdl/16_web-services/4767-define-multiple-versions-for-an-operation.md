---
title: "Define multiple versions for an operation"
source: "fgl-topics/c_gws_high_level_rest_multiple_version_operation.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Advanced REST API versioning > Operation-level versioning > Define multiple versions for an operation"
type: "concept"
---

# Define multiple versions for an operation

> Specify which API versions an operation applies to using the WSVersion attribute.

You can maintain multiple versions of the same operation when changes are not backward
compatible. Each version is implemented as a separate function, with a unique
`WSVersion` value.

The following example shows an operation (`prices_v1v2`) that is available in two
versions: `v1` and `v2`.

## Multiple versions

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION prices_v1v2(nb INTEGER ATTRIBUTES(WSParam))
  ATTRIBUTES (WSGet,WSPath = "/prices/{nb}",
               WSDescription = "Returns the price of an item for API v1 and v2",
               WSThrows = "404:item not found", 
               WSVersion = "v1,v2") 
  RETURNS (INTEGER)
   DEFINE price DECIMAL(10,2)
    TRY
      SELECT unit_price INTO price FROM unit_pricing
        WHERE unit_id = nb
      IF sqlca.sqlcode = 100 THEN
        CALL com.WebServiceEngine.SetRestError(404,NULL)
      END IF
    CATCH
      LET userError.message = SFMT(" - SQL error:%1 [%2]",
                                   sqlca.sqlcode, SQLERRMESSAGE)
      CALL com.WebServiceEngine.SetRestError(505,userError)
    END TRY
    RETURN price
END FUNCTION
```

When generating the OpenAPI document, GWS includes this operation in any version where it is
declared. In this example, the operation appears in both `v1` and `v2`:

- http://myhost/gas/ws/r/myGroup/myXcf/myService?openapi.json&version=v1
- http://myhost/gas/ws/r/myGroup/myXcf/myService?openapi.json&version=v2

## Related links

**Related concepts**  

[Set the default service version for OpenAPI documentation](4763-set-the-default-service-version-for-openapi-documentation.md "Define which version of a service is shown by default in the OpenAPI documentation.")

**Related tasks**  

[Get the OpenAPI description for a specific version](4777-get-openapi-description-for-a-version.md "Retrieve the OpenAPI description for a specific version of a RESTful web service.")
