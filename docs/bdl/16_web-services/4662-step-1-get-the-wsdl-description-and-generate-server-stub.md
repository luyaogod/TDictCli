---
title: "Step 1: Get the WSDL description and generate server stub"
source: "fgl-topics/c_gws_server_tutorial_011.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 2: Writing a server using third-party WSDL (the fglwsdl tool) > Step 1: Get the WSDL description and generate server stub"
type: "concept"
---

# Step 1: Get the WSDL description and generate server stub

> Use the fglwsdl tool to generate the BDL server stub from a WSDL.

This tutorial uses the Calculator service defined in [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service.") to obtain the WSDL
information. The `fglwsdl -s module-name` command generates
a corresponding BDL server stub. You must provide the name of the GWS server application
module ("serviceImp" in the sample) where the service finds the functions for the service
operations at
runtime:

```
fglwsdl -s serviceImp -soap12 -o example1 http://localhost:8090/MyCalculator?WSDL
```

The "MyCalculator" service must be running in order to obtain the WSDL information.

For an more information on using this command see [Generate server stub file](../13_programming-tools/2523-fglwsdl.md).

If you need to generate stub files that are compatible with your existing GWS client
application using `GLOBALS`, see [Get the WSDL description and generate legacy files](4663-get-the-wsdl-description-and-generate-legacy-files.md "Use the fglwsdl tool legacy option to generate legacy code (Genero 3.20 or prior) for the server stub from a WSDL.").

## The generated .4gl file

The generated stub file (example1Service.4gl in the example) contains a
single function that creates the service, publishes the operation, and registers the service. The
Web Service Style that is created is determined by the style specified in the WSDL information. The
functions in this file accomplish the same tasks as [Step 3: Create the service and operations](4658-step-3-create-the-service-and-operations.md "Describes how you provide your Web service and its operations to users who can access it on the net.") and [Step 4: Register the service](4659-step-4-register-the-service.md "Register the service with the Genero Web Services (GWS) server.") of Example 1. Do not modify this
file.

```
# example1Service.4gl
# Generated file containing the function Createexample1Service

IMPORT com
IMPORT FGL serviceImp    -- import the service operation module

# FUNCTION Createexample1Service
#   RETURNING soapstatus
FUNCTION Createexample1Service()
  DEFINE service      com.WebService
  DEFINE operation    com.WebOperation
  # Set ERROR handler
  WHENEVER ANY ERROR GOTO error
   # Create Web Service
  LET service = com.WebService.CreateWebService(
                             "MyCalculator",
                             "http://tempuri.org/webservices")

  # Publish Operation : Add
  LET operation = com.WebOperation.CreateRPCStyle(
                                 "serviceImp.Add",
                                 "Add",
                                 Add,
                                 AddResponse)
  CALL service.publishOperation(operation,"")
   # Register Service
  CALL com.WebServiceEngine.RegisterService(service)
  RETURN 0
  # ERROR handler
  LABEL error:
  RETURN status
  # Unset ERROR handler
  WHENEVER ANY ERROR STOP
END FUNCTION
```

## Related links

1. [Get the WSDL description and generate legacy files](4663-get-the-wsdl-description-and-generate-legacy-files.md)

   Use the fglwsdl tool legacy option to generate legacy code (Genero 3.20 or prior) for the server stub from a WSDL.

**Related concepts**  

[Step 2: Write a BDL function for your service operation](4664-step-2-write-a-bdl-function-for-your-service-operation.md "Write functions that implement the functions in the stub file. This allows you to create your own version of the function.")
