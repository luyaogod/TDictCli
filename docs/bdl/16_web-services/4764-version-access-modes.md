---
title: "Version access modes"
source: "fgl-topics/c_gws_high_level_rest_version_access_mode.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Advanced REST API versioning > Service-level versioning > Version access modes"
type: "concept"
---

# Version access modes

> Specify how clients access service versions in REST operations.

A web service [information
record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.") may include the `WSVersionMode` attribute, which
defines how clients access versioned operations. The mode applies to all operations in the REST
service.

The available version modes are:

- `uri`
- `header`
- `query`

In this example, `WSVersionMode` is set to `"query"`, which applies
the mode to the entire service:

```
PUBLIC DEFINE myInfo RECORD ATTRIBUTES(WSInfo, WSVersion = "v3", 
                                      WSVersionMode = "query")
   title STRING
   # ...
END RECORD
```

## Default version mode (URI)

If `WSVersionMode` is not set, or set to `"uri"`, GWS adds the
version directly to the resource path.

The OpenAPI document reflects this by prefixing the resource path with the version.

The following figures show examples for a service with one resource and for a service with
multiple resources.

![OpenAPI JSON description of a service with one resource using uri to set version mode.](../_images/rest_openapi_api_version_wsversionmode_uri_one_resource.png)

*Version mode uri ( service with one resource)*

![OpenAPI JSON description of a service with many resources using uri to set version mode.](../_images/rest_openapi_api_version_wsversionmode_uri_many_resources.png)

*Version mode uri (service with many resources)*

## Query mode

If `WSVersionMode` is set to `"query"`, GWS adds the
`api-version` query parameter to the OpenAPI description.

- The client must send the required version using this query parameter.
- The OpenAPI document includes the parameter definition using the `enum`
  keyword.
- GWS appends the query string automatically in requests, for example:
  `?api-version=v3`

![OpenAPI JSON description of a service using a query to set version mode.](../_images/rest_openapi_api_version_wsversionmode_query.png)

*Version mode query*

## Header mode

If `WSVersionMode` is set to `"header"`, GWS adds the
`api-version` HTTP header parameter to the OpenAPI description.

- The client must include this header in each request.
- The OpenAPI document includes the header definition using the `enum`
  keyword.

![OpenAPI JSON description of a service using a header to set version mode.](../_images/rest_openapi_api_version_wsversionmode_header.png)

*Version mode header*
