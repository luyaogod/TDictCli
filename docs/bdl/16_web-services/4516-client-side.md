---
title: "Client side"
source: "fgl-topics/c_gws_soap_008.html"
breadcrumb: "Web services > Concepts > SOAP features > SOAP Fault > Client side"
type: "concept"
---

# Client side

> A Genero Web Services client can receive a SOAP fault number in the operation status and act accordingly.

If a SOAP fault occurs, the operation returns the SOAP fault number in the operation status. The
SOAP fault number is defined in the generated stubs as a BDL constant prefixed with the
string `FaultID_`. A SOAP fault can occur in case of HTTP error 200 and
500.

For example in $FGLDIR/demo/WebServices/calculator/client/ws\_calculator.inc,
the `Divide` operation has a SOAP fault that informs the client when a number is
divided by zero.

```
# List of Soap fault constants
CONSTANT FaultID_DividedByZero = 1
...
# VARIABLE : DividedByZero
DEFINE DividedByZero STRING ATTRIBUTES(XMLName="DividedByZero",
  XMLNamespace="http://tempuri.org/")
...
# Operation: Divide
# FAULT #1: GLOBALS DividedByZero
```

You can test the operation status code accordingly and display
the SOAP fault message.

For example in $FGLDIR/demo/WebServices/calculator/client/calculatorClient.4gl,
when the divide operation status is 1, `DividedByZero` message is displayed.

```
ON ACTION divide
   CALL Divide(op1, op2) RETURNING wsstatus, result, remaind
   CASE wsstatus
      WHEN 0
         DISPLAY BY NAME result,remaind
         DISPLAY "OK" TO msg
      WHEN FaultID_DividedByZero
         DISPLAY DividedByZero TO msg
      OTHERWISE
         DISPLAY wsError.description TO msg
   END CASE
```
