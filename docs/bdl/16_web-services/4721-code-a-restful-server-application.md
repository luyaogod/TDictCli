---
title: "Code a RESTful server application (high-level framework)"
source: "fgl-topics/c_gws_restful_high_level_server.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application"
type: "concept"
---

# Code a RESTful server application (high-level framework)

> To create a RESTful Genero Web server application, you need to create a Genero BDL module that defines the service functions. When you publish it, a service description is available when required.

A Web service application demo can be found at
$FGLDIR/demo/WebServices/books.

## Child topics

- [Set up and test your environment](4722-set-up-and-test-your-environment.md)
- [Define functions in a module](4723-define-functions-in-a-module.md): A GWS REST service is defined in a module.
- [Provide service information](4753-provide-service-information.md): Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.
- [REST resource URI naming practice](4754-resource-uri-naming.md): The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.
- [Publish a REST service](4755-publish-a-rest-service.md): In publishing the service you provide your service to users who can access it on the net.
- [REST API catalog](4758-rest-api-catalog.md): A Genero application that publishes REST services automatically provides an RFC 9727 API catalog at /.well-known/api-catalog, listing each service and its OpenAPI description.
- [Version a REST web service API](4759-version-a-rest-service.md): Versioning your REST web service is important, especially when changes to the service would impact existing clients.
- [Setting MIME type at runtime](4772-setting-mime-type-at-runtime.md): Override the GWS default media format for messages.
