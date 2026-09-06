---
title: "Including the web services library"
source: "fgl-topics/c_gws_server_tutorial_003.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Writing a Web server application > Including the web services library"
type: "concept"
---

# Including the web services library

> Import the com class.

The methods associated with creating and publishing a Web Service are contained in the classes
that make up the [Genero Web Services
Library (com)](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.").

The XML library allows you to use attributes to map the BDL data types in a Genero application to
their corresponding XML data types. See [Attributes to Customize XML
Mapping](4958-xml-serialization-rules-and-customization.md) for additional information.

The [WSHelper](4923-wshelper-library.md "The WSHelper library provides a set of functions to help with web services.") library allows you to reference global HTTP
server variable type definitions for incoming and outgoing communication with the GWS server, and to
use its [wsError](4604-handle-gws-server-errors.md "When a Genero Web Services service operation returns a status that is non-zero, you can get a more detailed error description from the record wsError.") record for error
handling.

Include these lines at the top of each module of your GWS server application to import the
libraries:

```
IMPORT com
IMPORT xml
IMPORT FGL WSHelper
```
