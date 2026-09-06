---
title: "WSInfo"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSInfo.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSInfo"
type: "concept"
---

# WSInfo

> Defines general metadata about the REST service.

## Syntax

```
WSInfo
```

## Usage

You use this attribute to specify information that describes the service such as the service
title, version, and contact details.

The `WSInfo` attribute must be set in an `ATTRIBUTES()`
clause on the information record at the service module level. This allows you to [provide service
information](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.") that is useful to clients.

`WSInfo` is an optional attribute.

## Example using WSInfo in the service information record

The example service information record follows the [OpenAPI
standard](http://swagger.io/specification/#infoObject).

The GWS does not check the record structure used with `WSInfo` attribute.
You are therefore allowed to add the fields you want to be available in the OpenAPI
specification.

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

## Related links

**Related concepts**  

[WSScope (module)](4807-wsscope-module.md "Defines the access scope required to call any REST operation in the module.")

[Version a REST web service API](4759-version-a-rest-service.md "Versioning your REST web service is important, especially when changes to the service would impact existing clients.")

[Provide service information](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.")

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")
