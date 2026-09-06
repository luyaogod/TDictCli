---
title: "The Serializer class"
source: "fgl-topics/c_gws_XmlSerializer.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class"
type: "concept"
---

# The Serializer class

> The xml.Serializer class provides methods to manage options for the serializer engine, and to use the serializer engine to serialize variables and XML element nodes.

This class is provided in the `xml` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`xml` package with:

```
IMPORT xml
```

This class is a static class and does not have to be instantiated.

The `status` variable is set to zero after a successful method call.

## Child topics

- [xml.Serializer methods](4171-xml-serializer-methods.md): Methods for the xml.Serializer class.
- [Serialization option flags](4189-serialization-option-flags.md): Serialization option flags for the xml.Serializer class.
