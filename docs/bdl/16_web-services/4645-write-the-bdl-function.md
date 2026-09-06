---
title: "Write the BDL function"
source: "fgl-topics/c_gws_function_declaration_004.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web services server function > Write the BDL function"
type: "concept"
---

# Write the BDL function

> Your function defines an operation of the service.

A Web service function is a normal BDL function that uses the [input](4642-define-the-input-parameters.md "Define a record for the input message of the Web function.") and [output](4643-define-the-output-parameters.md "Define a record for the output message of the function.") records that you have defined.

## Example add operation

```
PUBLIC FUNCTION add()
   LET AddResponse.r = add.a + add.b
END FUNCTION
```
