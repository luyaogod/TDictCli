---
title: "Publish a REST service"
source: "fgl-topics/c_gws_restful_high_level_publish_service.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Publish a REST service"
type: "concept"
---

# Publish a REST service

> In publishing the service you provide your service to users who can access it on the net.

You must define functions of the service you want to publish as `PUBLIC`. The
fglrestful command only generates specifications for those functions in the
OpenAPI documentation. Functions that are private, are not available as operations of the service
and are therefore not generated.

How you register your service with the GWS REST Web service engine will depend on the number of
resources (modules) you are providing for the service:

- If your service has just one module, you call on the `com.WebServiceEngine.RegisterRestService()` method to register the service. See the
  example described in [Publish REST service module](4756-rest-service-with-one-module.md "Publish a service with one resource.").
- If your service has different resources divided into modules, you call on the `com.WebServiceEngine.RegisterRestResources` method to register the
  resources and the service. See the example described in [Publish REST service with multiple resources](4757-rest-service-with-several-modules.md "Publish a service with several resources.").

## Related links

**Related concepts**  

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")

[REST resource URI naming practice](4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.")

## Child topics

- [Publish REST service module](4756-rest-service-with-one-module.md): Publish a service with one resource.
- [Publish REST service with multiple resources](4757-rest-service-with-several-modules.md): Publish a service with several resources.
