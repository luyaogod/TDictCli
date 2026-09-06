---
title: "Advanced REST API versioning"
source: "fgl-topics/c_gws_high_level_rest_advanced_api_versioning.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Advanced REST API versioning"
type: "concept"
---

# Advanced REST API versioning

> Advanced REST API versioning lets you manage multiple versions of a REST web service with minimal coding.

You enable versioning by setting the `WSVersion` attribute on individual functions
or at the module level in the [service
information record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation."). You choose how clients specify the version using the [WSVersionMode](4809-wsversionmode.md "Specifies how the service version is selected for OpenAPI generation (URI, header, or query).") attribute.

## Versioning an operation

Assign a version name to an operation using the [WSVersion (function)](4825-wsversion-function.md "Specifies the version or versions in which the REST function is available.")
attribute (for example, `WSVersion="v2"`). This allows multiple versions of the same
path and HTTP verb to coexist.

Example:

```
PUBLIC FUNCTION prices_v2(nb INTEGER ATTRIBUTES(WSParam)) 
  ATTRIBUTES (WSGet,WSPath="/prices/{nb}", WSVersion="v2")
  RETURNS (INTEGER)
  # function code ... 
END FUNCTION
```

For details, go to [Create a new version of an operation](4766-create-a-new-version-of-an-operation.md "Specify the version in which the operation is available by setting the WSVersion attribute.")

## Versioning a service

You can assign a version to the entire service by setting `WSVersion` in the [information record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation."). All operations
inherit this version unless they define their own `WSVersion`.

Service-level versioning is useful when most operations change together or when introducing a
major version.

For details, go to [Service-level
versioning](4762-service-level-versioning.md "Overview of settings that control versioning for an entire REST web service.").

## Default service version for OpenAPI

You can define a default version for the OpenAPI documentation. Set the version name in the [WSVersion (module)](4808-wsversion-module.md "Sets the version of the REST service for OpenAPI documentation.")
attribute of the [information
record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation.").

```
# information record
PUBLIC DEFINE myInfo RECORD ATTRIBUTES(WSInfo, WSVersion = "v3")
   title STRING
   # ...
END RECORD
```

For details, go to [Set the default service version for OpenAPI documentation](4763-set-the-default-service-version-for-openapi-documentation.md "Define which version of a service is shown by default in the OpenAPI documentation.").

## Versioning mode

Choose how clients specify a version using the `WSVersionMode` attribute in the
[information record](4753-provide-service-information.md "Provide information about the service, such as title, version, contact details, etc., that is generated in the OpenAPI documentation."). Supported
methods include using the URI, a custom header, or a query string parameter.

```
# information record
PUBLIC DEFINE myInfo RECORD ATTRIBUTES(WSInfo, WSVersion="v3", WSVersionMode="query")
  title STRING
  # ...
END RECORD
```

For details, go to [Version access modes](4764-version-access-modes.md "Specify how clients access service versions in REST operations.").

Read the topics in this section for details about the versioning features provided by the Genero
Web Services, including operation-level versioning, service-level versioning, default versions, and
version-access modes.

If you’re new to versioning, the following topics are a good starting point:

- [Create a new version of an
  operation](4766-create-a-new-version-of-an-operation.md "Specify the version in which the operation is available by setting the WSVersion attribute.")
- [Set a default version of an
  operation](4768-set-the-default-version-for-an-operation.md "Define which version of an operation is shown by default in the OpenAPI documentation.")
- [Service-level
  versioning](4762-service-level-versioning.md "Overview of settings that control versioning for an entire REST web service.")

## Child topics

- [Service-level versioning](4762-service-level-versioning.md): Overview of settings that control versioning for an entire REST web service.
- [Operation-level versioning](4765-operation-level-versioning.md): Overview of settings that control versioning for individual REST operations.
