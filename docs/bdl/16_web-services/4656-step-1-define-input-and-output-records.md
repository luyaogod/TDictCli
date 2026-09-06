---
title: "Step 1: Define input and output records"
source: "fgl-topics/c_gws_server_tutorial_005.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Example 1: Writing the entire server application > Step 1: Define input and output records"
type: "concept"
---

# Step 1: Define input and output records

> Define records for the input and output messages of the Web function.

Based on the desired functionality of the operations that you plan for the Service, define the
input and output records for each operation. BDL functions that are written to implement a Web
Service operation cannot have input parameters or return values. Instead, each function's input and
output message must be defined as a public or module RECORD.

## The Input message

The fields of the public or module record represent each of the input parameters of the Web
function. The name of each field in the record corresponds to the name used in the SOAP request.
These fields are filled with the contents of the SOAP request by the Web Services engine just before
executing the corresponding BDL function.

## The Output message

The fields of the public or module record represent each of the output parameters of the Web
function. The name of each field in the record corresponds to the name used in the SOAP request.
These fields are retrieved from the Web Services engine immediately after executing the BDL
function, and sent back to the client.

Your Genero Web Services service has one planned operation that adds two integers and returns the
result. The input and output records are defined as follows:

## Define input and output records

```
# input record
PUBLIC DEFINE add_in RECORD 
  a INTEGER,
  b INTEGER
END RECORD

# output record
PUBLIC DEFINE add_out RECORD 
  r INTEGER
END RECORD
```

## Related links

**Related concepts**  

[Step 2: Write a BDL function for each service operation](4657-step-2-write-a-bdl-function-for-each-service-operation.md "Each function defines an operation of the service.")
