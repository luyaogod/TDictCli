---
title: "Handle GWS REST server errors"
source: "fgl-topics/c_gws_high_level_rest_error_handling.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Create the client application > Handle REST server errors"
type: "concept"
---

# Handle GWS REST server errors

> Handle the status code returned by a REST service call and process expected and unexpected errors in the client application.

The fglrestful tool generates types and constants on the client stub to help
you handle expected and unexpected errors returned by a REST service. The exact values depend on how
error handling is defined in the service. For more details, go to [Handling application level errors](4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service.").

Your client application must check the returned status code to determine how to handle the
response. In general, you:

- Check for a successful operation response.
- Check for known or expected errors, such as a bad request, resource not found, and so on.
- Check for unknown or unexpected errors, such as an internal server error at runtime.

See also: Trapping server errors on the client side.

## Expected errors

The fglrestful tool generates constants for expected errors. Status code 0
means success; values above 1000 indicate expected errors defined by the service.

```
# Error codes
PUBLIC CONSTANT C_SUCCESS = 0
PUBLIC CONSTANT C_USERERROR = 1001
PUBLIC CONSTANT C_NOT_FOUND = 1002
```

When an expected error occurs, the service populates the `userError` record.
Reference this record to get details.

```
# generated userErrorErrorType
PUBLIC TYPE userErrorErrorType RECORD
    message STRING
END RECORD

# Forced error
PUBLIC DEFINE userError userErrorErrorType
```

## Unexpected errors

Unexpected errors use the `wsError` record type. These can occur at runtime or
when an HTTP status code is returned that is not defined in the OpenAPI specification.

- If `wsError.code` < 0, a GWS runtime error occurred (for example, service not
  running).
- If `wsError.code` is an HTTP status (200–599), it indicates an error not defined
  in the service specification.

The `wsError.description` field provides additional
details.

```
PUBLIC DEFINE wsError RECORD
  code INTEGER,
  description STRING
END RECORD
```

## Trapping server errors on the client side

Define a variable to hold the returned status code. Evaluate it in a `CASE`
statement to handle success, expected errors, or unexpected errors.

- If `wsstatus = 0` (`clientStub.C_SUCCESS`), the call
  succeeded.
- If `wsstatus > 0`, an expected error occurred. Check `userError`
  for details.
- If `wsstatus < 0`, an unexpected error occurred. Check
  `wsError`.

Always include a `CASE OTHERWISE` branch to handle unexpected connection or
transport failures.

```
IMPORT FGL clientStub

MAIN
   CALL myWSRESTcall()
END MAIN

# Function performs call to the client stub function 
FUNCTION myWSRESTcall()
    DEFINE rets INTEGER
    DEFINE wsstatus INTEGER

    DEFINE p_body clientStub.updateUsersRequestBodyType
    LET p_body.users_id = "4" 
    LET p_body.users_name = "new name to be decided"
    # ...

     CALL clientStub.updateUsers(p_body.*) RETURNING wsstatus, rets
     CASE wsstatus
       WHEN clientStub.C_SUCCESS
         DISPLAY "Updated the user:"
       WHEN clientStub.C_NOT_FOUND
         DISPLAY "User not found"
         DISPLAY "error reason :", clientStub.userError.message
       WHEN clientStub.C_USERERROR
         #... Expected error
         DISPLAY "status code : ", wsstatus
         DISPLAY "Error reason :", clientStub.userError.message
       OTHERWISE
         #... Unexpected GWS error
         DISPLAY "Internal error"
         DISPLAY "status code : ", wsstatus
         DISPLAY "error code   :", clientStub.wsError.code
         DISPLAY "error reason :", clientStub.wsError.description
     END CASE
END FUNCTION
```

## Related links

**Related concepts**  

[Handling application level errors](4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service.")
