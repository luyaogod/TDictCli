---
title: "Set query, header, or cookie parameters"
source: "fgl-topics/c_gws_restful_high_level_set_query_cookie_header_parameter_examples.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Set query, header, or cookie parameters"
type: "concept"
---

# Set query, header, or cookie parameters

> Define the WSQuery, WSHeader, or WSCookie parameters in your function if the resource needs data passed as a query, cookie, or header.

The GWS REST engine provides support for the OpenAPI specification requirements for
these standard parameters:

- **Query parameters**. Set query parameters after a question mark (`?`) at the
  end of the resource URL. Ampersands (`&`) separate different
  `name=value` pairs. Query parameters can be required or optional.
- **Header parameters**, such as `X-MyHeader:Value`. These are custom headers
  sent with an HTTP request or response.
- **Cookie parameters**. The `Cookie` header passes these parameters, for
  example `Cookie:ctoken=BUSe35dohU4O1MZxDCU`.

## Parameter attributes

Set the attributes [WSQuery](4836-wsquery.md "Maps a parameter to a query string value in the request URL."), [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value."), and [WSCookie](4835-wscookie.md "Maps a parameter to a cookie value in the HTTP request.") in the parameters of
your function if there is requirement for them in your Web service. You can set
these parameters as optional with the [WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") attribute.

## Example use of parameter attributes

In this sample REST function a set of users is
returned based on values provided in input parameters. [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") attributes provide
descriptions of the parameters and the function. The function has three input parameters:

- The `stat` parameter is set with the `WSQuery` and the
  [WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") attributes. In a call to the
  function, the query string of the URI must be set by the name specified by `WSName`,
  such as http://myhost/gas/ws/r/myGroup/myXcf/Accounts/users?status=1
- The `ccode` parameter is set with the `WSCookie`,
  `WSName`, and the `WSOptional` attributes. If the client application
  calling the function provides a value for the cookie in the request, the name specified by
  `WSName` must be set in the query string of the URI, such as
  http://myhost/gas/ws/r/myGroup/myXcf/Accounts/users?status=1&country=FRA
- The `cat` parameter is set with the `WSHeader`,
  `WSName`, and the `WSOptional` attributes. The client
  application calling the function has the option of providing a value for a header named
  "category" in the request.

In the SQL query to the database, the `COALESCE` function is used to
produce a query that supports the parameters that are optional and may have no values.

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
     status INTEGER,
     ccode VARCHAR(3)
     # ...
   END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION getFilteredUsers(
    stat STRING ATTRIBUTES(WSQuery, WSName = "status",WSDescription = "Status is true(1) or false(0)" ),
    ccode STRING ATTRIBUTES(WSCookie, WSOptional, WSName = "country",WSDescription = "Country code"),
    cat STRING ATTRIBUTES(WSHeader, WSOptional, WSName = "category",WSDescription = "Product category")
    )
  ATTRIBUTES (WSGet,
              WSPath = "/filteredUsers",
              WSDescription = "Gets users based on query, cookie, and header values.",
              WSThrows = "400:Invalid,406:@userError" )
   RETURNS ( DYNAMIC ARRAY ATTRIBUTES(WSName = "Users_status",
             WSMedia = "application/json") OF profileType)
      
     DEFINE arr DYNAMIC ARRAY OF profileType
     DEFINE i INTEGER = 1
     TRY
       # The SQL COALESCE function produces a query that supports null parameters
       DECLARE c2 CURSOR FOR SELECT * FROM users
                               WHERE users.status = stat
                               AND users.country = COALESCE(ccode,users.country)
                               AND users.category = COALESCE(cat,users.category)
                               ORDER BY users.name ASC
       FOREACH c2 INTO arr[i].*
         LET i = i+1
       END FOREACH
       CALL arr.deleteElement(arr.getLength())
       # Remove the empty element implied by reference in FOREACH loop
     CATCH
       LET userError.message = SFMT("Error in SQL execution: %1 [%2]", 
                                     sqlca.sqlcode, SQLERRMESSAGE )
       CALL com.WebServiceEngine.SetRestError(406,userError)
     END TRY
     RETURN arr
END FUNCTION
```

## Related links

**Related concepts**  

[Designing REST Web services](4704-designing-rest-web-services.md "Identifying the resources to expose to a RESTful Web service client is an essential step in designing a RESTful Web service server.")

[Retrieve HTTP headers](4733-retrieve-http-headers.md "You can retrieve HTTP headers in your REST operation. There are two methods for doing this.")
