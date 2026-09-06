---
title: "The DomNodeList class"
source: "fgl-topics/c_gws_XmlDomNodeList.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNodeList class"
type: "concept"
---

# The DomNodeList class

> The xml.DomNodeList class provides methods to manipulate a list of DomNode objects.

This class is provided in the `xml` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`xml` package with:

```
IMPORT xml
```

You can create a DomNodeList object using selection methods in
the [DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.") and [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") classes.
The relationship between the DomNode objects in the list depends on
the method used to create the DomNodeList object.

The `status` variable is set to zero after a successful method call.

## Child topics

- [xml.DomNodeList methods](4086-xml-domnodelist-methods.md): Methods for the xml.DomNodeList class.
