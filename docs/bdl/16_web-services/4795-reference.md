---
title: "REST API reference (high-level framework)"
source: "fgl-topics/c_gws_high_level_rest_api.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference"
type: "concept"
---

# REST API reference (high-level framework)

> The reference information provided in this section allow you to create RESTful Web service server and client applications using the Genero Web Services high-level framework.

See [Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.") for examples of use of the
attributes in prototype functions. The [High-level RESTful Web service attributes](4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")
reference page lists attributes currently defined for high-level RESTful functions.

Some attributes are used in the attribute clause for variables at the modular level. See the
[WSError](4805-wserror.md "Defines a description for an HTTP status code returned by the service."), [WSInfo](4804-wsinfo.md "Defines general metadata about the REST service."), [WSScope (module)](4807-wsscope-module.md "Defines the access scope required to call any REST operation in the module."), and [WSContext](4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.") attributes.

See [REST function syntax with RESTful attributes](4797-rest-function-syntax-with-restful-attributes.md "A RESTful FUNCTION definition is specified with a set of input parameter attributes, function definition attributes, and return attributes that define it as an operation for a REST web service.") for details of the exact
syntax for using attributes. See [Function attributes](../08_language-basics/0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values.") for details of the
semantics of attributes applying to functions.

## Child topics

- [Query string parameters](4796-query-string-parameters.md): You can append query string parameters to a REST service URL to request OpenAPI documents or retrieve version information about the service.
- [REST function syntax with RESTful attributes](4797-rest-function-syntax-with-restful-attributes.md): A RESTful FUNCTION definition is specified with a set of input parameter attributes, function definition attributes, and return attributes that define it as an operation for a REST web service.
- [Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md): RESTful attributes define functions for your RESTful web service.
- [High-level RESTful Web service attributes](4802-high-level-restful-web-service-attributes.md): Attributes for high-level RESTful Genero Web Services.
- [Default media types](4846-default-media-types.md): For RESTful Web services developed using the high-level framework, you can specify the MIME type or you can accept the default. The default MIME type is based on the data type.
- [OpenAPI types mapping: BDL](4847-openapi-types-mapping-bdl.md): Conversion mapping for Genero BDL data types in OpenAPI documentation.
- [Limitations of RESTful Web Services (high-level framework)](4853-rest-limitations-high-level-apis.md): Be aware of the limitations of the RESTful Web Services high-level framework.
