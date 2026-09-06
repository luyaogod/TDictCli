---
title: "Provide service information"
source: "fgl-topics/c_gws_restful_high_level_service_information.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Provide service information"
type: "concept"
---

# Provide service information

> Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.

In your service module:

1. Define a `PUBLIC` record with a [`WSInfo`](4804-wsinfo.md "Defines general metadata about the REST service.")
   attribute as the service information record.

   If you do not set a service information
   record, the name of the service module is used for the service title.

   The service
   information record allows you to provide metadata about your REST service, which can be
   generated in the openapi.json documentation according to the [OpenAPI standard](http://swagger.io/specification/#infoObject) and used by clients. This example record follows that OpenAPI specification:

   ```
   # Web service server module

   PUBLIC DEFINE serviceInfo
      RECORD ATTRIBUTES(WSInfo, WSScope = "myusers",
                       WSVersion = "myv2", WSVersionMode = "uri")
         title STRING,
         description STRING,
         termOfService STRING,
         contact RECORD
           name STRING,
           url STRING,
           email STRING
         END RECORD,
         version STRING
      END RECORD = ( title: "my service", version: "1.0",
                     contact: ( email:"helpdesk@mysite.com" ) )

   # ... service functions ...
   ```
2. Add record fields as required. The GWS does not check the record structure used with
   `WSInfo` attribute. You are therefore allowed to add the fields you want to be
   available in the OpenAPI documentation.
3. Add values for the record fields.

## Related links

**Related concepts**  

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")

[Publish REST service module](4756-rest-service-with-one-module.md "Publish a service with one resource.")

[Version a REST web service API](4759-version-a-rest-service.md "Versioning your REST web service is important, especially when changes to the service would impact existing clients.")

[WSScope (module)](4807-wsscope-module.md "Defines the access scope required to call any REST operation in the module.")

[WSVersion (module)](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.")

[WSVersionMode](4809-wsversionmode.md "Specifies how the service version is selected for OpenAPI generation (URI, header, or query).")
