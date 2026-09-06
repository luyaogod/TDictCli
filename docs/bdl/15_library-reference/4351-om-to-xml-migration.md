---
title: "OM to XML Migration"
source: "fgl-topics/c_gws_OM_to_XML_001.html"
breadcrumb: "Library reference > Extension packages > The xml package > OM to XML Migration"
type: "concept"
---

# OM to XML Migration

> Code using the OM classes can be converted to XML classes in most cases.

The build-in `om` package provides basic XML handling. The Web Services extensions
`xml` package provides full support for XML document processing. You may need to
change code from `om` to `xml` classes. Before you convert code, make
sure that you are using the classes from the appropriate package:

- Classes from the `om` package exist to manipulate the AUI tree and also provide
  basic XML handling.
- Classes from the `xml` package provides classes and methods to handle any kind of
  XML document processing, and are recommended for XML document manipulation.

Why would you migrate from `om` to `xml` classes and methods?

- You need to be able to utilize a feature (such as a StyleSheet) that requires use of methods
  from the `xml` library classes.

The DOM API of the [om](3280-the-om-package.md "These topics cover the built-in classes of the om package") package
is designed to handle specific FGL files and to manipulate the user interface tree (the AUI
tree). For all other cases, we recommend that you use the DOM API of the Web Services [xml](3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.") package.

## Child topics

- [OM - XML Mapping](4352-om-xml-mapping.md): A reference guide to the DOM APIs of the om and xml classes.
