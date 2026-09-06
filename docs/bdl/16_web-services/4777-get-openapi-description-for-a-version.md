---
title: "Get the OpenAPI description for a specific version"
source: "fgl-topics/t_gws_restful_get_openapi_version_description.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Understanding the OpenAPI description of a Genero RESTful service > Get OpenAPI description for a version"
type: "task"
---

# Get the OpenAPI description for a specific version

> Retrieve the OpenAPI description for a specific version of a RESTful web service.

If the service defines multiple versions, you can request the [OpenAPI description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.") for a specific
version of the operations. The service must have a `WSVersion` definition for
versioned operations.

Retrieve the OpenAPI description for a specific version.

Append the [query
parameter](4796-query-string-parameters.md)
`&version=version_name` to the service URL to retrieve the
description for that version.

The following example requests the OpenAPI description for version v3:

http://myServer:myPort/gas/ws/r/myXcf/myService?openapi.json&version=v3

For details about how OpenAPI documents describe versions, go to [Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.") and [How GWS maps BDL types into the OpenAPI description](4775-how-gws-maps-bdl-types-into-the-openapi-description.md "GWS maps Genero BDL types to JSON schema objects in the OpenAPI description. The mapping depends on the content type and whether the data type is named or anonymous.").

## Related links

**Related concepts**  

[WSVersion (module)](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.")

**Related tasks**  

[Get the OpenAPI service description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.")
