---
title: "Setting MIME type at runtime"
source: "fgl-topics/c_gws_restful_high_level_set_mime_type_runtime.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Setting MIME type at runtime"
type: "concept"
---

# Setting MIME type at runtime

> Override the GWS default media format for messages.

If the MIME type is not set by the `WSMedia` attribute, you can specify the
format at runtime by calling the REST service engine option with either JSON or XML, or both. For
example:

```
com.WebServiceEngine.SetOption("server_restdefaultformat","xml")
```

```
com.WebServiceEngine.SetOption("server_restdefaultformat","both")
```

## Related links

**Related concepts**  

[Default media types](4846-default-media-types.md "For RESTful Web services developed using the high-level framework, you can specify the MIME type or you can accept the default. The default MIME type is based on the data type.")
