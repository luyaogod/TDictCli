---
title: "REST API catalog"
source: "fgl-topics/c_gws_restful_high_level_api_catalog.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > REST API catalog"
type: "concept"
---

# REST API catalog

> A Genero application that publishes REST services automatically provides an RFC 9727 API catalog at /.well-known/api-catalog, listing each service and its OpenAPI description.

You do not need to write any 4GL code to enable the catalog. The GWS engine builds and serves it
automatically from the services registered with
[`com.WebServiceEngine.RegisterRestService()`](../15_library-reference/3793-com-webserviceengine-registerrestservice.md "Registers a REST service in the engine.").
The catalog applies to the high-level REST framework only; services that use the low-level API are
not included. The catalog returns a linkset that contains links to the OpenAPI descriptions for each
REST service published by the application.

The GWS engine intercepts `GET` requests for
`/.well-known/api-catalog` before any user request handler is executed and sets the
`Content-Type` response header to:

```
application/linkset+json; profile="https://www.rfc-editor.org/info/rfc9727"
```

The response body contains a
`linkset` array with one entry for each published REST service. Each entry
contains:

- `anchor`: The absolute base URL of the service, reconstructed from the host,
  port, and path of the incoming request. When deployed on the Genero Application Server, the path
  includes the group-name and xcf-file name, as described in
  [REST resource URI naming
  practice](4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.").
- `service-desc`: Links to the OpenAPI description for the service. Each link has
  a `type` value of `application/json`.

## Service versions

1. A non-versioned service, or a versioned service with a default version, gets a single
   `service-desc` link at `?openapi.json` (the engine resolves it to
   the default version).
2. A versioned service with no default version gets one `service-desc` link per
   version (`?openapi.json&version=v`).

The full list of versions for a service is always available separately at
`service?version.json`. For the full URI format, see [Versioning with URI](4770-versioning-with-uri.md "You can version a resource using the WSPath attribute. The version is included in the resource URI.").

## Example

Request:

```
GET /.well-known/api-catalog HTTP/1.1
Host: myhost:8090
```

Response body for an application publishing a versioned `orders` service with a
default version, and a non-versioned `products` service:

```
{
"linkset": [
{
"anchor": "http://myhost:8090/orders",
"service-desc": [
{ "href": "http://myhost:8090/orders?openapi.json", "type": "application/json" }
]
},
{
"anchor": "http://myhost:8090/products",
"service-desc": [
{ "href": "http://myhost:8090/products?openapi.json", "type": "application/json" }
]
}
]
}
```

A versioned service with no declared default version returns one `service-desc`
link per version:

```
{
"anchor": "http://myhost:8090/reports",
"service-desc": [
{ "href": "http://myhost:8090/reports?openapi.json&version=v1", "type": "application/json" },
{ "href": "http://myhost:8090/reports?openapi.json&version=v2", "type": "application/json" }
]
}
```

An application that publishes no REST services returns an empty catalog:

```
{ "linkset": [] }
```

## Related links

**Related concepts**  

[Publish a REST service](4755-publish-a-rest-service.md "In publishing the service you provide your service to users who can access it on the net.")

**Related tasks**  

[Get the OpenAPI service description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.")
