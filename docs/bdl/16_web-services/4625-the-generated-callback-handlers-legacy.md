---
title: "The generated callback handlers (legacy)"
source: "fgl-topics/c_gws_handlers_client_006_legacy.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > WS client stubs and handlers > The generated callback handlers > The generated callback handlers (legacy)"
type: "concept"
---

# The generated callback handlers (legacy)

> Understand how to generate legacy stub files (.inc and .4gl) for compatibility with legacy code (Genero 3.20 or prior) in your client stub.

More and more Web services provide support for the different WS-\* specifications. To enable a
better interoperability with such services, the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool allows
the programmer to modify the SOAP request before it is sent, and to perform additional verifications
of the SOAP response before it is returned from the BDL function.

If you need to generate stub files (.inc and .4gl) for
compatibility with your legacy code (Genero 3.20 or prior), the fglwsdl tool has
a `-legacy` option for this purpose. For example, this command performs several
operations at
once:

```
fglwsdl -domHandler -legacy -soap12 -o myDomStub http://localhost:8090/MyService?WSDL
```

- It generates the client stub with DomHandler callbacks based on the WSDL for the service running
  on the localhost.
- It generates additional calls for each operation of a service to execute one of the [three callback handlers](4624-handler-definition.md "Describes the various callback handlers and how they are defined.") you have to implement.
- It generates the globals file (.inc) containing the definitions of the
  input and output records, and prototypes of the operations.
