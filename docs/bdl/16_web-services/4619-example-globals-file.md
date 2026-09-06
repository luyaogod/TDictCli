---
title: "Example globals file"
source: "fgl-topics/c_gws_handlers_client_008.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > Generate the stub file > Generate the stub files (legacy) > Example globals file"
type: "concept"
---

# Example globals file

> The example WSDL file for the Calculator Web Service provides information about the service.

The generated file ws\_calculator.inc lists the prototypes for the
following functions: `Add` and `Add_g` functions, the asynchronous
`AddRequest_g` and `AddResponse_g` functions, as well as the
definitions of the global variables **Add** and `AddResponse`:

```
# Operation: Add## FUNCTION: Add_g()  -- Function that uses the global input
                                      -- and output records
#   RETURNING: soapStatus             -- An integer where 0 represents success
#   INPUT: GLOBAL Add
#   OUTPUT: GLOBAL AddResponse
#
# FUNCTION: Add(p_a, p_b)             -- Function with input parameters that 
#  RETURNING: soapStatus ,p_r         -- correspond to the a and b variables 
                                      -- of the global INPUT record
                                      -- Return values are the status integer
                                      --  and the value in the r variable of 
                                      -- the global OUTPUT record
#
# FUNCTION: AddRequest_g()            -- Asynchronous function that uses the 
                                      -- global input record
#  RETURNING: soapStatus              -- An integer where 0 represents 
#   INPUT: GLOBAL Add                 -- success, -1 error and -2 means that 
                                      -- a previous request was sent 
                                      -- and that a response is in progress.
#
# FUNCTION: AddResponse_g()           -- Asynchronous function that uses 
                                      -- the global output record
#  RETURNING: soapStatus              -- An integer where 0 represents 
#   OUTPUT: GLOBAL AddResponse        -- success, -1 error and -2 means that
                                      -- the response was not 
                                      -- yet received, and that a new call
                                      -- should be done later.

#VARIABLE : Add   -- defines the global INPUT record
DEFINE Add RECORD ATTRIBUTES(XMLName="Add",
                            XMLNamespace="http://tempuri.org/")
       a INTEGER ATTRIBUTES(XMLName="a",XMLNamespace=""),
       b INTEGER ATTRIBUTES(XMLName="b",XMLNamespace="")
    END RECORD

# VARIABLE : AddResponse   -- defines the global OUTPUT record
DEFINE AddResponse RECORD ATTRIBUTES(XMLName="AddResponse",
                          XMLNamespace="http://tempuri.org/")
       r INTEGER ATTRIBUTES(XMLName="r",XMLNamespace="")
    END RECORD
```

## Related links

**Related concepts**  

[Generating stub files for a legacy GWS client](4618-generate-the-stub-files-legacy.md "Use the fglwsdl tool to generate legacy client stub files (.inc and .4gl) compatible with apps created with Genero 3.20 or prior.")
