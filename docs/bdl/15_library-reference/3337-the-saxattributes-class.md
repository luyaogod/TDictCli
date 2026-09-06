---
title: "The SaxAttributes class"
source: "fgl-topics/c_fgl_ClassSaxAttributes.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class"
type: "concept"
---

# The SaxAttributes class

> The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.

To process SAX attributes, create a `om.SaxAttributes` object with a [SAX reader](3367-the-xmlreader-class.md "The om.XmlReader class provides methods to read and process a file written in XML format.") or [SAX
writer](3377-the-xmlwriter-class.md "The om.XmlWriter class implements methods to write XML to a stream.") object.

Get an instance of SaxAttributes with the `om.XmlReader.getAttributes()`
method.

> **Important:**
>
> Build-in XML support classes `om.*` have known limitations. Refere to the [Limitations of XML built-in classes](../09_advanced-features/0951-limitations-of-xml-built-in-classes.md "Built-in XML classes have some limitations you must be aware of.") page for more details.

## Child topics

- [om.SaxAttributes methods](3338-om-saxattributes-methods.md): Methods of the om.SaxAttributes class.
- [Examples](3349-examples.md): om.SaxAttributes usage examples.
