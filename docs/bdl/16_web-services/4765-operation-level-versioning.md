---
title: "Operation-level versioning"
source: "fgl-topics/c_gws_high_level_rest_operation_level_versioning.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Advanced REST API versioning > Operation-level versioning"
type: "concept"
---

# Operation-level versioning

> Overview of settings that control versioning for individual REST operations.

Operation-level versioning defines how individual REST operations are associated with
one or more API versions. These settings apply to specific functions within a Web service
and determine how each operation appears in the OpenAPI documentation for a given version.

Version an operation when changes to its request or response would break existing clients.
By introducing a new version of the function and assigning a unique version value, you can
publish updated behavior without affecting existing users who rely on earlier versions.

Operation-level versioning settings operate independently of any versioning applied
at the service level.

## Child topics

- [Create a new version of an operation](4766-create-a-new-version-of-an-operation.md): Specify the version in which the operation is available by setting the WSVersion attribute.
- [Define multiple versions for an operation](4767-define-multiple-versions-for-an-operation.md): Specify which API versions an operation applies to using the WSVersion attribute.
- [Set the default version for an operation](4768-set-the-default-version-for-an-operation.md): Define which version of an operation is shown by default in the OpenAPI documentation.
