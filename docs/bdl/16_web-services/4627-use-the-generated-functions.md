---
title: "Use the generated functions"
source: "fgl-topics/c_gws_handlers_client_009.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > Use the generated functions"
type: "concept"
---

# Use the generated functions

> Learn how to use the functions and records in the stub file to write client applications that use the Web service.

If you have [generated stub
files](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.") for compatibility with your legacy GWS client application using
`GLOBALS`, see [Using functions generated with globals](4628-using-functions-generated-with-globals.md "Learn how to use the functions and global records in the stub file to write client applications that use the Web service.")

The information about the Add function from the stub file, for example, allows you to write code in
your own .4gl module that uses this operation as part of your Client application.

It is recommended to call functions in the stub via their qualified names. For example, use
`CALL ws_myStub.Add(1,2)` instead of `CALL
Add(1,2)`.

## Using parameters and return values

You call the `Add` function in your client application, defining variables for the
parameters and return values. The input variables for our example are simple integers.

```
FUNCTION myWScall()  
  DEFINE op1 INTEGER
  DEFINE op2 INTEGER
  DEFINE result INTEGER
  DEFINE wsstatus  INTEGER
  ...
  LET op1 = 6 
  LET op2 = 8
  CALL ws_myStub.Add(op1, op2) 
    RETURNING wsstatus, result ...
  DISPLAY result
```

## Using public records

An alternative option is to call the `Add_g` function instead, using the public
records `Add` and `AddResponse` directly. If the input variables are
complex structures like records or arrays, you are required to use this
function.

```
FUNCTION myWScall()
  DEFINE wsstatus INTEGER
  # ...
  LET ws_myStub.Add.a = 6
  LET ws_myStub.Add.b = 8
  LET wsstatus = ws_myStub.Add_g()
  # ...
  DISPLAY ws_myStub.AddResponse.r
```

In this case, the status is returned by the function, which has also put the result in the
`AddResponse` public record.

See [Tutorial: Writing a
Client Application](4595-steps-to-write-a-gws-client.md "A SOAP Web service contains functions that you need to call. Create a Genero BDL application that requests the functions of the service.") for more information. The demo/WebServices
subdirectory of your Genero installation directory contains complete examples of Client
Applications.

## Using asynchronous calls

If you don't want your application to be blocked when waiting for the response to a request, it
is recommended that you first call `AddRequest_g`; this will send the request using
the public `Add` record to the server. It returns a status of 0 (zero) if everything
goes well, -1 in case of error, or -2 if you tried to resend a new request before the previous
response was retrieved.

```
FUNCTION sendMyWScall()
  DEFINE wsstatus INTEGER
  # ...
  LET ws_myStub.Add.a = 6
  LET ws_myStub.Add.b = 8
  LET wsstatus = ws_myStub.AddRequest_g()
  IF wstatus <> 0 THEN
    DISPLAY "ERROR :", WSHelper.wsError.code
  END IF
  # ...
```

You can then call `AddResponse_g` to retrieve the response in the
`AddResponse` public record of the previous request. If the returned status is 0
(zero), the response was successfully received; -1 means that there was an error, and -2 means that
the response was not yet received and that the function needs to be called
later.

```
FUNCTION retrieveMyWScall()
  DEFINE wsstatus INTEGER
  # ...
  LET wsstatus = ws_myStub.AddResponse_g()  
  CASE wstatus
    WHEN -2
      DISPLAY "No response available, try later"
    WHEN 0
      DISPLAY "Response is :",ws_myStub.AddResponse.r  
    OTHERWISE
      DISPLAY "ERROR :", WSHelper.wsError.code
  END CASE
  # ...
```

You can mix the asynchronous call with the synchronous one as they are using two different
requests. In other words, you can perform an asynchronous request with
`AddRequest_g`, then a synchronous call with `Add_g`, and then
retrieve the response of the previous asynchronous request with `AddResponse_g`.

> **Important:**
>
> In development mode, a single BDL Web Service server can only handle one request at a time,
> and several asynchronous requests in a row without retrieving the corresponding response will
> lead to a deadlock. To support several asynchronous requests in a row, it is recommended that
> you are in [deployment](4908-web-services-server-program-deployment.md "The Genero Application Server (GAS) manages web services. You must consider GAS configuration when deploying your web service in a production environment.") mode
> with a GAS as the front-end.

## Child topics

- [Using functions generated with globals](4628-using-functions-generated-with-globals.md): Learn how to use the functions and global records in the stub file to write client applications that use the Web service.
