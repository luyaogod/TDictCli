---
title: "WSCookie"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSCookie.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes set on input parameters only > WSCookie"
type: "concept"
---

# WSCookie

> Maps a parameter to a cookie value in the HTTP request.

## Syntax

```
WSCookie
```

`WSCookie` is an optional attribute.

## Usage

You set this attribute to indicate that a parameter value will transport a cookie.
The client application needs to provide the appropriate header and value when making
a call to the function.

`WSCookie` is defined in an `ATTRIBUTES()` clause of an
input parameter of the function.

## Example WSCookie and WSOptional

In this sample REST function a set of users is returned based on a value provided in a
cookie sent as a header. The function's `ccode` input parameter is set with
the `WSCookie`, `WSName`, and the `WSOptional`
attributes.

If the client application calling the function provides a value for the cookie in the
request, the name specified by `WSName` must be set in the query string of the URI,
such as
http://myhost/gas/ws/r/myGroup/myXcf/Accounts/usersbycountry?country=FRA

In the SQL query of the database, the `COALESCE` function is used to produce
a query that supports the optional parameter. If the cookie is not provided, all users are
returned.

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
     country VARCHAR(3)
     # ...
   END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION getUsersByCountry(
     ccode STRING ATTRIBUTE(WSCookie, WSOptional, WSName = "country",
                            WSDescription = "Country code" )
                            )
  ATTRIBUTES (WSGet,
              WSPath = "/usersbycountry",
              WSDescription = "Gets users with the optional cookie value applied.",
              WSThrows = "400:Invalid,406:@userError" )
   RETURNS ( DYNAMIC ARRAY ATTRIBUTE(WSName = "Users",
             WSMedia = "application/json") OF profileType)
      
     DEFINE arr DYNAMIC ARRAY OF profileType
     DEFINE i INTEGER = 1
     TRY
       # code to get users
       DECLARE c4 CURSOR FOR SELECT * FROM users
                               WHERE users.country = COALESCE(ccode,users.country)
                               ORDER BY users.name ASC         
       # COALESCE function is used to produce a query that supports the optional parameter
       FOREACH c4 INTO arr[i].*
         LET i = i+1
       END FOREACH
       CALL arr.deleteElement(arr.getLength())
       # Remove the empty element implied by reference in FOREACH loop
     CATCH
       LET userError.message = SFMT("Error in SQL execution: %1 [%2]", sqlca.sqlcode, SQLERRMESSAGE )
       CALL com.WebServiceEngine.SetRestError(406,userError)
     END TRY
     RETURN arr
END FUNCTION
```
