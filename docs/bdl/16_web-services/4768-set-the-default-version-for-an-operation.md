---
title: "Set the default version for an operation"
source: "fgl-topics/c_gws_high_level_rest_default_version_operation.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Advanced REST API versioning > Operation-level versioning > Set the default version for an operation"
type: "concept"
---

# Set the default version for an operation

> Define which version of an operation is shown by default in the OpenAPI documentation.

If an operation does not change across versions, you can
leave it unversioned or set `WSVersion="default"`. The default version is used unless
the client requests a specific version.

This setting applies at the operation level.

## Example setting WSVersion in functions

In this example, the function `prices_default` is marked as the default version of
the `prices` operation. The web service defines two versions of this operation:
`v2` and `v3`.

When generating the OpenAPI document for version `v3` (for example:

`http://myhost/gas/ws/r/myGroup/myXcf/myService?openapi.json&version=v3`), GWS
selects the `prices_default` operation because it is declared as the default
version.

For version `v2`, GWS instead selects the `prices_v2` operation
because that version of the function is explicitly available.

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

TYPE priceType RECORD
  item VARCHAR(150),
  unit_measure VARCHAR(150),
  price DECIMAL(10,2)
END RECORD

PUBLIC FUNCTION prices_default(nb INTEGER ATTRIBUTE(WSParam))
  ATTRIBUTES (WSGet,WSPath = "/prices/{nb}",
               WSDescription = "Returns the price of an item",
               WSThrows = "404:item not found",
               WSVersion = "default")
  RETURNS (INTEGER)
   DEFINE price DECIMAL(10,2)
    TRY
      SELECT unit_price INTO price FROM unit_pricing
        WHERE @unit_id = nb
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

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION quantity_v3(id INTEGER ATTRIBUTE(WSParam) )
 ATTRIBUTES (WSGet,WSPath = "/quantity/{id}",
              WSDescription = "Returns an item quantity for API v3",
              WSThrows = "404:item not found",
              WSVersion = "v3")
 RETURNS (INTEGER)
   DEFINE quantity INTEGER
   
     TRY
      SELECT qty INTO quantity FROM inventory
        WHERE itemid = id
      IF sqlca.sqlcode == NOTFOUND THEN
        CALL com.WebServiceEngine.SetRestError(404,NULL)
      END IF
    CATCH
      LET userError.message = SFMT(" - SQL error:%1 [%2]",
                                   sqlca.sqlcode, SQLERRMESSAGE)
      CALL com.WebServiceEngine.SetRestError(400,userError)
    END TRY
  RETURN quantity
 END FUNCTION
```

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

TYPE priceType RECORD
  item VARCHAR(150),
  unit_measure VARCHAR(150),
  price DECIMAL(10,2)
END RECORD

PUBLIC FUNCTION prices_v2(nb INTEGER ATTRIBUTE(WSParam))
  ATTRIBUTES (WSGet,WSPath = "/prices/{nb}",
               WSDescription = "Returns price details for an item for API v2",
               WSThrows = "404:item not found",
               WSVersion = "v2")
  RETURNS (priceType)
    DEFINE price priceType
    TRY
      SELECT unit_title, measure_unit, unit_price INTO price.*
        FROM unit_pricing
        INNER JOIN measure_type
        ON unit_pricing.measure_typeid = measure_type.measure_id
        WHERE unit_pricing.unit_id = nb

      IF sqlca.sqlcode == NOTFOUND THEN
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

## Related links

**Related concepts**  

[WSVersion (function)](4825-wsversion-function.md "Specifies the version or versions in which the REST function is available.")

[Set the default service version for OpenAPI documentation](4763-set-the-default-service-version-for-openapi-documentation.md "Define which version of a service is shown by default in the OpenAPI documentation.")

[Define multiple versions for an operation](4767-define-multiple-versions-for-an-operation.md "Specify which API versions an operation applies to using the WSVersion attribute.")

**Related tasks**  

[Get the OpenAPI description for a specific version](4777-get-openapi-description-for-a-version.md "Retrieve the OpenAPI description for a specific version of a RESTful web service.")
