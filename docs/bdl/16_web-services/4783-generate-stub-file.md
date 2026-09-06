---
title: "Generate a REST stub file"
source: "fgl-topics/t_gws_rest_stub_generate.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Before generating the stub file > Generate stub file"
type: "task"
---

# Generate a REST stub file

> Use the fglrestful tool to generate a client stub file from a REST service URL or an OpenAPI description file.

The stub file is generated from the [OpenAPI description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.") of the RESTful
service. You can obtain the description directly from the service URL or use an existing OpenAPI
file.

1. Generate the stub file from the service URL.

   ```
   fglrestful -o clientStub http://host:port/gas/ws/r/group/xcf/service?openapi.json
   ```

   The query string ?openapi.json retrieves the OpenAPI description from the
   running service. The generated file is named clientStub.4gl.
2. Generate the stub file from an OpenAPI description file.

   ```
   fglrestful -o clientStub service-spec-file
   ```

   Use this option when you already have a local OpenAPI file provided by the service developer.
3. Confirm that the stub file was created.

   The generated clientStub.4gl file contains all the types and functions
   needed to call the service. Compile and import it into your Genero REST client application. For
   details, see [REST stub file overview](4782-rest-stub-file-overview.md "The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.").

## Related links

**Related concepts**  

[Information needed to generate a REST stub file](4781-before-generating-the-stub-file.md "You need the OpenAPI description of the RESTful service before generating the stub file or implementing client-side calls.")

[REST stub file overview](4782-rest-stub-file-overview.md "The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.")
