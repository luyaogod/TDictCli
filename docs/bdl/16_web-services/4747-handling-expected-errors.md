---
title: "Example: handling expected application errors"
source: "fgl-topics/c_gws_restful_high_level_wsthrows_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling application level errors > Handling expected errors"
type: "concept"
---

# Example: handling expected application errors

> Error handling is supported with the WSThrows, WSError, and WSErrorHeader attributes.

In your REST function expected errors, such as application level errors, are listed in the
`WSThrows` attribute defined
in the `ATTRIBUTES` clause of the function.

## Example using WSError with WSThrows

In this sample REST function there is an example of error handling when a query for a customer
specified by id encounters an error. The function has an input parameter
id defined as type `INTEGER` with a `WSQuery` attribute.

The `WSThrows` attribute is set in the `ATTRIBUTES` clause of
the function to trap for two HTTP status codes, 404 (Not found) and 400 (bad request). The
function also references the "userError" variable (`@userError`) defined with
the [WSError](4805-wserror.md "Defines a description for an HTTP status code returned by the service.") attribute to return a
description of the error at runtime.

In the example, the module variable `userError` is a record but you can
define this as any suitable Genero BDL simple type.

If an error is encountered in the SQL query, this is handled in the
`TRY/CATCH` block. The `WSError` attributed variable is set
with a specific description of the HTTP status code, and the `SetRestError()` method
is called to return it.

You can define a custom header to return the error response of your web service function. For
more details, see [WSErrorHeader](4831-wserrorheader.md "Defines a custom HTTP header to include in an error response.").

```
IMPORT com

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError="User error")
  message STRING
END RECORD

PUBLIC FUNCTION GetCustomersNameById( id STRING ATTRIBUTES(WSQuery) )
  ATTRIBUTES(WSGet,
             WSPath = "/customers",
             WSThrows = "400:@userError")
  RETURNS STRING
    DEFINE s STRING
    TRY
      SELECT lname INTO s FROM customers WHERE @id  = id
      IF sqlca.sqlcode = NOTFOUND THEN
        LET userError.message = SFMT("Could not find customer id :%1",id)
        CALL com.WebServiceEngine.SetRestError(400,userError)
      END IF
    CATCH
      LET userError.message = SFMT(" - SQL error:%1 [%2]",
                                   sqlca.sqlcode, SQLERRMESSAGE)
      CALL com.WebServiceEngine.SetRestError(400,userError)
    END TRY
    RETURN s
END FUNCTION
```

## Related links

**Related concepts**  

[Example: handling an unexpected error](4748-handling-unexpected-errors.md "Example of a method you can use to code for unexpected errors.")

[WSErrorHeader](4831-wserrorheader.md "Defines a custom HTTP header to include in an error response.")
