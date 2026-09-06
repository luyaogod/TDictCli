---
title: "Versioning with custom header"
source: "fgl-topics/c_gws_restful_high_level_versioning_header.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Version a REST service > Hand-coding REST versioning > Versioning with custom header"
type: "concept"
---

# Versioning with custom header

> You can set version using a custom header.

The version can be included in a custom
header.

```
api-version:v2.0
```

The version must be
specified by the client. Clients need to know which version to specify before requesting a resource.

The API function retrieves the value of the header. If the version matches one specified, then
that is the method that will be invoked. Based on this, the API can decide if it is capable of
fulfilling the request and respond to the client appropriately.

## Related links

**Related concepts**  

[WSInfo](4804-wsinfo.md "Defines general metadata about the REST service.")

[REST resource URI naming practice](4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.")
