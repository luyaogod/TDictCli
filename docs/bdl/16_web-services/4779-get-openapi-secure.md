---
title: "Get OpenAPI description for a secure service"
source: "fgl-topics/t_gws_client_get_secure_service_openapi_spec.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Understanding the OpenAPI description of a Genero RESTful service > Get OpenAPI (secure)"
type: "task"
---

# Get OpenAPI description for a secure service

> Retrieve the OpenAPI description from a secure RESTful service by obtaining an access token and including it in the service URL.

Secure RESTful services require an access token before you can retrieve the [OpenAPI description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML."). Use the
`GetToken` tool to obtain the token, then include it in the Web service URL.

The `GetToken` tool is located in the
FGLDIR/web\_utilities/services/gip/bin/gettoken
directory.

1. Use the `GetToken` tool to obtain an access token.

   ```
   GetToken password -u user -p mypw --idp https://host:port/gas/ws/r/services/GeneroIdentityProvider --savetofile mytoken.json myWSScope
   ```

   The token is saved to mytoken.json.

   The token is valid for a limited time (for example, 10 minutes).
2. Add the token to the Web service URL to retrieve the OpenAPI description.

   ```
   http://host:port/gas/ws/r/group/xcf/service?openapi.json&access_token=token
   ```

   The OpenAPI description is displayed in JSON format in your browser. You can save the output to a
   file if needed.
3. (Optional) Save the OpenAPI description to a file for later use.

   After retrieving the description, you can generate the REST stub file. See [Generate a REST stub file](4783-generate-stub-file.md "Use the fglrestful tool to generate a client stub file from a REST service URL or an OpenAPI description file.")

## Related links

**Related tasks**  

[Get OpenAPI description for a non-secure service](4780-get-openapi-non-secure.md "Retrieve the OpenAPI description from a non-secure RESTful service using the service URL with the ?openapi.json query string.")

[Generate a REST stub file](4783-generate-stub-file.md "Use the fglrestful tool to generate a client stub file from a REST service URL or an OpenAPI description file.")
