---
title: "Get available versions of a service"
source: "fgl-topics/t_gws_high_level_rest_retrieve_versions.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Understanding the OpenAPI description of a Genero RESTful service > Get available versions of a service"
type: "task"
---

# Get available versions of a service

> Retrieve the list of available versions of a REST service and identify the default version if defined.

When a REST service uses [`WSVersion`](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation."), GWS provides a way to retrieve all available versions of the
service, as well as the default version if one is defined. This is done by extending the service URL
with a dedicated [query
parameter](4796-query-string-parameters.md).

GWS evaluates service versions based on the following rules:

- Versions are returned in lexicographical order.
- No semantic or numeric interpretation of version values is performed.

**Steps to generate the version list**

1. Append the `version`
   [query parameter](4796-query-string-parameters.md)
   to the service URL, followed by the response format suffix.

   The presence of the `version` query parameter triggers version listing mode. Any
   value associated with the parameter is ignored.

   The suffix (`.json` or `.yaml`) defines the response format.

   **JSON format
   request:**

   http://myServer:myPort/gas/ws/r/myXcf/myService?version.json

   **YAML
   format
   request:**

   http://myServer:myPort/gas/ws/r/myXcf/myService?version.yaml
2. Parse the response to retrieve the list of versions and the default version.

   **JSON format response:**

   ```
   {
   "list": [ "version1", "version2", "version3" ],
   "default": "version2"
   }
   ```

   **YAML format response:**

   ```
   list:
   - version1
   - version2
   - version3
   default: version2
   ```

The response provides:

- A `list` property containing all declared versions in lexicographical order.
- A `default` property identifying the default version, if one is defined.
  If no default version exists, this property is omitted.

## Related links

**Related concepts**  

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")

[WSVersion (module)](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.")
