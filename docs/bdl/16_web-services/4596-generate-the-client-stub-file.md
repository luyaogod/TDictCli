---
title: "Generate a GWS client stub file from the WSDL"
source: "fgl-topics/c_gws_client_get_wsdl.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Generate the client stub file"
type: "concept"
---

# Generate a GWS client stub file from the WSDL

> To access a SOAP Web service, retrieve the WSDL from the service provider to learn about the functions provided by the Web service. The fglwsdl command-line tool can use the WSDL to generate much of the client code for you.

To explain how you code a Genero client application that accesses a
SOAP Web service, this documentation uses the example of a calculator client application. This
application accesses the Add operation in the **MyCalculator** Web service. To write the client,
you simply need to access the WSDL for the Web service. To learn how to program the service, see
[Writing a Web server application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .").

Sample services can be found through UDDI registries or on other sites.

The GWS package provides the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool. This tool can use the
information provided by the WSDL to create a client stub file. For example, this command retrieves
the WSDL (the Web service information) for the "MyCalculator" service and creates a client stub file
named "Example2Client". The service must be running on the specified port in order to provide the
WSDL
information.

```
fglwsdl -soap12 -o ws_calculator http://localhost:8090/MyCalculator?WSDL
```

The generated client stub file (ws\_calculator.4gl in this example) is a
module containing:

- The definitions of the input and output records.
- The functions to call in your GWS client application, to perform a requested Web Service
  operation.
- The code that manages the Web Service request.

To generate legacy stub files (.inc and .4gl)
compatible with GWS client applications created with Genero 3.20 or prior, see [Generate GWS client stub files from the WSDL (legacy)](4597-generate-client-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.").

## Input and output records

The stub file contains the data types for the operations. They are defined as
`PUBLIC` modular
variables.

```
PUBLIC DEFINE Add RECORD
  ATTRIBUTES( XMLName="Add", 
             XMLNamespace="http://tempuri.org/webservices" )
  a INTEGER ATTRIBUTES( XMLName="a", XMLNamespace="" ),
  b INTEGER ATTRIBUTES( XMLName="b", XMLNamespace="" )
END RECORD

PUBLIC DEFINE AddResponse RECORD 
  ATTRIBUTES( XMLName="AddResponse", 
             XMLNamespace="http://tempuri.org/webservices" )
  r INTEGER ATTRIBUTES(XMLName="r",XMLNamespace="" )
END RECORD
```

## Function for the operations

The stub file contains functions generated for each operation. For example, the Add operation has
two functions that developers may use for different situations and data types.

The **Add** function uses input and output parameters, and returns the status and result. This
function can only be used if the input and output parameters are not complex structures such as
arrays or records. Using this function, developers do not access the public records directly.

```
FUNCTION Add(p_a, p_b)
  DEFINE p_a	 INTEGER
  DEFINE p_b	 INTEGER
  DEFINE soapStatus INTEGER

  LET Add.a = p_a
  LET Add.b = p_b

  LET soapStatus = Add_g()

  RETURN soapStatus, AddResponse.r
END FUNCTION
```

The **Add\_g** function can be used with the public input and output
records. Before calling this function, you must set the values in the variables of the public input
record.

For examples using the functions, see [Call the web service](4598-call-the-web-service.md "Import the required libraries and write the program to call a SOAP Web Service.").

## Related links

1. [Generate GWS client stub files from the WSDL (legacy)](4597-generate-client-stub-files-legacy.md)

   Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.

**Related concepts**  

[SOAP 1.1 and 1.2](4513-soap-1-1-and-1-2.md "GWS is able to communicate with Web services provided with SOAP 1.1 or SOAP 1.2.")

[WS client stubs and handlers](4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.")
