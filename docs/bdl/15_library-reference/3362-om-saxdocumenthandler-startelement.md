---
title: "om.SaxDocumentHandler.startElement"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_startElement.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > om.SaxDocumentHandler methods > om.SaxDocumentHandler.startElement"
type: "concept"
---

# om.SaxDocumentHandler.startElement

> Processes the beginning of an element.

## Syntax

```
startElement(
   name STRING,
   atts om.SaxAttributes )
```

1. name is the tag name of the element.
2. atts is the list of attributes of the element.

## Usage

The `startElement()` method processes the beginning of an element with the SAX
interface.

Use the `om.SaxAttributes` methods to handle the attributes of an element.

## Example

```
DEFINE out om.SaxDocumentHandler
       node om.DomNode,
       attrs om.SaxAttributes,
       x, c INTEGER
...
CALL attrs.clear()
LET c = node.getChildCount()
FOR x=1 TO c
    CALL attrs.addAttribute( node.getAttributeName(x),
                             node.getAttributeValue(x) )
END FOR
CALL out.startElement( node.getTagName(), attrs )
```

## Related links

**Related concepts**  

[om.SaxDocumentHandler.endElement](3357-om-saxdocumenthandler-endelement.md "Processes the end of an element.")

[The SaxAttributes class](3337-the-saxattributes-class.md "The om.SaxAttributes class holds a set of attributes to process with a SAX reader or writer.")
