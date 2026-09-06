---
title: "Example"
source: "fgl-topics/c_gws_XmlStaxWriter_examples.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxWriter class > Example"
type: "concept"
---

# Example

> This example uses methods from the xml.StaxWriter class.

```
IMPORT xml

FUNCTION save(file)
  DEFINE file STRING
  DEFINE writer xml.StaxWriter
  TRY
    LET writer = xml.StaxWriter.Create()
    CALL writer.setFeature("format-pretty-print",TRUE)
    CALL writer.writeTo(file)
    CALL writer.startDocument("utf-8","1.0",true)
    CALL writer.comment("This is my first comment using a stax writer")
    CALL writer.setPrefix("c","http://www.mycompany.com/c")
    CALL writer.setPrefix("d","http://www.mycompany.com/d")
    CALL writer.setDefaultNamespace("http://www.mycompany.com/d")
    CALL writer.startElementNS("root", "http://www.mycompany.com/d")
    CALL writer.attribute("attr1","value1")
    CALL writer.attribute("attr2","value2")
    CALL writer.attributeNS("attr3", "http://www.mycompany.com/d","value3")
    CALL writer.comment("This is a comment using a stax writer")
    CALL writer.startElementNS("eltA", "http://www.mycompany.com/d")
    CALL writer.cdata("<this is a CData section>")
    CALL writer.endElement()
    CALL writer.startElementNS("eltB", "http://www.mycompany.com/c")
    CALL writer.characters("Hello world, I'm from the development team")
    CALL writer.entityRef("one")
    CALL writer.endElement()
    CALL writer.processingInstruction("command1","do what you want")
    CALL writer.endElement()
    CALL writer.comment("This is my last comment using a stax writer")
    CALL writer.endDocument()
    CALL writer.close()
    RETURN TRUE
  CATCH
    DISPLAY "StaxWriter ERROR :",status, sqlca.sqlerrm
    RETURN FALSE
  END TRY
END FUNCTION
```
