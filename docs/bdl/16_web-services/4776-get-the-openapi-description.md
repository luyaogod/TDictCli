---
title: "Get the OpenAPI service description"
source: "fgl-topics/t_gws_restful_high_level_generate_description.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Understanding the OpenAPI description of a Genero RESTful service > Get the OpenAPI description"
type: "task"
---

# Get the OpenAPI service description

> Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.

OpenAPI provides a standardized model for describing RESTful web services. You can retrieve the description directly from the service URL when the service supports OpenAPI.

If you have set a [default
version](4763-set-the-default-service-version-for-openapi-documentation.md "Define which version of a service is shown by default in the OpenAPI documentation.") in the [information
record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation."), there is no need to specify the version in the query string in the URI. For example,
with http://myhost/gas/ws/r/myGroup/myXcf/myService?openapi.json the GWS
generates documentation for the version of operations set as default.

The web service must be running on the Genero Application Server to provide the output.

Retrieve the OpenAPI description in JSON or YAML.

- To retrieve the description in JSON, append `?openapi.json` to the service URL:

  http://myServer:myPort/gas/ws/r/myXcf/myService?openapi.json
- To retrieve the description in YAML, append `?openapi.yaml` to the service URL:

  http://myServer:myPort/gas/ws/r/myXcf/myService?openapi.yaml

  > **Warning:**
  >
  > The YAML output feature is experimental and may change.

The OpenAPI description provides details about the operations defined in the web service. For
more information, go to [Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.") and [How GWS maps BDL types into the OpenAPI description](4775-how-gws-maps-bdl-types-into-the-openapi-description.md "GWS maps Genero BDL types to JSON schema objects in the OpenAPI description. The mapping depends on the content type and whether the data type is named or anonymous.").

## Related links

**Related concepts**  

[Information needed to generate a REST stub file](4781-before-generating-the-stub-file.md "You need the OpenAPI description of the RESTful service before generating the stub file or implementing client-side calls.")

[Access to secure web service](4787-access-to-secure-web-service.md "To access a secure RESTful Web service, the client application must have a valid access token.")

**Related tasks**  

[Get the OpenAPI description for a specific version](4777-get-openapi-description-for-a-version.md "Retrieve the OpenAPI description for a specific version of a RESTful web service.")
