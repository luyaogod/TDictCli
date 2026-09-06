---
title: "Define your resource operations"
source: "fgl-topics/c_gws_restful_high_level_operations.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Define your resource operations"
type: "concept"
---

# Define your resource operations

> Operations are the HTTP verbs used to manipulate the resources of your Web service.

In GWS REST the attributes that perform operations are identifiable by the HTTP verb in the name,
for example, [WSGet](4816-wsget.md "In order to retrieve a resource, you set the WSGet attribute."), [WSDelete](4815-wsdelete.md "In order to remove an existing resource, you set the WSDelete attribute."), [WSPost](4820-wspost.md "In order to create a new resource, you set the WSPost attribute."), and [WSPut](4821-wsput.md "Update an existing resource with the WSPut attribute."), etc. Your REST function must have one of
these attributes in order to define its service. You can use these attributes in the attributes
clause of your function.

## Publishing operations

When you generate the service description, the operations you declared are published in a node.
For example, there is a "get" node if it describes a `GET` operation. You find the
operation node under the path for that resource in your [OpenAPI description](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")  file.

## Related links

**Related concepts**  

[HTTP verbs and attributes](4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.")

[Designing REST Web services](4704-designing-rest-web-services.md "Identifying the resources to expose to a RESTful Web service client is an essential step in designing a RESTful Web service server.")

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")

[REST resource URI naming practice](4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.")

[HTTP operation attributes (Verbs)](4814-http-operation-attributes-verbs.md "Attributes that map an HTTP operation or verb action in a function to a REST resource.")

## Child topics

- [Example: Get operation with WSGet](4725-get-resource-data-with-wsget.md): Example of methods you can use to get data from a resource with the WSGet attribute.
- [Example: create resource with WSPost](4726-create-a-resource-with-wspost.md): Create a new resource with the WSPost attribute.
- [Example: Update resource with WSPut](4727-update-a-resource-with-wsput.md): Update a resource with the WSPut attribute.
- [Example: Delete a resource with WSDelete](4728-delete-a-resource-with-wsdelete.md): Delete a resource with the WSDelete attribute.
