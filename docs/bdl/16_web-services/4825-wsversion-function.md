---
title: "WSVersion (function)"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSVersion_function.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > WSVersion (function)"
type: "concept"
---

# WSVersion (function)

> Specifies the version or versions in which the REST function is available.

## Syntax

```
WSVersion = "{ version } [,...]"
```

Where `WSVersion` is a comma-separated list of versions and where:

1. version defines the version of the operation.

`WSVersion` is an optional attribute.

## Usage

You use this attribute to specify version at the function level. It activates versioning for the
Web service. When the GWS detects the `WSVersion`
attribute, it automatically switches to versioning mode. This affects how it generates the
OpenAPI documentation and ultimately how resources are accessed.  You must now add the query "&version=nameVersion" in
the URI to get the OpenAPI documentation. For example:
http://myhost/gas/ws/r/myservice?openapi.json&version=v2.

Set a version name on the `WSVersion` attribute in the
`ATTRIBUTES()` clause of the REST function. By doing this you specify the version of
the Web service in which the operation will be accessible to clients.

You can use this attribute to update a previous version of the same operation (with the same path
and verb) to a new version of the operation. Existing versions are not affected and can still be
used.

If an operation does not change across versions, you can
leave it unversioned or set `WSVersion="default"`. The default version is used unless
the client requests a specific version.

> **Important:**
>
> Setting a default version is optional. However, if no default version is defined and the client
> does not specify a version in the query string, the request returns a `404 (Not
> found)` and the server raises [error
> 41](../15_library-reference/3805-error-codes-of-com-webservicesengine.md) internally.

When an operation is accessible in one or more versions, you can
set these specifically within the attribute, for example,
`WSVersion="v1,v2,v3"`.

You can display the documentation for a specific version of the Web service, by adding the query
"&version=nameVersion" in the URI. For example:
http://myhost/gas/ws/r/fruits?openapi.json&version=v2

You can specify the version of operations of the Web service that are displayed in the OpenAPI
document and in the generated stub file by setting `WSVersion` in the service
information record of the module. For more information on specifying the default version of
operations of the Web service, see [WSVersion (module)](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.").

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

[How to name versions](4760-how-to-name-versions.md "There are no absolute rules to naming versions, but there are best practices.")
