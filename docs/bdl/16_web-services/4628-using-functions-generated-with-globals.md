---
title: "Using functions generated with globals"
source: "fgl-topics/c_gws_handlers_client_009_globals.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > Use the generated functions > Using functions generated with globals"
type: "concept"
---

# Using functions generated with globals

> Learn how to use the functions and global records in the stub file to write client applications that use the Web service.

If you have [generated stub
files](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.") for compatibility with your existing GWS client application using
`GLOBALS`, the globals file contains record definitions and the function prototypes
obtained from the WSDL.

For example, the information about the Add function from the
ws\_calculator.inc file, allows you to write code in your own
.4gl module that uses this operation as part of your client application.

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
  CALL Add(op1, op2) 
    RETURNING wsstatus, result ...
  DISPLAY result
```

## Using global records

An alternative option is to call the `Add_g` function instead, using the global
records `Add` and `AddResponse` directly. If the input variables are
complex structures like records or arrays, you are required to use this
function.

```
FUNCTION myWScall()
  DEFINE wsstatus INTEGER
  ...
  LET Add.a = 6
  LET Add.b = 8
  LET wsstatus = Add_g()
  ...
  DISPLAY AddResponse.r
```

In this case, the status is returned by the function, which has also put the result in the
`AddResponse` global record.

See [Tutorial: Writing a
Client Application](4595-steps-to-write-a-gws-client.md "A SOAP Web service contains functions that you need to call. Create a Genero BDL application that requests the functions of the service.") for more information. The demo/WebServices
subdirectory of your Genero installation directory contains complete examples of Client
Applications.

## Using asynchronous calls

If you don't want your application to be blocked when waiting for the response to a request, it
is recommended that you first call `AddRequest_g`; this will send the request using
the global `Add` record to the server. It returns a status of 0 (zero) if everything
goes well, -1 in case of error, or -2 if you tried to resend a new request before the previous
response was retrieved.

```
FUNCTION sendMyWScall()
  DEFINE wsstatus INTEGER
  ...
  LET Add.a = 6
  LET Add.b = 8
  LET wsstatus = AddRequest_g()
  IF wstatus <> 0 THEN
    DISPLAY "ERROR :", wsError.code
  END IF
  ...
```

You can then call `AddResponse_g` to retrieve the response in the
`AddResponse` global record of the previous request. If the returned status is 0
(zero), the response was successfully received; -1 means that there was an error, and -2 means that
the response was not yet received and that the function needs to be called
later.

```
FUNCTION retrieveMyWScall()
  DEFINE wsstatus INTEGER
  ...
  LET wsstatus = AddResponse_g()  
  CASE wstatus
    WHEN -2
      DISPLAY "No response available, try later"
    WHEN 0
      DISPLAY "Response is :",AddResponse.r  
    OTHERWISE
      DISPLAY "ERROR :", wsError.code
  END CASE
  ...
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
