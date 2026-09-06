---
title: "Example: handling an unexpected error"
source: "fgl-topics/c_gws_restful_high_level_unexpected_error_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Handling application level errors > Handling unexpected errors"
type: "concept"
---

# Example: handling an unexpected error

> Example of a method you can use to code for unexpected errors.

> **Tip:**
>
> In general, the recommended option is to list all the possible errors in order
> to have them generated in the OpenAPI documentation for the Web service and trapped on the client
> side.

Not all errors can be anticipated. In this case, an unexpected error can be handled in your REST
function. Unexpected error are the ones **not** listed in the `WSThrows` attribute.

## Example: managing unexpected errors

In this sample REST function there is an example of error handling when an unexpected error is
encountered. The function has an input parameter id defined as
type `INTEGER` with a `WSParam`
attribute.

The `WSThrows` attribute is set in the `ATTRIBUTES`
clause of the function to trap for an HTTP status code "400" error. The function
also references the "userError" variable (`@userError`) defined with
the [WSError](4805-wserror.md "Defines a description for an HTTP status code returned by the service.") attribute to
return a description of the error at runtime.

If an error is encountered in the SQL query, this is handled in the
`TRY/CATCH` block. If the account is not found, the
`WSError` attributed variable is set with a specific description
of the HTTP 400 status code, and the `SetRestError()` method is called to return it.

If the error encountered is unexpected, the `SetRestError()` method
is called to return error code 505. Here the description is set to `NULL` to
allow the HTTP standard error description to be returned, but you could also return a
description in this case; optionally using the variable defined with the [WSError](4805-wserror.md "Defines a description for an HTTP status code returned by the service.") attribute.

```
IMPORT com

TYPE profileType RECORD
     id INTEGER,
     name VARCHAR(100),
     email VARCHAR(255)
     # ...
   END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError="User error")
  message STRING
END RECORD

PUBLIC FUNCTION queryAccountsById( id VARCHAR(10) ATTRIBUTES(WSParam) )
  ATTRIBUTES(WSGet,
             WSPath = "/accounts/{id}",
             WSThrows = "400:@userError")
  RETURNS profileType ATTRIBUTES(WSName = "body")
    DEFINE thisAccount profileType
    TRY
      SELECT * INTO thisAccount.* FROM accounts WHERE @id = id
      IF sqlca.sqlcode = NOTFOUND THEN
        LET userError.message = SFMT("Could not find account id :%1",id)
        CALL com.WebServiceEngine.SetRestError(400,userError)
      END IF
    CATCH
      CALL com.WebServiceEngine.SetRestError(505,NULL)
    END TRY
    RETURN thisAccount
END FUNCTION
```

## Related links

**Related concepts**  

[Example: handling expected application errors](4747-handling-expected-errors.md "Error handling is supported with the WSThrows, WSError, and WSErrorHeader attributes.")
