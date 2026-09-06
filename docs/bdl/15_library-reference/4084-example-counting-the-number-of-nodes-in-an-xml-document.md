---
title: "Example Counting the number of nodes in an XML document"
source: "fgl-topics/c_gws_XML_DomNode_class_012.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomNode class > Examples > Example Counting the number of nodes in an XML document"
type: "concept"
description: "This code example counts the number of nodes of each type. IMPORT XML DEFINE nbElt INTEGER DEFINE nbAttr INTEGER DEFINE nbComment INTEGER DEFINE nbPI INTEGER DEFINE nbTxt INTEGER DEFINE nbCData ..."
---

# Example Counting the number of nodes in an XML document

This code example counts the number of nodes of each type.

```
IMPORT XML

DEFINE nbElt INTEGER
DEFINE nbAttr INTEGER
DEFINE nbComment INTEGER
DEFINE nbPI INTEGER
DEFINE nbTxt INTEGER
DEFINE nbCData INTEGER

MAIN
  DEFINE document xml.DomDocument
  DEFINE ind INTEGER
  # Handle arguments
  IF num_args() !=1 THEN
    CALL ExitHelp()
  END IF
  # Create document, load it, and count the nodes
  LET document = xml.DomDocument.Create()
  CALL document.load(arg_val(1))
  CALL CountDoc(document)
  # Display result
  DISPLAY "Results: "
  DISPLAY " Elements: ",nbElt
  DISPLAY " Attributes:",nbAttr
  DISPLAY " Comments: ",nbComment
  DISPLAY " PI: ",nbPI
  DISPLAY " Texts: ",nbTxt
  DISPLAY " CData: ",nbCData
END MAIN

FUNCTION CountDoc(d)
  DEFINE d xml.DomDocument
  DEFINE n xml.DomNode
  LET n = d.getFirstDocumentNode()
  WHILE (n IS NOT NULL )
    CALL Count(n)
    LET n = n.getNextSibling()
  END WHILE
END FUNCTION

FUNCTION Count(n)
  DEFINE n xml.DomNode
  DEFINE child xml.DomNode
  DEFINE next xml.DomNode
  DEFINE node xml.DomNode
  DEFINE ind INTEGER
  DEFINE name STRING
  IF n IS NOT NULL THEN
    IF n.getNodeType() == "COMMENT_NODE" THEN
      LET nbComment = nbComment + 1
    END IF
    IF n.getNodeType() == "ATTRIBUTE_NODE" THEN
      LET nbAttr = nbAttr + 1
    END IF
    IF n.getNodeType() == "PROCESSING_INSTRUCTION_NODE " THEN
      LET nbPI = nbPI + 1
    END IF
    IF n.getNodeType() == "ELEMENT_NODE" THEN
      LET nbElt = nbElt + 1
    END IF
    IF n.getNodeType() == "TEXT_NODE" THEN
      LET nbTxt = nbTxt +1
    END IF
    IF n.getNodeType() == "CDATA_SECTION_NODE" THEN
      LET nbCData = nbCData + 1
    END IF
    IF n.hasChildNodes() THEN
      LET name = n.getLocalName()
      LET child = n.getFirstChild()
      WHILE (child IS NOT NULL )
        CALL Count(child)
        LET child = child.getNextSibling()
      END WHILE
    END IF
    IF n.hasAttributes() THEN
      FOR ind = 1 TO n.getAttributesCount()
        LET node = n.getAttributeNodeItem(ind)
        CALL Count(node)
      END FOR
    END IF
  END IF
END FUNCTION

FUNCTION ExitHelp()
  DISPLAY "DomCount <xml>"
  EXIT PROGRAM
END FUNCTION
```
