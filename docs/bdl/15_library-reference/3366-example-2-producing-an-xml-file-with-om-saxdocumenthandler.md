---
title: "Example 2: Producing an XML file with om.SaxDocumentHandler"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_example_2.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > Examples > Example 2: Producing an XML file with om.SaxDocumentHandler"
type: "concept"
description: "This example shows how to write an XML file by using a SAX document handler object created from the om.XmlWriter.createFileWriter() method: MAIN DEFINE w om.SaxDocumentHandler DEFINE a ..."
---

# Example 2: Producing an XML file with om.SaxDocumentHandler

This example shows how to write an XML file by using a SAX document handler object created from
the [`om.XmlWriter.createFileWriter()`](3380-om-xmlwriter-createfilewriter.md "Creates an om.SaxDocumentHandler object writing to a file.") method:

```
MAIN
    DEFINE w om.SaxDocumentHandler
    DEFINE a om.SaxAttributes

    LET a = om.SaxAttributes.create()

    LET w = om.XmlWriter.createFileWriter("output.xml")
    CALL w.startDocument()
    CALL w.startElement("Foo", NULL)
    CALL w.processingInstruction("PI2", "val")

    CALL w.startElement("Foo1", NULL)
    CALL w.characters("foo bar")
    CALL w.skippedEntity("nbsp")
    CALL w.EndElement("Foo1")

    CALL w.endElement("Foo")
    CALL w.endDocument()

END MAIN
```
