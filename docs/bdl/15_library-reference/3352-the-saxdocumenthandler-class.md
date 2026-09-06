---
title: "The SaxDocumentHandler class"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class"
type: "concept"
---

# The SaxDocumentHandler class

> The om.SaxDocumentHandler class provides an interface to write an XML filter with events.

This class follows the [SAX](../09_advanced-features/0949-dom-and-sax-standards.md "DOM and SAX are both programming interfaces that can work with XML.") standards.

A `om.SaxDocumentHandler` object can be used in two different ways:

1. To implement an XML SAX filter, based on functions defined in a .4gl
   module, by using the [`createForName()`](3354-om-saxdocumenthandler-createforname.md "Creates a new SAX document handler object for the given .4gl module.") class method.
2. To write an XML document to a file, process or socket output, by using [`om.XmlWriter`](3377-the-xmlwriter-class.md "The om.XmlWriter class implements methods to write XML to a stream.") creation
   methods, and the `om.SaxDocumentHandler` processing methods.

The `om.SaxDocumentHandler` class also provides methods to process all SAX
events by hand. This is useful if you want to chain SAX filters.

> **Important:**
>
> Build-in XML support classes `om.*` have known limitations. Refere to the [Limitations of XML built-in classes](../09_advanced-features/0951-limitations-of-xml-built-in-classes.md "Built-in XML classes have some limitations you must be aware of.") page for more details.

## Child topics

- [om.SaxDocumentHandler methods](3353-om-saxdocumenthandler-methods.md): Methods of the om.SaxDocumentHandler class.
- [Examples](3364-examples.md): om.SaxDocumentHandler usage examples.
