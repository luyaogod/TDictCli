---
title: "fgl_ws_server_publishFunction() (version 1.3)"
source: "fgl-topics/c_gws_server_API_fgl_ws_server_publishFunction.html"
breadcrumb: "Web services > Reference > Server API functions - version 1.3 only > fgl_ws_server_publishFunction() (version 1.3)"
type: "concept"
---

# fgl_ws_server_publishFunction() (version 1.3)

> Publishes the given BDL function as a Web-Function on the Web.

> **Warning:**
>
> This function is valid for backward compatibility, but is not a preferred way to handle Genero
> Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for the preferred
> classes and methods for handling Web services.

## Syntax

```
FUNCTION fgl_ws_server_publishFunction(
   operationName VARCHAR,
   inputNamespace VARCHAR,
   inputRecordName VARCHAR,
   outputNamespace VARCHAR,
   outputRecord VARCHAR,
   functionName VARCHAR)
```

1. operationName is the name by which the operation will be defined on the Web.
   The name is case sensitive.
2. inputNamespace is the namespace of the incoming operation message.
3. inputRecordName is the name of the BDL record representing the Web Function
   input message or `""` if there is none.
4. outputNamespace is the namespace of the outgoing operation message.
5. outputRecord is the name of the BDL record representing the Web Function
   output message or `""` if there is none.
6. functionName is the name of the BDL function that is executed when the Web
   Service engine receives a request with the operation name defined above.

## Example

```
CALL fgl_ws_server_publishFunction(
 "MyWebOperation",
 "http://www.tempuri.org/webservices/","myfunction_input",
 "http://www.tempuri.org/webservices/","myfunction_output",
 "my_bdl_function")
```

## Possible runtime errors

- -15503: FUNCTION\_ALREADY\_EXISTS
- -15501: FUNCTION\_ERROR
- -15502: FUNCTION\_DECLARATION\_ERROR
- -15512: INPUT\_VARIABLE\_ERROR
- -15513: OUTPUT\_VARIABLE\_ERROR
- -15503: BDL\_XML\_ERROR
- -15518: INPUT\_NAMESPACE\_MISSING
- -15519: OUTPUT\_NAMESPACE\_MISSING
