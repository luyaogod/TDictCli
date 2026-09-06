---
title: "Known issues"
source: "fgl-topics/c_gws_installation_known_issues.html"
breadcrumb: "Web services > General > Known issues"
type: "concept"
---

# Known issues

> There are some known issues when working with Web services.

These issues are specific to SOAP Web Services.

## Forcing RPC style convention when no input message

In RPC style, the convention defines names for input messages and output messages,
but if there is no input message, its name cannot be redefined.

To workaround this issue, respect RPC style convention in WSDL, or force RPC
convention (on client and server side) by using the `-fRPC` option of the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool.

## Variable names conflicts with library names

The fglwsdl tool can
generate variable names conflicting with `IMPORT` library names.

For example:

```
DEFINE xml xml.DomDocument
```

will conflict with the
`xml` library, if the code defines also the
instruction:

```
IMPORT xml
```
