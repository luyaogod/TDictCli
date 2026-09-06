---
title: "Version a REST web service API"
source: "fgl-topics/c_gws_high_level_rest_versioning.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service"
type: "concept"
---

# Version a REST web service API

> Versioning your REST web service is important, especially when changes to the service would impact existing clients.

Changes to a web service API can impact existing clients. For example, a change in the format of the response data or
type (an integer to a float) could result in an error for clients. To ensure that clients are not disrupted, manage the
change by versioning your web service operations.

You version an operation if changes to an existing function, such as the request and/or response data, would break
existing use of the operation. You need to create a new version of the function to apply the changes.

The REST specifications do not provide specific versioning guidelines; however, this section describes commonly used
approaches.

You can either implement versioning using advanced REST API versioning (recommended) or you can hand-code
versioning.

## Child topics

- [How to name versions](4760-how-to-name-versions.md): There are no absolute rules to naming versions, but there are best practices.
- [Advanced REST API versioning](4761-advanced-rest-api-versioning.md): Advanced REST API versioning lets you manage multiple versions of a REST web service with minimal coding.
- [Hand-coding REST versioning](4769-hand-coding-rest-versioning.md): Hand-coding versioning is the alternative to using the advanced REST API versioning. You must hand-code for each REST operation.
