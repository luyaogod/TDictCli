---
title: "Example 1: Writing XML to a file"
source: "fgl-topics/c_fgl_ClassXMLWriter_example_1.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlWriter class > Examples > Example 1: Writing XML to a file"
type: "concept"
description: "MAIN DEFINE w om.SaxDocumentHandler DEFINE a,n om.SaxAttributes LET w = om.XmlWriter.createFileWriter(\"sample.html\") LET a = om.SaxAttributes.create() LET n = om.SaxAttributes.create() CALL n.clear() ..."
---

# Example 1: Writing XML to a file

```
MAIN
  DEFINE w om.SaxDocumentHandler 
  DEFINE a,n om.SaxAttributes 

  LET w = om.XmlWriter.createFileWriter("sample.html")
  LET a = om.SaxAttributes.create()
  LET n = om.SaxAttributes.create()

  CALL n.clear()

  CALL w.startDocument()

    CALL w.startElement("HTML",n)

      CALL w.startElement("HEAD",n)

        CALL w.startElement("TITLE",n)
        CALL w.characters("HTML page generated with XmlWriter")
        CALL w.endElement("TITLE")

        CALL a.clear()
        CALL a.addAttribute("type","text/css")
        CALL w.startElement("STYLE",a)
        CALL w.characters("\nBODY { background-color:#c0c0c0; }\n")  
        CALL w.endElement("STYLE")

      CALL w.endElement("HEAD")

      CALL w.startElement("BODY",n)

        CALL addHLine(w)
        CALL addTitle(w,"What is XML?",1,"55ff55")
        CALL addParagraph(w,"XML = eXtensible Markup Language ...")

        CALL addHLine(w)
        CALL addTitle(w,"What is SAX?",1,"55ff55")
        CALL addParagraph(w,"SAX = Simple Api for XML ...")

      CALL w.endElement("BODY")

    CALL w.endElement("HTML")

  CALL w.endDocument()

END MAIN

FUNCTION addHLine(w)
  DEFINE w om.SaxDocumentHandler 
  DEFINE a om.SaxAttributes 
  LET a = om.SaxAttributes.create()
  CALL a.clear()
  CALL a.addAttribute("width","100%")
  CALL w.startElement("HR",a)
  CALL w.endElement("HR")
END FUNCTION

FUNCTION addTitle(w,t,x,c)
  DEFINE w om.SaxDocumentHandler 
  DEFINE t VARCHAR(100)
  DEFINE x INTEGER DEFINE c VARCHAR(20)
  DEFINE a om.SaxAttributes 
  DEFINE n varchar(10)
  LET a = om.SaxAttributes.create()
  LET n = "h" || x 
  CALL a.clear()
  CALL w.startElement(n,a)
  IF c IS NOT NULL THEN CALL a.addAttribute("color",c)
  END IF CALL w.startElement("FONT",a)
  CALL w.characters(t)
  CALL w.endElement("FONT")
  CALL w.endElement(n)
END FUNCTION

FUNCTION addParagraph(w,t)
  DEFINE w om.SaxDocumentHandler 
  DEFINE t VARCHAR(2000)
  DEFINE a om.SaxAttributes 
  LET a = om.SaxAttributes.create()
  CALL a.clear()
  CALL w.startElement("P",a)
  CALL w.characters("Text is:")
  CALL w.skippedEntity("nbsp") # Add a non breaking space: &nbsp;
  CALL w.characters("is")
  CALL w.characters(t)
  CALL w.endElement("P")
END FUNCTION
```
