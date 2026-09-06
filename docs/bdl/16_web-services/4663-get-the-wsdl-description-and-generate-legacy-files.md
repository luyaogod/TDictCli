---
title: "Get the WSDL description and generate legacy files"
source: "fgl-topics/c_gws_server_tutorial_011_legacy.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 2: Writing a server using third-party WSDL (the fglwsdl tool) > Step 1: Get the WSDL description and generate server stub > Get the WSDL description and generate legacy files"
type: "concept"
---

# Get the WSDL description and generate legacy files

> Use the fglwsdl tool legacy option to generate legacy code (Genero 3.20 or prior) for the server stub from a WSDL.

This command gets the WSDL information for the Calculator Service defined in [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service."). It uses the
`-legacy` option of the [fglwsdl](../13_programming-tools/2523-fglwsdl.md) tool:

```
fglwsdl -s -legacy -soap12 -o example1 http://localhost:8090/MyCalculator?WSDL
```

The command generates two corresponding BDL files:

- The [globals file](4663-get-the-wsdl-description-and-generate-legacy-files.md), containing declarations of global variables that
  can be used as input or output to functions accessing the Web Service operations.
- A [.4gl file](4663-get-the-wsdl-description-and-generate-legacy-files.md) containing a function that creates the service
  described in the WSDL, publishes the operations of the service, and registers the service.

The "MyCalculator" Genero Web Services service created in [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service.") must be running in order
to obtain the WSDL information.

## The generated globals file

The globals file example1Service.inc provides the definition of the global
input and output records as described in [Step 1: Define input and output records](4656-step-1-define-input-and-output-records.md "Define records for the input and output messages of the Web function.") of the [Example 1: Writing the entire server application](4655-example-1-writing-the-entire-server-application.md "Design a simple Web service.") GWS Server program. The names of the input and output
records have been assigned by fglwsdl, in accordance with the Style of the Web
Service "MyCalculator" (created as [RPC Style](4499-web-services-style-options.md "Information on Web services Style options available for SOAP Genero Web services. There is no style concept in REST.") in the
Example1 program). Do not modify this file.

## Example input and output records:

```
# VARIABLE : Add
DEFINE Add RECORD
           ATTRIBUTES( XMLName="Add",
                      XMLNamespace="http://tempuri.org/webservices" )
  a INTEGER ATTRIBUTES(XMLName="a",XMLNamespace=""),
  b INTEGER ATTRIBUTES(XMLName="b",XMLNamespace="")
END RECORD

# VARIABLE : AddResponse
DEFINE AddResponse RECORD 
                   ATTRIBUTES(XMLName="AddResponse", 
                             XMLNamespace="http://tempuri.org/webservices" )
  r INTEGER ATTRIBUTES(XMLName="r",XMLNamespace="")
END RECORD
```

## The generated .4gl file

The example1Service.4gl file contains a single function that creates the
service, publishes the operation, and registers the service. The Web Service Style that is created
is determined by the style specified in the WSDL information. The functions in this file accomplish
the same tasks as [Step 3: Create the service and operations](4658-step-3-create-the-service-and-operations.md "Describes how you provide your Web service and its operations to users who can access it on the net.") and [Step 4: Register the service](4659-step-4-register-the-service.md "Register the service with the Genero Web Services (GWS) server.") of Example 1. Do not modify this
file.

```
# example1Service.4gl
# Generated file containing the function Createexample1Service

IMPORT com
GLOBALS "example1Service.inc"

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
   # Operation: Add
   # Publish Operation : Add
  LET operation = com.WebOperation.CreateRPCStyle(
                                 "Add",
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

**Related concepts**  

[Step 2: Write a BDL function for your service operation](4664-step-2-write-a-bdl-function-for-your-service-operation.md "Write functions that implement the functions in the stub file. This allows you to create your own version of the function.")
