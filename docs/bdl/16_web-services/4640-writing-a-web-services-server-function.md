---
title: "Writing a Web services server function"
source: "fgl-topics/c_gws_function_declaration_001.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web services server function"
type: "concept"
---

# Writing a Web services server function

> You create a standard Genero function and publish it as a Web function (Web services operation) using methods from the classes in the COM library.

There are some restrictions on the function - input and output parameters are not allowed. By
using global or module variables, however, you can work around this exception.

See also [Tutorial:
Writing a GWS Server Application](4653-writing-a-web-server-application.md "Follow examples showing you how to write a complete Web service application for the SOAP protocol. .")

The steps for writing a Web Services
function:

## Related links

1. [Import the libraries](4641-import-the-libraries.md)
2. [Define the input parameters](4642-define-the-input-parameters.md)

   Define a record for the input message of the Web function.
3. [Define the output parameters](4643-define-the-output-parameters.md)

   Define a record for the output message of the function.
4. [Define HTTP variables](4644-define-http-variables.md)

   Define variables for the HTTP request and response communication of the service.
5. [Write the BDL function](4645-write-the-bdl-function.md)

   Your function defines an operation of the service.
6. [Create and publish the Web services operation](4646-create-and-publish-the-web-services-operation.md)

   Provide your Web service and its operation to users who can access it on the net.
