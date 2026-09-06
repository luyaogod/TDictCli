---
title: "Service-level versioning"
source: "fgl-topics/c_gws_high_level_rest_service_level_versioning.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Advanced REST API versioning > Service-level versioning"
type: "concept"
---

# Service-level versioning

> Overview of settings that control versioning for an entire REST web service.

Service-level versioning defines how a REST Web service exposes and manages its API versions.
These settings apply to the web service as a whole and determine how versions are selected, how
clients specify a version, and how versions appear in the OpenAPI documentation.

The features in this section let you control versioning consistently across all operations of the service.
They operate at the service level, independently of any version settings applied to individual operations.

## Child topics

- [Set the default service version for OpenAPI documentation](4763-set-the-default-service-version-for-openapi-documentation.md): Define which version of a service is shown by default in the OpenAPI documentation.
- [Version access modes](4764-version-access-modes.md): Specify how clients access service versions in REST operations.
