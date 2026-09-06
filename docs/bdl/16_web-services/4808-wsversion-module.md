---
title: "WSVersion (module)"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSVersion_module.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSVersion (module)"
type: "concept"
---

# WSVersion (module)

> Sets the version of the REST service for OpenAPI documentation.

## Syntax

```
WSVersion = "{ version }"
```

Where:

1. version is a string value setting the version of the Web service.

`WSVersion` is an optional attribute.

## Usage

You use this attribute to specify version at the module level. It activates versioning for the
web service. When the GWS detects the `WSVersion`
attribute, it automatically switches to versioning mode. This affects how it generates the
OpenAPI documentation and ultimately how resources are accessed.

Set a version name on the `WSVersion` attribute in the [service information record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.") of the
module. By doing this you specify the version of operations of the web service that are displayed in
the OpenAPI document and in the generated stub file.

> **Warning:**
>
> When you use `WSVersion` at the module level, you can only assign a
> single version, and the value cannot be "default". Multiple versions or using
> `WSVersion="default"` are only allowed at the function level. If either is detected
> at the module level, the compiler reports error [error-9140](../15_library-reference/4483-genero-bdl-errors.md).

## OpenAPI

When a default version is defined for a service, you do not need to specify the version in
the query string when requesting the OpenAPI document. For example, the following request returns
the OpenAPI documentation for the default version of the service:
http://myServer:myPort/gas/ws/r/myGroup/myXcf/myService?openapi.json
> **Important:**
>
> Setting a default version is optional. However, if no default version is defined and the client
> does not specify a version in the query string, the request returns a `404 (Not
> found)` and the server raises [error
> 41](../15_library-reference/3805-error-codes-of-com-webservicesengine.md) internally.

When choosing a default version for the OpenAPI document,
make sure the version represents a complete set of operations for the service. The default version
can include:

- operations whose WSVersion value matches the version name you set at the module level, and
- operations that are unversioned or use `WSVersion="default"`.

You can also choose a version name that is not used by any operation. In this case, the OpenAPI
document only includes operations without WSVersion or with `WSVersion="default"`.
This is useful when introducing versioning in the first release of a service; for example, setting
`WSVersion="v1"` even though none of the operations are versioned yet. It helps users
understand that versioning is supported and prepares them for future changes.

To learn more setting a default version of operations, see [Set the default service version for OpenAPI documentation](4763-set-the-default-service-version-for-openapi-documentation.md "Define which version of a service is shown by default in the OpenAPI documentation.").

## Example using WSVersion in the information record

In the example, the `WSVersion` is set to "v3". This makes v3 the default version
used when generating the OpenAPI documentation for the web service.

The main purpose of the service information record (defined with the
`WSInfo` attribute) is to document the REST service. If you are not setting
the version at the modular level here, the record is still needed to provide service
information.

```
PUBLIC DEFINE serviceInfo RECORD ATTRIBUTES(WSInfo, WSScope="users", WSVersion="v3")
  title STRING,
  description STRING,
  termOfService STRING,
  contact RECORD
    name STRING,
    url STRING,
    email STRING
    END RECORD,
    version STRING
  END RECORD = (
    title: "my service", 
    version: "1.0", 
    contact: ( email:"helpdesk@mysite.com") )
```

## Related links

**Related concepts**  

[WSVersionMode](4809-wsversionmode.md "Specifies how the service version is selected for OpenAPI generation (URI, header, or query).")

[Version a REST web service API](4759-version-a-rest-service.md "Versioning your REST web service is important, especially when changes to the service would impact existing clients.")

[How to name versions](4760-how-to-name-versions.md "There are no absolute rules to naming versions, but there are best practices.")

[Understanding the OpenAPI description of a Genero RESTful service](4774-understanding-the-openapi-description-of-a-genero-restful-se.md "The OpenAPI description provides a structured view of your RESTful web service, including service information, paths, parameters, and the request and response bodies generated from your Genero BDL code.")
