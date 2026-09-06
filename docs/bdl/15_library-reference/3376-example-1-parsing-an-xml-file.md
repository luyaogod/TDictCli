---
title: "Example 1: Parsing an XML file"
source: "fgl-topics/c_fgl_ClassXMLReader_example_1.html"
breadcrumb: "Library reference > Built-in packages > The om package > The XmlReader class > Examples > Example 1: Parsing an XML file"
type: "concept"
description: "MAIN DEFINE i, l INTEGER DEFINE r om.XmlReader DEFINE e String DEFINE a om.SaxAttributes LET r = om.XmlReader.createFileReader(\"myfile.xml\") LET l = 0 LET e = r.read() WHILE e IS NOT NULL CASE e WHEN ..."
---

# Example 1: Parsing an XML file

```
MAIN
   DEFINE i, l INTEGER
   DEFINE r om.XmlReader 
   DEFINE e String 
   DEFINE a om.SaxAttributes 
   LET r = om.XmlReader.createFileReader("myfile.xml")
   LET l = 0
   LET e = r.read()
   WHILE e IS NOT NULL
     CASE e 
       WHEN "StartDocument"
         DISPLAY "StartDocument:"
       WHEN "StartElement"
         LET l=l+1
         DISPLAY l SPACES, "StartElement:", r.getTagName()
         LET a = r.getAttributes()
         FOR i=1 to a.getLength()
           DISPLAY l SPACES,"  ",
             a.getName(i)," = ",
             a.getValueByIndex(i)
         END FOR
       WHEN "Characters"
         DISPLAY l SPACES, "  Characters:'",r.getCharacters(),"'"
       WHEN "SkippedEntity"
         DISPLAY "Entity:'",r.skippedEntity(),"'"
       WHEN "EndElement"
         DISPLAY l SPACES, "EndElement:", r.getTagName()
         LET l=l-1
       WHEN "EndDocument"
         DISPLAY "EndDocument:"
       OTHERWISE
         DISPLAY "Invalid event: ",e 
     END CASE
     LET e=r.read()
   END WHILE
END MAIN
```
