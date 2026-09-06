---
title: "The xml package"
source: "fgl-topics/c_gws_XML_Library_001.html"
breadcrumb: "Library reference > Extension packages > The xml package"
type: "concept"
---

# The xml package

> The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.

The library provides a W3C-compatible DOM API, integrating additional XML Schema and DTD
validation methods. There is also an API compatible with StAX for writing or reading XML documents
where performance and speed are important.

The DOM API of the [om](3280-the-om-package.md "These topics cover the built-in classes of the om package") package
is designed to handle specific FGL files and to manipulate the user interface tree (the AUI
tree). For all other cases, we recommend that you use the DOM API of the Web Services [xml](3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.") package.

## Child topics

- [The Document Object Modeling (DOM) classes](3958-the-document-object-modeling-dom-classes.md): The Document Object Modeling (DOM) classes manage XML documents entirely in memory with support for XML Schema and DTD validation.
- [The streaming API for XML (StAX) classes](4089-the-streaming-api-for-xml-stax-classes.md): The streaming API for XML (StAX) classes use streaming while managing XML documents.
- [XML serialization classes](4169-xml-serialization-classes.md): The XML serialization classes convert BDL variables to XML and XML to BDL variables.
- [XML security classes](4190-xml-security-classes.md): XML Security classes handle encryption and signature of XML documents entirely in memory with keys and certificates.
- [XML transformation classes](4340-xml-transformation-classes.md): The XML transformation classes transform XML.
- [OM to XML Migration](4351-om-to-xml-migration.md): Code using the OM classes can be converted to XML classes in most cases.
