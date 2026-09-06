---
title: "Define HTTP variables"
source: "fgl-topics/c_gws_function_declaration_003a_http_variables.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web services server function > Define HTTP variables"
type: "concept"
---

# Define HTTP variables

> Define variables for the HTTP request and response communication of the service.

Types defined in the [WSHelper](4923-wshelper-library.md "The WSHelper library provides a set of functions to help with web services.") module provide global
variables that register methods for handling input and output messages. Define public variables that
reference these in your module:

## Example: HTTP input variable

```
# HTTP INPUT VARIABLE : HttpIn
PUBLIC DEFINE HttpIn WSHelper.tGlobalServerHttpInputVariableType
```

## Example: HTTP output variable

```
# HTTP OUTPUT VARIABLE : HttpOut
PUBLIC DEFINE HttpOut WSHelper.tGlobalServerHttpOutputVariableType
```
