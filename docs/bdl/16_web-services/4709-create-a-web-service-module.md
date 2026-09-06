---
title: "Create a web service module"
source: "fgl-topics/t_gws_restful_high_level_quick_start_service_module.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 1: RESTful server application > Create a web service module"
type: "task"
---

# Create a web service module

> Resources are identified as a set of functions within the Web services module that provides the service.

In this task you create a Genero BDL module and code two functions that define access to [resources](4707-quick-start-1-restful-server-application.md).

In the RESTful high-level framework you must define functions as `PUBLIC` in order
to expose the resource.

A RESTful Web service function is a normal BDL function that uses [attributes](4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services."). These attributes
(`WS*`) are defined specifically for high-level REST.

In the function attributes clause there is an attribute that determines [the HTTP verb](4814-http-operation-attributes-verbs.md "Attributes that map an HTTP operation or verb action in a function to a REST resource.") (for example,
`WSGet`). The `WSPath` attribute determines the resource endpoint.
These two attributes must be included in all RESTful Web service functions.

Other attributes are set in the `ATTRIBUTES` property of input and output
parameters to define HTTP headers, etc.

1. Create your application module. For example, create a file named
   myservice.4gl.

   Start with importing the `com` Web services module, and declare that you want
   to use the `test1schema` schema:

   ```
   IMPORT com
   SCHEMA test1schema
   ```
2. Define types used by the REST functions.

   We define the following record type at the modular
   level:

   ```
   TYPE custinfoType RECORD LIKE custinfo.*
   ```
3. Create a function to get all customer records (copy and paste the code below)

   In the function's `ATTRIBUTES` clause the following attributes are set:
   - [WSGet](4816-wsget.md "In order to retrieve a resource, you set the WSGet attribute.") –
     specifies the GET HTTP verb.
   - [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") –
     specifies the resource endpoint.
   - [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") – provides a description of the
     resource which will be displayed in the generated OpenAPI specification file.

   In the function's `RETURNS` clause, we specify a `DYNAMIC ARRAY OF
   custinfoType` to return the complete set of customer records.

   ```
   PUBLIC FUNCTION getCustomerRecords()
       ATTRIBUTES(WSGet,
           WSPath = "/customers",
           WSDescription = "Fetches all customer records.")
       RETURNS (DYNAMIC ARRAY OF custinfoType)
       DEFINE custlist DYNAMIC ARRAY OF custinfoType
       DEFINE i INTEGER = 1
       DECLARE customerCurs CURSOR FOR SELECT * FROM custinfo ORDER BY cust_lname
       FOREACH customerCurs INTO custlist[i].*
           LET i = i + 1
       END FOREACH
       # Remove the empty element implied by reference in FOREACH loop
       CALL custlist.deleteElement(custlist.getLength())
       RETURN custlist
   END FUNCTION
   ```
4. Create a function to get a customer's details (copy and paste the code below)

   In the function's `ATTRIBUTES` clause the following attributes are set:
   - [WSGet](4816-wsget.md "In order to retrieve a resource, you set the WSGet attribute.") –
     specifies the GET HTTP verb.
   - [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") –
     specifies the resource endpoint. It includes a `{WSParam-identifier}` as
     a template path. The value will be specified in an input parameter with an [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") attribute.
   - [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") – provides a description of the
     resource which will be displayed in the generated OpenAPI specification file.
   - [WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") - specifies how we handle
     different outcomes of the request. We set these in the attribute with a standard HTTP status code
     and a message as follows:
     - `404:Not Found`, if the resource is not found.
     - `400:SQL error`, when an unexpected SQL error occurs.We trap errors in the `TRY/CATCH` block. We check the `SQLCA` record after the execution of the SQL
     query. The call to `SetRestError()` with `NULL` handles the return of the message defined
     in `WSThrows`.

   In the function's input parameter we set the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") as the part of the path to a specific
   customer resource. The customer id will be provided in the URI request, for
   example:

   ```
   GET http://localhost:8090/MyService/customers/2
   ```

   In the function's `RETURNS` clause we return the customer record in the
   `fCustomerRec` variable of type `custinfoType`.

   ```
   PUBLIC FUNCTION getCustomerById(
       resourceId INTEGER ATTRIBUTES(WSParam))
       ATTRIBUTES(WSGet,
           WSPath = "/customers/{resourceId}",
           WSDescription = "Fetches a single record from the Customer resource.",
           WSThrows = "404:Not Found,400:SQL Error")
       RETURNS(custinfoType)
       DEFINE fCustomerRec custinfoType
       TRY
           SELECT custinfo.*
               INTO fCustomerRec.*
               FROM custinfo
               WHERE custinfo.cust_num = resourceId
           IF sqlca.sqlcode == NOTFOUND THEN
               # 404 is defined by the WSThrows attribute
               CALL com.WebServiceEngine.SetRestError(404, NULL)
           END IF
       CATCH
           DISPLAY "SQL error: ", sqlca.sqlcode
           CALL com.WebServiceEngine.SetRestError(400, NULL)
       END TRY
       RETURN fCustomerRec
   END FUNCTION
   ```

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

**Related reference**  

[High-level RESTful Web service attributes](4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")
