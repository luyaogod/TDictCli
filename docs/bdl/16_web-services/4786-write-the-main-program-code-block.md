---
title: "Write the MAIN program code block"
source: "fgl-topics/c_gws_rest_high_level_client_main.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Create the client application > Write the MAIN program code block"
type: "concept"
---

# Write the MAIN program code block

> Create the MAIN program block that imports the client stub and calls its functions.

After generating the [client stub](4782-rest-stub-file-overview.md "The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.")
for the REST service, write the `MAIN` program block that performs the calls to the
service functions.

In your `MAIN` module, you:

1. Import the client stub module.
2. If the service requires authentication, obtain the access token before calling any functions. See [Main program code for access to secure service](4788-code-to-access-secure-service.md "Code to get an access token for a secure RESTful Web service from a Genero application also secured.")
3. Call the functions provided by the stub.

```
IMPORT FGL clientStub
MAIN
  DEFINE wsstatus INTEGER
  DEFINE rets STRING
  DEFINE opt1 INTEGER
  LET opt1=22
  CALL clientStub.getUsersName(opt1) RETURNING wsstatus, rets
  IF wsstatus = clientStub.C_SUCCESS THEN
    DISPLAY "User name is :",rets
  ELSE
    DISPLAY "Error in getting user details"
  END IF
END MAIN
```

This example is a basic version of the code. For a complete client application, see $FGLDIR/demo/WebServices/books.

## Related links

1. [Access to secure web service](4787-access-to-secure-web-service.md)

   To access a secure RESTful Web service, the client application must have a valid access token.
2. [Main program code for access to secure service](4788-code-to-access-secure-service.md)

   Code to get an access token for a secure RESTful Web service from a Genero application also secured.

**Related concepts**  

[Call the stub functions](4789-call-the-stub-functions.md "Call the functions defined in the client stub to interact with the REST service from your application code.")

**Related tasks**  

[Main program code for access to secure service](4788-code-to-access-secure-service.md "Code to get an access token for a secure RESTful Web service from a Genero application also secured.")
