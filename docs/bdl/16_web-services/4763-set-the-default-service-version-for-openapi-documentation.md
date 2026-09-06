---
title: "Set the default service version for OpenAPI documentation"
source: "fgl-topics/c_gws_high_level_rest_default_version_service.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Advanced REST API versioning > Service-level versioning > Set the default service version for OpenAPI documentation"
type: "concept"
---

# Set the default service version for OpenAPI documentation

> Define which version of a service is shown by default in the OpenAPI documentation.

You can define a default version for an entire web service. The default version determines which versions of the operations appear in the OpenAPI documentation when no specific version is requested.

The default version is set in the service [information record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.") using the
[WSVersion (module)](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.") attribute.

This setting applies to the **entire web service**.

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

## Default OpenAPI version

In this example, `WSVersion` is set to "v3". This makes v3 the default version
used when generating the OpenAPI documentation for the web service.
> **Warning:**
>
> When you use `WSVersion` at the module level, you can only assign a
> single version, and the value cannot be "default". Multiple versions or using
> `WSVersion="default"` are only allowed at the function level. If either is detected
> at the module level, the compiler reports error [error-9140](../15_library-reference/4483-genero-bdl-errors.md).

```
PUBLIC DEFINE myInfo RECORD ATTRIBUTES(WSInfo, WSVersion = "v3")
   title STRING
   # ...
END RECORD
```

## Related links

**Related concepts**  

[Set the default version for an operation](4768-set-the-default-version-for-an-operation.md "Define which version of an operation is shown by default in the OpenAPI documentation.")

[Version access modes](4764-version-access-modes.md "Specify how clients access service versions in REST operations.")

[Operation-level versioning](4765-operation-level-versioning.md "Overview of settings that control versioning for individual REST operations.")

**Related tasks**  

[Get the OpenAPI service description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.")

[Get the OpenAPI description for a specific version](4777-get-openapi-description-for-a-version.md "Retrieve the OpenAPI description for a specific version of a RESTful web service.")

[Get available versions of a service](4778-get-available-versions-of-a-service.md "Retrieve the list of available versions of a REST service and identify the default version if defined.")
