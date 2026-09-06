---
title: "Step 3: Write the MAIN program block"
source: "fgl-topics/c_gws_client_tutorial_006.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Call the web service > Step 3: Write the MAIN program block"
type: "concept"
description: "To explain how you code a Genero client application that accesses a SOAP Web service, this documentation uses the example of a calculator client application. This application accesses the Add ..."
---

# Step 3: Write the MAIN program block

To explain how you code a Genero client application that accesses a
SOAP Web service, this documentation uses the example of a calculator client application. This
application accesses the Add operation in the **MyCalculator** Web service. To write the client,
you simply need to access the WSDL for the Web service. To learn how to program the service, see
[Writing a Web server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .").

Provide values for the input and output messages of the operation, and call one of the generated
functions. Since the input and output messages are simple integers, we can call the **Add**
function defined in the stub file (ws\_calculator).

```
MAIN
  DEFINE op1        INTEGER
  DEFINE op2        INTEGER
  DEFINE result     INTEGER
  DEFINE wsstatus   INTEGER

  LET op1 = 1
  LET op2 = 2
  CALL ws_calculator.Add(op1, op2) RETURNING wsstatus, result
  IF wsstatus = 0 THEN
    DISPLAY "Result: ", result
  ELSE
    -- Use the wsError record
    DISPLAY "Error: ", wsError.description
  END IF
END MAIN
```

Alternatively, we can use the [input and
output](4596-generate-the-client-stub-file.md) records directly, calling the **Add\_g**
function:

```
MAIN
  DEFINE wsstatus INTEGER

  LET ws_calculator.Add.a = 1
  LET ws_calculator.Add.b = 2
  LET wsstatus = ws_myStub.Add_g()
  IF wsstatus != 0 THEN
    -- Use the wsError record
    DISPLAY "Error :", wsError.Description
  ELSE
    DISPLAY "Result: ", AddResponse.r
  END IF
END MAIN
```

These examples are very basic versions of the code. For complete examples, see the code samples
provided with the package in demo/WebServices.

## Backward compatibility for globals

If you have [generated stub
files](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.") for compatibility with your GWS client app prior to version 4.00, you have a reference
in the [`GLOBALS`](4601-step-2-import-the-stub-file-legacy.md)
statement at the top of the .4gl module that links the stub file. This means
you can reference the variables and call the function directly. In the first code
sample you can change to this method to call the stub
function:

```
# ...
CALL Add(op1, op2) RETURNING wsstatus, result
#...
```

In the
second code sample, you can set the variables and call the function as
follows:

```
# ...
 LET Add.a = 1
 LET Add.b = 2
 LET wsstatus = Add_g()
#...
```
