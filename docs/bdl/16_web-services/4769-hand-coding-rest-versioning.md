---
title: "Hand-coding REST versioning"
source: "fgl-topics/c_gws_restful_high_level_versioning_by_hand.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Hand-coding REST versioning"
type: "concept"
---

# Hand-coding REST versioning

> Hand-coding versioning is the alternative to using the advanced REST API versioning. You must hand-code for each REST operation.

There are two options for hand-coding versioning.

- Set the `WSPath` attribute to specify the version in the URI. See [Versioning with URI](4770-versioning-with-uri.md "You can version a resource using the WSPath attribute. The version is included in the resource URI.").
- Add a custom HTTP header, such as "`api-version`", with the version. See [Versioning with custom header](4771-versioning-with-custom-header.md "You can set version using a custom header.").

There are advantages and disadvantage in using either method and the choice depends on your
preference.

## Related links

**Related concepts**  

[Advanced REST API versioning](4761-advanced-rest-api-versioning.md "Advanced REST API versioning lets you manage multiple versions of a REST web service with minimal coding.")

## Child topics

- [Versioning with URI](4770-versioning-with-uri.md): You can version a resource using the WSPath attribute. The version is included in the resource URI.
- [Versioning with custom header](4771-versioning-with-custom-header.md): You can set version using a custom header.
