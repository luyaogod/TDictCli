---
title: "Example output"
source: "fgl-topics/c_gws_handlers_server_004.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > WS server stubs and handlers > Example output"
type: "concept"
---

# Example output

> What you can expect to find in the server stub files.

In the generated file ws\_calculatorService.inc, the definitions of the
variables for the input and output record are the same as those generated for the Web Service client
application:

```
#VARIABLE : Add -- defines the global INPUT record
DEFINE Add RECORD ATTRIBUTES(XMLName="Add",
                            XMLNamespace="http://tempuri.org/")
  a INTEGER ATTRIBUTES(XMLName="a",XMLNamespace=""),
  b INTEGER ATTRIBUTES(XMLName="b",XMLNamespace="")
END RECORD

# VARIABLE : AddResponse -- defines the global OUTPUT record
DEFINE AddResponse RECORD ATTRIBUTES(XMLName="AddResponse",
                                    XMLNamespace="http://tempuri.org/")
  r INTEGER ATTRIBUTES(XMLName="r",XMLNamespace="")
END RECORD
```

The generated file ws\_calculatorService.4gl contains a single function that
creates the Calculator service, creates and publishes the service operations, and registers the
Calculator
service:

```
FUNCTION Createws_calculatorService()
  DEFINE service com.WebService
  DEFINE operation com.WebOperation
  ...  # Create Web Service
  LET service = com.WebService.CreateWebService("Calculator",
    "http://tempuri.org/")
  # Publish Operation : Add
  LET operation = com.WebOperation.CreateRPCStyle("Add","Add",
    Add,AddResponse)
  CALL service.publishOperation(operation,"")  ...
  # Register Service
  CALL com.WebServiceEngine.RegisterService(service)
  RETURN 0
  ...
END FUNCTION
```
