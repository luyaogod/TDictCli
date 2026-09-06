---
title: "Create a new version of an operation"
source: "fgl-topics/t_gws_high_level_rest_version_operation.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Advanced REST API versioning > Operation-level versioning > Create a new version of an operation"
type: "task"
---

# Create a new version of an operation

> Specify the version in which the operation is available by setting the WSVersion attribute.

You version an operation if changes to an existing function, such as the request and/or response
data, would break existing use of the operation. You need to create a new version of the function to
apply the changes. The new version of the function will be offered with the latest release of your
REST web service. This will not affect existing users who will continue to use the previous
version.

If an operation does not change across versions, you can
leave it unversioned or set `WSVersion="default"`. The default version is used unless
the client requests a specific version.

Create a new version of an operation when you introduce changes that are not backward compatible
with the existing implementation, such as updates to the request or response data.

With versioning, you can reuse the same operation path and verb as many times as needed, provided
each version has a unique `WSVersion` value.

To create a new version of the operation:

1. Create a copy of the existing function, and paste it in the same file.
2. Rename the copied function. 

   For example operationId\_version\_name.
3. Add the `WSVersion` attribute to the `ATTRIBUTES` clause and set
   the version name.

   For example:

   ```
   IMPORT com

   PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
     message STRING
   END RECORD

   PUBLIC FUNCTION prices_v1(nb INTEGER ATTRIBUTES(WSParam))
     ATTRIBUTES (WSGet,WSPath = "/prices/{nb}",
                  WSDescription = "Returns the price of an item",
                  WSThrows = "404:item not found", 
                  WSVersion = "v1") 
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
4. Apply the necessary changes to the function.

To generate the OpenAPI documentation for this version, add
`&version=version_name` to the operation query string, for
example
http://myhost/gas/ws/r/myGroup/myXcf/myService?openapi.json&version=v1.

For information about generating a version of the OpenAPI document, see [Get the OpenAPI description for a specific version](4777-get-openapi-description-for-a-version.md "Retrieve the OpenAPI description for a specific version of a RESTful web service.").

## Related links

**Related concepts**  

[Service-level versioning](4762-service-level-versioning.md "Overview of settings that control versioning for an entire REST web service.")

[Operation-level versioning](4765-operation-level-versioning.md "Overview of settings that control versioning for individual REST operations.")

[Set the default version for an operation](4768-set-the-default-version-for-an-operation.md "Define which version of an operation is shown by default in the OpenAPI documentation.")

[Define multiple versions for an operation](4767-define-multiple-versions-for-an-operation.md "Specify which API versions an operation applies to using the WSVersion attribute.")

**Related tasks**  

[Get the OpenAPI description for a specific version](4777-get-openapi-description-for-a-version.md "Retrieve the OpenAPI description for a specific version of a RESTful web service.")

[Get available versions of a service](4778-get-available-versions-of-a-service.md "Retrieve the list of available versions of a REST service and identify the default version if defined.")
