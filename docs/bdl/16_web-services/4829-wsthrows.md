---
title: "WSThrows"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSThrows.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > Error handling attributes > WSThrows"
type: "concept"
---

# WSThrows

> Defines the list of error codes the REST function may return.

## Syntax

```
WSThrows = "{ code | code:description | code:@variable } [,...]"
```

`WSThrows` has a comma-separated list of status codes and/or error
descriptions where:

1. code is an HTTP status code in the range 400 to 599
2. description is text describing the error
3. @variable is a reference to a variable defined to handle error
   descriptions. The variable must be defined with a `WSError` attribute. It is
   returned in the response message body.

- There is no space before or after the colon (`:`) between
  code and description.
- If you use the code option on its own without a description, the
  default fault description defined in [RFC 2616](http://tools.ietf.org/html/rfc2616) is returned to the client.

`WSThrows` is an optional attribute.

## Usage

You use this attribute to manage application level errors, such as a database record not
found. If your function can potentially return one or more errors, you use
`WSThrows` to declare those so that they can be generated in the OpenAPI
documentation for that resource.

For instance, HTTP status code 404 means resource not found, but using the
`WSThrows` attribute you can respond to the client with `404:user id
not found` in your function. To return a description of the status code like this,
you must use `WSThrows` with the
"code:@variable" option of the attribute, and
reference a variable with the `WSError` attribute.

If you simply need to return an error code without a description, use
`WSThrows` with code, or with a description use
code:description.

For example, these are valid uses of `WSThrows`:
> **Important:**
>
> You must use valid HTTP status codes. Codes between 400 to 599 are valid, but use of
> codes between 200 to 399 representing successful HTTP status, result in `error-9117`.

**`WSThrows = "404, 402"`**

In this example, there are no descriptions of the error, just standard HTTP error
messages ( see [RFC 2616](http://tools.ietf.org/html/rfc2616#section-10.5) ) are returned and displayed on the
client side.

**`WSThrows= "404:not found error, 402:hello world"`**

In this example, the content after the colon (`:`) is the description
of the error displayed on the client side.

**`WSThrows= "404:@error1, 402:@error2"`**

Here error descriptions are provided in the `WSError` attributed
variables referenced in `@error1` and `@error2`.

If you need to
ensure an error description other than the standard protocol message is returned to the client, you
must use `WSThrows` with the
"`code:@variable`" option. Otherwise, the web
server may replace your description with the standard HTTP error message.

## Example function using WSThrows

In this sample REST function a user resource is updated. The function's
`thisUser` parameter of type `profileType` contains the
values to update. The `thisUser` data is passed in the message body in either
JSON or XML format.

The `WSThrows` attribute is set to handle errors that may be encountered. It
has options to respond to three HTTP errors:

- Error 404 when the resource is not found
- Error 412 when the server can not respond to the request
- Error 406 when the server throws an internal error and can not respond to the
  request

In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. Depending on the error thrown, error descriptions are
provided in the `WSError` attributed variables referenced in
`@userError` or `@fatalError`. A call to [SetRestError()](../15_library-reference/3800-com-webserviceengine-setresterror.md "Manages error handling for a REST high-level Web Service function.") returns the
message defined in `WSThrows` for the error.

```
IMPORT com  

TYPE profileType RECORD
     id INTEGER,
     name VARCHAR(100),
     email VARCHAR(255)
     # ...
   END RECORD

PUBLIC DEFINE fatalError INTEGER ATTRIBUTES(WSError="fatal error")

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION updateUser( thisUser profileType )
  ATTRIBUTES (WSPut, 
              WSPath ="/users/update",
              WSDescription ="Update user details",
              WSThrows ="404:@userError,
                         412:@fatalError,
                         406:should not happen" )
  RETURNS STRING
    DEFINE ret STRING
    TRY
      UPDATE users SET 
         name = thisUser.name, 
         email = thisUser.email 
         WHERE id = thisUser.id
         IF sqlca.sqlerrd[3] = 1 THEN  # sqlerrd[3] = processed rows
           LET ret = SFMT("Updated user with ID: %1",thisUser.id)
         ELSE
           LET userError.message = SFMT("No user with ID: %1",thisUser.id)
           CALL com.WebServiceEngine.SetRestError(404,userError)
         END IF
    CATCH
      CASE 
         WHEN sqlca.sqlcode < 0 
            LET fatalError = sqlca.sqlcode
            CALL com.WebServiceEngine.SetRestError(412,fatalError)
         OTHERWISE 
            CALL com.WebServiceEngine.SetRestError(406,NULL)
      END CASE
    END TRY
    RETURN ret
END FUNCTION
```

## Related links

**Related concepts**  

[Handling application level errors](4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service.")

[WSErrorHeader](4831-wserrorheader.md "Defines a custom HTTP header to include in an error response.")
