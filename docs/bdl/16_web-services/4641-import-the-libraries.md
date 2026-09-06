---
title: "Import the libraries"
source: "fgl-topics/c_gws_function_declaration_002_import.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web services server function > Import the libraries"
type: "concept"
description: "The methods associated with creating and publishing a Web Service are contained in the classes that make up the Genero Web Services Library ( com ). If your application uses other classes, you need to ..."
---

# Import the libraries

The methods associated with creating and publishing a Web Service are contained in the classes
that make up the Genero Web Services Library (`com`). If your application uses other
classes, you need to import those. If you are using the [wsError](4604-handle-gws-server-errors.md "When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.") record, you must import the
[WSHelper](4923-wshelper-library.md "The WSHelper library provides a set of functions to help with web services.") module.

If the functions where the operations of the service are defined is in a separate Genero BDL
module, ( service\_implementation in the example), you need to import this so that
the service module can publish the service operation.

```
IMPORT com
IMPORT xml
IMPORT FGL WSHelper
IMPORT FGL service_implementation
```

## Backward compatibility for globals

If you are creating a service that uses legacy code (Genero 3.20 or prior)
(.inc and .4gl), import Genero Web Services Library
(`com`) and any other classes your service application
uses.

```
IMPORT com
```
