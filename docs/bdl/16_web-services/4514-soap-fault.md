---
title: "SOAP Fault"
source: "fgl-topics/c_gws_soap_006.html"
breadcrumb: "Web services > Concepts > SOAP features > SOAP Fault"
type: "concept"
---

# SOAP Fault

> Genero Web Services supports SOAP's built-in error handling.

Since 2.40, Genero Web Services supports SOAP fault.

For backward compatibility, the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool
provides the `-ignoreFaults` option to disable SOAP fault management.

## Child topics

- [Server side](4515-server-side.md): A Genero Web Services server can throw a SOAP fault when a processing error is encountered.
- [Client side](4516-client-side.md): A Genero Web Services client can receive a SOAP fault number in the operation status and act accordingly.
