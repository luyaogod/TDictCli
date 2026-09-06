---
title: "DOM and SAX built-in classes"
source: "fgl-topics/c_fgl_xml_utils_004.html"
breadcrumb: "Advanced features > XML support > DOM and SAX built-in classes"
type: "concept"
---

# DOM and SAX built-in classes

> The DOM and SAX APIs both contain a set of built-in classes.

The DOM API is composed of:

- The [`om.DomDocument` class](../15_library-reference/3281-the-domdocument-class.md "The om.DomDocument class provides methods to manipulate a data tree, following the DOM standards."),
  that defines the interface to a DOM document. Instances of this class can be used to identify
  and manipulate an XML tree. *DomNode* object manipulation methods are provided by this
  class.
- The [`om.DomNode` class](../15_library-reference/3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree."), that
  defines the interface to an DOM node. Instances of this class can be used to identify and
  manipulate a branch of an XML tree. Child nodes and node attributes management methods are
  provided by this class.
- The [`om.NodeList` class](../15_library-reference/3330-the-nodelist-class.md "A om.NodeList object hold a list of DOM nodes."), to
  handle a list of DomNode objects.

The SAX API is composed of:

- The [`om.SaxAttributes`
  class](../15_library-reference/3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.") represents a set of element attributes. It is used with an
  `om.XmlReader` or an `om.XmlWriter` object.
- The [`om.XmlReader` class](../15_library-reference/3367-the-xmlreader-class.md "The om.XmlReader class provides methods to read and process a file written in XML format."),
  that is defined to read XML. The XML document processing is based on SAX events.
- The [`om.XmlWriter` class](../15_library-reference/3377-the-xmlwriter-class.md "The om.XmlWriter class implements methods to write XML to a stream."),
  that is defined to write XML. The XML document processing is based on SAX events.
- The [`om.SaxDocumentHandler`
  class](../15_library-reference/3352-the-saxdocumenthandler-class.md "The om.SaxDocumentHandler class provides an interface to write an XML filter with events."), which provides an interface to implement a SAX driver using [functions](../08_language-basics/0761-functions.md "Describes user defined functions.") defined in a .4gl
  module loaded dynamically.
