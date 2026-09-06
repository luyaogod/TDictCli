---
title: "REST resource URI naming practice"
source: "fgl-topics/c_gws_restful_high_level_uri_convention.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Resource URI naming"
type: "concept"
---

# REST resource URI naming practice

> The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.

RESTful Web services rely on providing the URL to access the resource. The URI you use in
development mode will be different to what you use in your production environment.

You can have one or several resource modules, depending on how you want to organize your
resources and [publish your Web
service](4755-publish-a-rest-service.md "In publishing the service you provide your service to users who can access it on the net."). This affects how the path for endpoints are created.

## Resource URI syntax

In a production environment, the Genero Application Server (GAS) provides the base URL from the
server where the Web service is deployed. This follows the format:

```
http[s]://host:port/gas/ws/r[/group-name]/xcf-file/endpoint-path[?query-params]
```

where endpoint-path contains several path segments:

```
service-name[/resource-name]/resource-path
```

1. host:port/gas/ws/r is
   the base URL from the server. For more information, see the Application Web Address
   page in the Genero Application Server User Guide.
2. group-name is the name of the group. If the Web service
   belongs to the default group on the GAS, this is optional. If the web service is part of a group
   that is not the default, you must provide the group name. For more information, see the Genero Application Server User Guide.
3. xcf-file is the name of the web service configuration file
   on the GAS where the web service is deployed. For example, if the configuration file is
   myService.xcf, you specify myService in the
   xcf-file path segment. For more information, see the Genero Application Server User Guide.
4. service-name is the name of the web service.
5. resource-name is the name of the resource. The
   resource-name is part of the path when you have more than one resource registered
   in the GWS. It is **not** part of the path when you only have one resource registered in the GWS.
   Use the appropriate function to register your REST Web services:
   - [com.WebServiceEngine.RegisterRestService](../15_library-reference/3793-com-webserviceengine-registerrestservice.md "Registers a REST service in the engine.") to register a service with one
     resource.
   - [com.WebServiceEngine.RegisterRestResources](../15_library-reference/3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine.") to register a service with
     many resources
6. resource-path is the path to the resource, as defined
   by [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL."). This is a logical
   (not a physical) URI that determines the operation you are requesting.
7. query-params is a set of URL query parameters.

## URI syntax for the OpenAPI documentation of a Web service

If you publish a web service with a single resource module or many resource modules, the URL for
the service's OpenAPI documentation is:

```
http[s]://host:port/gas/ws/r[/group-name]/xcf-file/service-name?openapi.json
```

1. host:port/gas/ws/r is
   the base URL from the server. For more information, see the Application Web Address
   page in the Genero Application Server User Guide.
2. group-name is the name of the group. If the Web service
   belongs to the default group on the GAS, this is optional. If the web service is part of a group
   that is not the default, you must provide the group name. For more information, see the Genero Application Server User Guide.
3. xcf-file is the name of the web service configuration file
   on the GAS where the web service is deployed. For example, if the configuration file is
   myService.xcf, you specify myService in the
   xcf-file path segment. For more information, see the Genero Application Server User Guide.
4. service-name is the name of the web service.

For example:

http://myhost/gas/ws/r/myGroup/myXcf/accounts?openapi.json

## URI syntax for the REST API catalog of a Web service

The RFC 9727 REST API catalog for the web service configuration (FGLGWS 6.00.03 or later) is
available at:

```
http[s]://host:port/gas/ws/r[/group-name]/xcf-file/.well-known/api-catalog
```

1. host:port/gas/ws/r is
   the base URL from the server. For more information, see the Application Web Address
   page in the Genero Application Server User Guide.
2. group-name is the name of the group. If the Web service
   belongs to the default group on the GAS, this is optional. If the web service is part of a group
   that is not the default, you must provide the group name. For more information, see the Genero Application Server User Guide.
3. xcf-file is the name of the web service configuration file
   on the GAS where the web service is deployed. For example, if the configuration file is
   myService.xcf, you specify myService in the
   xcf-file path segment. For more information, see the Genero Application Server User Guide.

For example:

http://myhost/gas/ws/r/myGroup/myXcf/.well-known/api-catalog

For the response format, see [REST API
catalog](4758-rest-api-catalog.md "A Genero application that publishes REST services automatically provides an RFC 9727 API catalog at /.well-known/api-catalog, listing each service and its OpenAPI description.").

## URI for Web service with one resource or for RFC

If you have only one resource or when you want to implement a Remote Function Call (RFC) service,
you register the service on the GWS with [com.WebServiceEngine.RegisterRestService](../15_library-reference/3793-com-webserviceengine-registerrestservice.md "Registers a REST service in the engine."). This adds the service name to the path.
The URI specification is:

```
http[s]://host:port/gas/ws/r[/group-name]/xcf-file/service-name/resource-path
```

1. host:port/gas/ws/r is
   the base URL from the server. For more information, see the Application Web Address
   page in the Genero Application Server User Guide.
2. group-name is the name of the group. If the Web service
   belongs to the default group on the GAS, this is optional. If the web service is part of a group
   that is not the default, you must provide the group name. For more information, see the Genero Application Server User Guide.
3. xcf-file is the name of the web service configuration file
   on the GAS where the web service is deployed. For example, if the configuration file is
   myService.xcf, you specify myService in the
   xcf-file path segment. For more information, see the Genero Application Server User Guide.
4. service-name is the name of the web service.
5. resource-path is the path to the resource, as defined
   by [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL."). This is a logical
   (not a physical) URI that determines the operation you are requesting.

In this example, you have a Web service that manages user accounts. It has one resource module
(accounts) registered on the GWS with [com.WebServiceEngine.RegisterRestService](../15_library-reference/3793-com-webserviceengine-registerrestservice.md "Registers a REST service in the engine."). It is deployed on a GAS where users can
access its operations. For instance, a URI to access the "users" operation is at:

http://myhost/gas/ws/r/myGroup/myXcf/accounts/users

Where:

- myhost/gas/ws/r is the base URL from the server.
- myGroup is the name of the group on the GAS where the Web service is
  deployed.
- myXcf is the name of the Web service configuration
  .xcf file.
- accounts is the name of the Web service.
- users is a resource-path of a specific resource
  operation set with [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL.").

The service's OpenAPI documentation is at:

http://myhost/gas/ws/r/myGroup/myXcf/accounts?openapi.json

## URI for Web service with many resources

If you have more than one resource, typically, you have one resource per
.4gl module. You register the service on the GWS with [com.WebServiceEngine.RegisterRestResources](../15_library-reference/3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine."). This builds the path for endpoints
with the service name followed by the resource module file name. The URI documentation is:

```
http[s]://host:port/gas/ws/r[/group-name]/xcf-file/service-name/resource-name/resource-path
```

1. host:port/gas/ws/r is
   the base URL from the server. For more information, see the Application Web Address
   page in the Genero Application Server User Guide.
2. group-name is the name of the group. If the Web service
   belongs to the default group on the GAS, this is optional. If the web service is part of a group
   that is not the default, you must provide the group name. For more information, see the Genero Application Server User Guide.
3. xcf-file is the name of the web service configuration file
   on the GAS where the web service is deployed. For example, if the configuration file is
   myService.xcf, you specify myService in the
   xcf-file path segment. For more information, see the Genero Application Server User Guide.
4. service-name is the name of the web service.
5. resource-name is the name of the resource. The
   resource-name is part of the path when you have more than one resource registered
   in the GWS by [com.WebServiceEngine.RegisterRestResources](../15_library-reference/3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine.").
6. resource-path is the path to the resource, as defined
   by [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL."). This is a logical
   (not a physical) URI that determines the operation you are requesting.

In the next example we have a Web service that manages fruits. It has several resources (for
example, "apple", "peach", etc.) registered on the GWS with [com.WebServiceEngine.RegisterRestResources](../15_library-reference/3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine."). It is deployed on a GAS where users
can access its resources. For instance, a URI to access the "apple" resource "prices" operation is
at:

http://myhost/gas/ws/r/myXcf/fruits/apple/prices

Where:

- myhost/gas/ws/r is the base URL from the server.
- For this example, there is no group-name in the URI as the service is
  deployed in the default group of the GAS.
- myXcf is the name of the Web service .xcf
  configuration file.
- fruits is the name of the Web service (service-name)
- apple is the name of the resource (resource-name).
- prices is the resource-path (defined with [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL.")) for the "apple"
  resource-name.

The service's OpenAPI documentation is at:

http://myhost/gas/ws/r/myXcf/fruits?openapi.json

## URI in development

When running the Web service in development, you access resources and the OpenAPI documentation
differently depending on the mode used.

In development HTTPS is not supported. Using the standalone dispatcher on localhost, the URI
specification is:

```
http://host:port/ws/r[/group-name]/xcf-file/service-name/resource-name/resource-path
```

1. host:port/ws/r is the base URL using
   the standalone dispatcher. For more information, see the Application Web Address page
   in the Genero Application Server User Guide.
2. group-name is the name of the group. If the Web service
   belongs to the default group on the GAS, this is optional. If the web service is part of a group
   that is not the default, you must provide the group name. For more information, see the Genero Application Server User Guide.
3. xcf-file is the name of the Web service configuration file on the GAS. For
   example, if the configuration file is myService.xcf, you specify
   myService in the xcf-file path segment.
4. service-name is the name of the web service.
5. resource-name is the name of the resource. The
   resource-name is part of the path when you have more than one resource registered
   in the GWS by [com.WebServiceEngine.RegisterRestResources](../15_library-reference/3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine.").
6. resource-path is the path to the resource, as defined
   by [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL."). This is a logical
   (not a physical) URI that determines the operation you are requesting.

For example, running the Web service for managing fruits at the default port (6394):

- A URI to access the "apple" resource "prices" operation
  is:

  http://localhost:6394/ws/r/myXcf/fruits/apple/prices

  Where:
  - localhost:6394/ws/r is the base URL for the standalone dispatcher.
  - For this example, there is no group-name in the URI as the service is
    deployed in the default group of the GAS.
  - myXcf is the name of the Web service .xcf
    configuration file.
  - fruits is the name of the Web service
    (service-name).
  - apple is the name of the resource (resource-name).
  - prices is the resource-path (defined with [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL.")) for the "apple"
    resource-name.
- The service's OpenAPI documentation is at:

  http://localhost:6394/ws/r/myXcf/fruits?openapi.json

In direct mode the URI specification
is:

```
http://host:port/service-name/resource-name/resource-path
```

1. host:port is the base URL using
   direct mode. For more information, see the Application Web Address page in the Genero Application Server User Guide.
2. service-name is the name of the web service.
3. resource-name is the name of the resource. The
   resource-name is part of the path when you have more than one resource registered
   in the GWS by [com.WebServiceEngine.RegisterRestResources](../15_library-reference/3792-com-webserviceengine-registerrestresources.md "Registers the resources of a REST service in the engine.").
4. resource-path is the path to the resource, as defined
   by [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL."). This is a logical
   (not a physical) URI that determines the operation you are requesting.

For example, running the Web service for managing fruits with FGLAPPSERVER set at port 8090 on
localhost:

- A URI to access the "apple" resource "prices" operation is
  at:

  http://localhost:8090/fruits/apple/prices

  Where:
  - localhost:8090 is the base URL for direct mode.
  - There is no configuration file (xcf-file) because the GAS is not involved in
    direct mode.
  - fruits is the name of the Web service
    (service-name).
  - apple is the name of the resource (resource-name).
  - prices is the resource-path (defined with [WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL.")) for the "apple"
    resource-name.
- The service's OpenAPI documentation is at:

  http://localhost:8090/fruits?openapi.json

## Related links

**Related concepts**  

[Version a REST web service API](4759-version-a-rest-service.md "Versioning your REST web service is important, especially when changes to the service would impact existing clients.")
