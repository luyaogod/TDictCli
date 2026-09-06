---
title: "Get OpenAPI description for a non-secure service"
source: "fgl-topics/t_gws_client_get_nonsecure_service_openapi_spec.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Understanding the OpenAPI description of a Genero RESTful service > Get OpenAPI (non-secure)"
type: "task"
---

# Get OpenAPI description for a non-secure service

> Retrieve the OpenAPI description from a non-secure RESTful service using the service URL with the ?openapi.json query string.

A client application needs the [OpenAPI description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.") of the RESTful service before generating a stub file or calling the
service. For non-secure services, you can retrieve the description directly from the service URL.

1. Use the service URL with the `?openapi.json` query string to
   retrieve the OpenAPI description.

   ```
   http://host:port/gas/ws/r/group/xcf/service?openapi.json
   ```

   The OpenAPI description is displayed in JSON format in your browser. You can save the output to a
   file if needed.
2. (Optional) Save the JSON output to a file if you prefer to generate the stub
   from a local specification.

   You can now generate the stub file using the saved JSON or directly from the service URL. For
   details, go to [Generate a REST stub file](4783-generate-stub-file.md "Use the fglrestful tool to generate a client stub file from a REST service URL or an OpenAPI description file.").

## Related links

**Related concepts**  

[Information needed to generate a REST stub file](4781-before-generating-the-stub-file.md "You need the OpenAPI description of the RESTful service before generating the stub file or implementing client-side calls.")

**Related tasks**  

[Get OpenAPI description for a secure service](4779-get-openapi-secure.md "Retrieve the OpenAPI description from a secure RESTful service by obtaining an access token and including it in the service URL.")

[Generate a REST stub file](4783-generate-stub-file.md "Use the fglrestful tool to generate a client stub file from a REST service URL or an OpenAPI description file.")
