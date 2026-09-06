---
title: "WSDescription"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSDescription.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes related to function parameters > WSDescription"
type: "concept"
---

# WSDescription

> Describes the REST function, parameters, return values, and members of user-defined types.

## Syntax

```
WSDescription=" description "
```

Where:

1. description provides a description of the parameter, function, or return
   value it is applied to.

`WSDescription` is an optional attribute.

## Usage

You use this attribute to provide useful description for parameters. You can set
`WSDescription` on:

- Any input and output parameter of a REST operation.
- The attribute clause of the function.
- Members of a user-defined [type](../08_language-basics/0753-type.md "Types define a synonym for a base or structured data type.").

Information you specify with this attribute, is generated in
the OpenAPI documentation file. You can control generation of descriptions in the stub file with the
`--comment` option of the [fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.") tool.

## Example 1: using WSDescription in the function and parameters

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

## Example 2: using WSDescription in a type definition

In this example, the `WSDescription` attribute describes members of the
user-defined type `profileType`.

```
TYPE profileType RECORD
     id INTEGER ATTRIBUTES (WSDescription="Internal Identifier"), 
     name VARCHAR(100) ATTRIBUTES (WSDescription="Lastname"),
     email VARCHAR(255),
     category VARCHAR(10) ATTRIBUTES (WSDescription="User Demographic"),
     status INTEGER,
     ccode VARCHAR(3) ATTRIBUTES (WSDescription="Country Code")
     # ...
   END RECORD
```

In the OpenAPI documentation, the `profileType` will be referenced in the
`#/components/schema` section. The GWS engine exposes the description property for
the elements defined with `WSDescription` in the `schema`.

The output shown in this example is from the Firefox™ browser, which formats JSON for readability. The appearance may vary depending on your browser.

![Image from the OpenAPI document showing the profileType JSON schema with description properties](../_images/rest_openapi_api_wsdescription_json_schema.png)

*Sample JSON schema with description properties*

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

[Set a request body](4731-set-a-request-body.md "Functions that create or update a resource need to set a request body for the incoming payload. You specify the request body in an input parameter.")

[Set a response body and header](4732-set-a-response-body-and-header.md "You specify a response body in a return parameter without an attribute. Other return values can be sent in headers, using the WSHeader attribute.")
