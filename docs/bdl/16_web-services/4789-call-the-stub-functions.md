---
title: "Call the stub functions"
source: "fgl-topics/c_gws_rest_high_level_client_using_functions.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Create the client application > Call the stub functions"
type: "concept"
---

# Call the stub functions

> Call the functions defined in the client stub to interact with the REST service from your application code.

The [client stub](4782-rest-stub-file-overview.md "The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.") provides the
functions and record types that your application uses to call the REST service. This topic shows how
to call a stub function from a client module.

## Calling a stub function in the client application

In this example, the function `myWScall` updates a user resource by calling a function provided in the client stub.

The variable `this` is defined as a stub-generated record type (`clientStub.addUserRequestBodyType`). Values are assigned to the fields before calling the stub function.

Two variables are defined for the return:

- `wsstatus` of type `INTEGER` for the status code
- `res` of type `STRING` for details returned by the service

The response is evaluated in a `CASE` statement. If `wsstatus` is not equal to `clientStub.C_SUCCESS`, the error details are displayed. For more details about handling errors, go to [Handle GWS REST server errors](4790-handle-rest-server-errors.md "Handle the status code returned by a REST service call and process expected and unexpected errors in the client application.").

```
IMPORT FGL clientStub

MAIN
CALL myWScall()
END MAIN

# Function performs call to the client stub function
FUNCTION myWScall()
  DEFINE this clientStub.addUserRequestBodyType
  DEFINE wsstatus INTEGER
  DEFINE res STRING

  LET this.user_name = "Mike Pantock"
  # ...
  CALL clientStub.addUser(this.*) RETURNING wsstatus, res
  CASE wsstatus
    WHEN clientStub.C_SUCCESS
      DISPLAY "Success adding new user"
    OTHERWISE
      DISPLAY "Unexpected error :", wsstatus, res
  END CASE

END FUNCTION
```

Find a complete client application at $FGLDIR/demo/WebServices/books.

## Related links

**Related concepts**  

[Write the MAIN program code block](4786-write-the-main-program-code-block.md "Create the MAIN program block that imports the client stub and calls its functions.")

[Handle GWS REST server errors](4790-handle-rest-server-errors.md "Handle the status code returned by a REST service call and process expected and unexpected errors in the client application.")

**Related tasks**  

[Main program code for access to secure service](4788-code-to-access-secure-service.md "Code to get an access token for a secure RESTful Web service from a Genero application also secured.")
