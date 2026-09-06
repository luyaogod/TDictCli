---
title: "The generated callback handlers"
source: "fgl-topics/c_gws_handlers_client_006.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > The generated callback handlers"
type: "concept"
---

# The generated callback handlers

> Understand the function of callback handlers, and how to generate them in your client stub.

More and more Web services provide support for the different WS-\* specifications. To enable a
better interoperability with such services, the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool allows
the programmer to modify the SOAP request before it is sent, and to perform additional verifications
of the SOAP response before it is returned from the BDL function.

A fglwsdl command with the `-domHandler
module-name` option creates a client stub file based entirely on the DOM
API. This eases the manipulation of the XML requests and responses. In the
module-name you must provide the name of the module where you implement the [three callback handlers](4624-handler-definition.md "Describes the various callback handlers and how they are defined.") for the service operations.

For example, this command performs several operations at
once:

```
fglwsdl -domHandler myCallbacks -soap12 -o myDomStub http://localhost:8090/MyService?WSDL
```

- It generates the client stub with DomHandler callbacks based on the WSDL for the service running
  on the localhost.
- In the generated stub file (.4gl) it includes an import statement
  (`IMPORT FGL myCallbacks` in the sample) to reference the callback functions at
  runtime.
- In the generated stub file it adds calls for each operation of a service to execute callbacks you
  implement in the callback module (`myCallbacks` in the example).

## Related links

**Related concepts**  

[The generated callback handlers (legacy)](4625-the-generated-callback-handlers-legacy.md "Understand how to generate legacy stub files (.inc and .4gl) for compatibility with legacy code (Genero 3.20 or prior) in your client stub.")

## Child topics

- [Handler definition](4624-handler-definition.md): Describes the various callback handlers and how they are defined.
- [The generated callback handlers (legacy)](4625-the-generated-callback-handlers-legacy.md): Understand how to generate legacy stub files (.inc and .4gl) for compatibility with legacy code (Genero 3.20 or prior) in your client stub.
