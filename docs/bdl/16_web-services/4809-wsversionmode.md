---
title: "WSVersionMode"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSVersionMode_module.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSVersionMode"
type: "concept"
---

# WSVersionMode

> Specifies how the service version is selected for OpenAPI generation (URI, header, or query).

## Syntax

```
WSVersionMode = "{ uri | header | query }"
```

Where:

1. WSVersionMode has three options: `uri`,
   `header`, or `query`. It is a string value setting. It can only
   contain one value.

`WSVersionMode` is an optional attribute that is only allowed in the [service information record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.")
of the module. A [WSVersion (module)](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.")
attribute must also be set , otherwise [error 9142](../15_library-reference/4483-genero-bdl-errors.md) is thrown.

## Usage

`WSVersion` function and module attributes
activate versioning mode. In the default versioning mode, the GWS automatically prefixes the
version to the operation's resource path in the OpenAPI documentation. This is the
`uri` versioning mode.

For example, if before versioning, the resource path as defined by `WSPath`
was:

`/prices/{nb}`

After versioning the operation ("v2" in the example), the resource path is:

`/v2/prices/{nb}`

Access to operations can be changed with the optional attribute `WSVersionMode` on
the [service information
record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation."). The main purpose of the service information record (defined with the
`WSInfo` attribute) is to document the REST service. If you are not setting
the version mode, the record is still needed to provide service information.

You can set the versioning mode to one of the three supported versioning methods; versioning the
URI, adding a custom header, or using a query string, depending on your preference. The GWS
generates the OpenAPI documentation and creates the functions in the stub file for the version based
on the mode. The whole REST Web service will use the same access mode.

To learn about version mode (to include how the GWS handles version mode in the OpenAPI
description of operations ), see [Version access modes](4764-version-access-modes.md "Specify how clients access service versions in REST operations.").

## Example using WSVersionMode

In this sample information record the `WSVersionMode` is set to "query".
This means access to operations of the service is made by a parameter adding a query string
to the URI, such as `/users?api-version=v2`. Client applications need to
provide a value for the query parameter (`api-version`) when making a call to
the service.

```
PUBLIC DEFINE serviceInfo RECORD ATTRIBUTES(WSInfo, 
                                           WSScope="users", 
                                           WSVersion="v2" WSVersionMode="query")
  title STRING,
  description STRING,
  termOfService STRING,
  contact RECORD
    name STRING,
    url STRING,
    email STRING
    END RECORD,
    version STRING
  END RECORD = (
    title: "my service", 
    version: "1.0", 
    contact: ( email:"helpdesk@mysite.com") )
```

## Related links

**Related concepts**  

[WSVersion (module)](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.")

[Version a REST web service API](4759-version-a-rest-service.md "Versioning your REST web service is important, especially when changes to the service would impact existing clients.")

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")
