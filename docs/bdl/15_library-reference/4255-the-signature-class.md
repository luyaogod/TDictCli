---
title: "The Signature class"
source: "fgl-topics/c_gws_XmlSignature.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class"
type: "concept"
---

# The Signature class

> The xml.Signature class provides methods to create detached, enveloped or enveloping XML signatures of one or more references of XML documents or document fragments, and to determine whether a signed referenced document has been modified afterwards.

This class is provided in the `xml` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`xml` package with:

```
IMPORT xml
```

It follows the [XML-Signature](http://www.w3.org/TR/xmldsig-core/) specifications.

The `status` variable is set to zero after a successful method call.

## Child topics

- [xml.Signature methods](4256-signature-methods.md): Methods for the xml.Signature class.
- [XML Signature concepts](4287-xml-signature-concepts.md): The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.
- [Digest identifier](4295-digest-identifier.md)
- [Transformation identifier](4296-transformation-identifier.md)
- [Examples](4297-examples.md): xml.Signature usage examples.
