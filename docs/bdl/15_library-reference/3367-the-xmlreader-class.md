---
title: "The XmlReader class"
source: "fgl-topics/c_fgl_ClassXMLReader.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class"
type: "concept"
---

# The XmlReader class

> The om.XmlReader class provides methods to read and process a file written in XML format.

The processing of the XML file is streamed-data based; the file is loaded and processed
sequentially with events. To process XML element attributes, an `om.XmlReader`
object must cooperate with a [`om.SaxAttributes`](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.") object. The
XmlReader class can only read from a file. To write to a file, use the [`om.XmlWriter`](3377-the-xmlwriter-class.md "The om.XmlWriter class implements methods to write XML to a stream.") class.

Steps to use a XML reader:

1. Declare a variable with the `om.XmlReader` type.
2. Create the reader object with the [`createFileReader()`](3369-om-xmlreader-createfilereader.md "Creates an XML reader object from a file.") method and assign the reference to the
   variable.
3. Process SAX events in a `WHILE` loop, by reading document fragments with
   the [`read()`](3373-om-xmlreader-read.md "Reads the next SAX event to process.")
   method.
4. Inside the loop, depending on the SAX event, process element attributes with [getAttributes()](3370-om-xmlreader-getattributes.md "Builds an attribute list for the current processed element.") or get the element data with
   the [getCharacters()](3371-om-xmlreader-getcharacters.md "Returns the character data of the current processed element.") methods.

> **Important:**
>
> Build-in XML support classes `om.*` have known limitations. Refere to the [Limitations of XML built-in classes](../09_advanced-features/0951-limitations-of-xml-built-in-classes.md "Built-in XML classes have some limitations you must be aware of.") page for more details.

## Child topics

- [om.XmlReader methods](3368-om-xmlreader-methods.md): Methods of the om.XmlReader class.
- [Examples](3375-examples.md): om.XmlReader usage examples.
