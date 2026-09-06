---
title: "Information needed to generate a REST stub file"
source: "fgl-topics/c_gws_client_get_service_openapi_spec.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Before generating the stub file"
type: "concept"
---

# Information needed to generate a REST stub file

> You need the OpenAPI description of the RESTful service before generating the stub file or implementing client-side calls.

Before coding a RESTful client application, ensure that you have the [OpenAPI description](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.") for the service
you intend to call. The description defines the available operations, the required parameters, and
the data structures used for request and response bodies. This information is essential for
generating the REST stub file and for understanding how to implement the client-side logic.

You can retrieve the OpenAPI description directly from the service URL or by using the [fglrestful](../13_programming-tools/2524-fglrestful.md) tool
included with the Genero Web Services package. The retrieved description becomes the input to the
stub-generation process.

If the service is secured, authentication may be required before accessing the OpenAPI
description. For details, go to [Get OpenAPI description for a secure service](4779-get-openapi-secure.md "Retrieve the OpenAPI description from a secure RESTful service by obtaining an access token and including it in the service URL.").

Some services publish multiple API versions. Use the version that matches the service you intend
to call, as the stub file generated from the OpenAPI description must correspond to the version
deployed on the server.

## Related links

**Related concepts**  

[REST stub file overview](4782-rest-stub-file-overview.md "The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.")

## Child topics

- [REST stub file overview](4782-rest-stub-file-overview.md): The stub file provides client-side functions that call the operations of a RESTful Web service. It acts as a proxy layer between your client application and the service.
- [Generate a REST stub file](4783-generate-stub-file.md): Use the fglrestful tool to generate a client stub file from a REST service URL or an OpenAPI description file.
- [Troubleshooting stub file creation](4784-troubleshooting.md): Errors or warnings reported during stub file creation indicate issues in the OpenAPI description or in the data types defined by the service. This topic describes how to interpret and resolve these messages.
