---
title: "Server side"
source: "fgl-topics/c_gws_soap_007.html"
breadcrumb: "Web services > Concepts > SOAP features > SOAP Fault > Server side"
type: "concept"
---

# Server side

> A Genero Web Services server can throw a SOAP fault when a processing error is encountered.

To generate a SOAP fault you need to:

- create the fault variable with [com.WebService.createFault()](../15_library-reference/3756-com-webservice-createfault.md "Creates a global fault for a Web Service object.")
- add it to your operation with [com.WebOperation.addFault](../15_library-reference/3770-the-weboperation-class.md "The com.WebOperation class provides an interface to create and manage the operations of a Genero Web Service.")
- use it with [com.WebServiceEngine.SetFaultDetail](../15_library-reference/3785-the-webserviceengine-class.md "The com.WebServiceEngine class provides an interface to manage the Web Services engine.")

For example in
$FGLDIR/demo/WebServices/calculator/server/calculatorServer.4gl, the calculator
server has a **divide\_by\_zero** SOAP fault. The SOAP fault is raised when you try to divide a
number by zero. To generate a SOAP fault proceed as follows:

## Create a SOAP fault

You define the variable to send as a SOAP fault. It can be a simple string like in this example
or a complex type. Remember to assign an XMLName to the
variable.

```
DEFINE divide_by_zero STRING ATTRIBUTES (XMLName="DividedByZero")
```

Then you provide the fault variable to the service using the function [com.WebService.createFault()](../15_library-reference/3756-com-webservice-createfault.md "Creates a global fault for a Web Service object.").

```
LET serv = com.WebService.CreateWebService("Calculator",serviceNS)
CALL serv.createFault(divide_by_zero,FALSE)
```

## Add the SOAP fault to an operation

A SOAP fault can be used by an operation in response to the client when an error has occurred in
the service request. An operation can use different SOAP faults but **only one at a
time**.

```
LET op = com.WebOperation.CreateRPCStyle("divide","Divide",divide_in,divide_out)
CALL op.addFault(divide_by_zero,NULL)
```

Here, the SOAP fault is added to the "divide" operation.

## Send the SOAP fault

Set the values to the fault variable. The fault message is sent to the client at the end of the
operation
processing.

```
LET divide_by_zero = "Cannot divide "||divide_in.a||" by zero"
CALL com.WebServiceEngine.SetFaultDetail(divide_by_zero)
```
