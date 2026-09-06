---
title: "High-level RESTful Web service attributes"
source: "fgl-topics/r_gws_high_level_rest_api_attributes.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes"
type: "reference"
---

# High-level RESTful Web service attributes

> Attributes for high-level RESTful Genero Web Services.

This reference page describes all attributes supported by high‑level RESTful Web services.
Attributes are grouped by their area of use, including HTTP verb declarations, URL paths,
module‑level configuration, error handling, security, parameters, and function‑level metadata. Use
this page to identify the attributes available in each context and to review the syntax and behavior
of each one.

| Attribute | Description |
| --- | --- |
| WSDelete | In order to remove an existing resource, you set the WSDelete attribute. |
| WSGet | In order to retrieve a resource, you set the WSGet attribute. |
| WSHead | In order to retrieve resource headers, you define the WSHead attribute. |
| WSOptions | In order to retrieve the HTTP request methods supported by the server or the resource, you define the WSOptions attribute. |
| WSPatch | In order to partially update an existing resource, you define the WSPatch attribute. |
| WSPost | In order to create a new resource, you set the WSPost attribute. |
| WSPut | Update an existing resource with the WSPut attribute. |
| WSTrace | You define the WSTrace attribute for debugging purposes. |

| Attribute | Description |
| --- | --- |
| WSPath = "/{ path-element \| value-template } [/...]" | Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL. |

| Attribute | Description |
| --- | --- |
| WSCookie | Maps a parameter to a cookie value in the HTTP request. |
| WSParam | Maps a parameter to a value in the resource path template. |
| WSQuery | Maps a parameter to a query string value in the request URL. |

| Attribute | Description |
| --- | --- |
| WSAttachment [= "regexp_pattern"] | Defines file attachments in the REST message. |
| WSHeader | Defines a custom HTTP header for a parameter or return value. |
| WSMedia = " MIME-type [,...]" | Defines the supported media (MIME) types for a parameter or return value. |
| WSName=" description " | Specifies an alternative name for a parameter or return value in the REST message. |
| WSOptional | Marks a parameter or return value as optional in the REST message. |

| Attribute | Description |
| --- | --- |
| WSContext | Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module. |
| WSError= "description" | Defines a description for an HTTP status code returned by the service. |
| WSInfo | Defines general metadata about the REST service. |
| WSScope = "{ scope } [,...]" | Defines the access scope required to call any REST operation in the module. |
| WSTypeDescription=" description " | Provides a description for a user‑defined type in the REST service. |
| WSVersion = "{ version }" | Sets the version of the REST service for OpenAPI documentation. |
| WSVersionMode = "{ uri \| header \| query }" | Specifies how the service version is selected for OpenAPI generation (URI, header, or query). |

| Attribute | Description |
| --- | --- |
| WSErrorHeader= "{ code } [,...]" | Defines a custom HTTP header to include in an error response. |
| WSRetCode= "{ code \| code:description } " | Sets the HTTP success status code returned by the REST function. |
| WSThrows = "{ code \| code:description \| code:@variable } [,...]" | Defines the list of error codes the REST function may return. |

| Attribute | Description |
| --- | --- |
| WSDescription=" description " | Describes the REST function, parameters, return values, and members of user-defined types. |
| WSOperationId=" operationId " | Specifies a unique string to identify the operation. It corresponds to the `operationId` keyword in the OpenAPI specification. |
| WSPreProcessing = "[ wsquery \| wsheader \| wsparam ]" | Specify functions in a REST web service module that are used as preprocessing request handlers. |
| WSPostProcessing = "[ wsheader ]" | Specify functions in a REST web service module that are used as postprocessing request handlers. |
| WSScope = "{ scope } [,...]" | Specifies the security scope required to access this REST function. Use this attribute to restrict access at the function level. |
| WSSummary=" summary " | Specifies a short, human-readable summary of a REST operation. It corresponds to the `summary` keyword in the OpenAPI specification. |
| WSTags = "{ tag } [,...]" | Assign one or more tags to a REST operation. The value corresponds to the "tags" keyword in the OpenAPI specification and is used for grouping and documentation. |
| WSVersion = "{ version } [,...]" | Specifies the version or versions in which the REST function is available. |

## Child topics

- [Attributes set at the module level](4803-attributes-set-at-the-module-level.md): Attributes that apply to the whole Web service (at the module level).
- [Attributes set at the function level](4813-attributes-set-at-the-function-level.md): Attributes applied to the function that affect the Web service operation (the function).
- [Attributes set on parameters and returns](4832-attributes-set-on-parameters-and-returns.md): Attributes that affect the input parameters and return values of the function.
