---
title: "Generate GWS client stub files from the WSDL (legacy)"
source: "fgl-topics/c_gws_client_get_wsdl_globals.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Generate the client stub file > Generate client stub files (legacy)"
type: "concept"
---

# Generate GWS client stub files from the WSDL (legacy)

> Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.

This example command requests the Web service information for the "MyCalculator" service. The
[`fglwsdl
-legacy`](../13_programming-tools/2523-fglwsdl.md) option specifies to generate legacy code, and the output files
(.inc and .4gl) will have the base name
"ws\_calculator".

```
fglwsdl -soap12 -legacy -o ws_calculator http://localhost:8090/Calculator?WSDL
```

The
"MyCalculator" GWS service must be running on the specified port in order to provide the
WSDL information.

This generates two files:

- ws\_calculator.inc - the globals file containing the definitions of the
   input and
  output records, and the prototypes of the operations.
- ws\_calculator.4gl - a module containing prototype definitions of the
  functions that can be used in your GWS client application to perform the requested Web Service
  operation, and the code that manages the Web Service request.

## Input and Output records

The globals file contains record definitions for the data types of the operations obtained from
the WSDL. They are defined as variables within the `GLOBALS` instruction.

```
GLOBALS 
#...
DEFINE Add RECORD 
  ATTRIBUTES( XMLName="Add", 
             XMLNamespace="http://tempuri.org/webservices" )
  a INTEGER ATTRIBUTES( XMLName="a", XMLNamespace="" ),
  b INTEGER ATTRIBUTES( XMLName="b", XMLNamespace="" )
END RECORD

DEFINE AddResponse RECORD 
  ATTRIBUTES( XMLName="AddResponse", 
             XMLNamespace="http://tempuri.org/webservices" )
  r INTEGER ATTRIBUTES(XMLName="r",XMLNamespace="" )
END RECORD

#...
END GLOBALS
```

## Function prototypes for the operations

The globals file contains prototypes for the generated functions for each operation. For example,
the Add operation has two functions that developers may use for different situations and data types.

The **Add** function uses input and output parameters, and returns the status and result. This
function can only be used if the input and output parameters are not complex structures such as
arrays or records. Using this function, developers do not access the global records directly.

The **Add\_g** function can be used with the global input and
output records. Before calling this function, you must set the values in the variables of the
global input record.

```
Operation: Add
#
# FUNCTION: Add_g()
#   RETURNING: soapStatus
#   INPUT: GLOBAL Add
#   OUTPUT: GLOBAL AddResponse
#
# FUNCTION: Add(p_a, p_b)
#   RETURNING: soapStatus ,p_r
```

For examples using the functions, see [Call the web service](4598-call-the-web-service.md "Import the required libraries and write the program to call a SOAP Web Service.").

## Related links

**Related concepts**  

[SOAP 1.1 and 1.2](4513-soap-1-1-and-1-2.md "GWS is able to communicate with Web services provided with SOAP 1.1 or SOAP 1.2.")

[Generate a GWS client stub file from the WSDL](4596-generate-the-client-stub-file.md "To access a SOAP Web service, retrieve the WSDL from the service provider to learn about the functions provided by the Web service. The fglwsdl command-line tool can use the WSDL to generate much of the client code for you.")

[WS client stubs and handlers](4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.")
