---
title: "The Encryption class"
source: "fgl-topics/c_gws_XmlEncryption.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class"
type: "concept"
---

# The Encryption class

> The xml.Encryption class provides methods to encrypt and decrypt XML documents, nodes or symmetric keys.

This class is provided in the `xml` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`xml` package with:

```
IMPORT xml
```

It follows the [XML-Encryption](http://www.w3.org/TR/xmlenc-core/) specifications.

The `status` variable is set to zero after a successful method call.

## Child topics

- [xml.Encryption methods](4310-encryption-methods.md): Methods for the xml.Encryption class.
- [Examples](4330-examples.md): xml.Encryption usage examples.
