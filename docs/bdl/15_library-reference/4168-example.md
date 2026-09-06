---
title: "Example"
source: "fgl-topics/c_gws_XmlStaxReader_examples.html"
breadcrumb: "Library reference > Extension packages > The xml package > The streaming API for XML (StAX) classes > The StaxReader class > Example"
type: "concept"
---

# Example

> Example using methods of the xml.StaxReader class.

```
IMPORT  xml

FUNCTION  parse(file)
  DEFINE  file     STRING
  DEFINE  event    STRING
  DEFINE  ind      INTEGER
  DEFINE  reader   xml.StaxReader
  TRY
    LET  reader=xml.StaxReader.Create()
    CALL  reader.readFrom(file)
    WHILE  (true)
      LET  event=reader.getEventType()
      CASE  event
        WHEN  "START_DOCUMENT"
          DISPLAY  "Document reading started"
          DISPLAY  "XML Version  : ",reader.getVersion()
          DISPLAY  "XML Encoding : ",reader.getEncoding()
          IF  reader.standaloneSet() THEN
            IF  reader.isStandalone() THEN
              DISPLAY  "Standalone   : yes"
            ELSE
              DISPLAY  "Standalone   : no"
            END IF
          END IF
        WHEN  "END_DOCUMENT"
          DISPLAY  "Document reading finished"
        WHEN  "START_ELEMENT"
          IF  reader.isEmptyElement() THEN
            DISPLAY  "<"||reader.getName()||"/>"
          ELSE
            DISPLAY  "<"||reader.getName()||">"
          END IF
          FOR  ind=1 TO  reader.getNamespaceCount()
            DISPLAY  "xmlns:"||reader.getNamespacePrefix(ind)||"="
                        ||reader.getNamespaceURI(ind)
          END FOR
          FOR  ind=1 TO reader.getAttributeCount()
            IF  reader.getAttributePrefix(ind) THEN
              DISPLAY  reader.getAttributePrefix(ind)||":"
                          ||reader.getAttributeLocalName(ind)||"="
                          ||reader.getAttributeValue(ind)
            ELSE
              DISPLAY  reader.getAttributeLocalName(ind)||"="
                          ||reader.getAttributeValue(ind)
            END IF
          END FOR
        WHEN  "END_ELEMENT"
          DISPLAY  "</"||reader.getName()||">"
        WHEN  "CHARACTERS"
          IF  reader.hasText() AND NOT reader.isIgnorableWhitespace() THEN
            DISPLAY  "CHARACTERS :",reader.getText()
          END IF
        WHEN  "COMMENT"
          IF  reader.hasText() THEN
            DISPLAY  "Comment :",reader.getText()
          END IF
        WHEN  "CDATA"
          IF  reader.hasText() THEN
            DISPLAY  "CDATA :", reader.getText()
          END IF
        WHEN  "PROCESSING_INSTRUCTION"
          DISPLAY  "PI :",reader.getPITarget(),reader.getPIData()
        WHEN  "ENTITY_REFERENCE"
          DISPLAY  "Entity name :",reader.getName()
        OTHERWISE
          DISPLAY  "Unknown "||event||" node"
      END CASE
      IF  reader.hasNext() THEN
        CALL  reader.next()
      ELSE
        CALL  reader.close()
        EXIT WHILE
      END IF
    END WHILE
  CATCH
    DISPLAY  "StaxReader ERROR :",status||" ("||sqlca.sqlerrm||")"    
  END TRY
END FUNCTION
```
