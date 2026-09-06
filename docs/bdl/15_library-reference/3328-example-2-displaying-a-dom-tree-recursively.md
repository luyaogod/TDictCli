---
title: "Example 2: Displaying a DOM tree recursively"
source: "fgl-topics/c_fgl_ClassDomNode_example_2.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > Examples > Example 2: Displaying a DOM tree recursively"
type: "concept"
description: "The following example displays a DOM tree content recursively: FUNCTION displayDomNode(n,e) DEFINE n om.DomNode DEFINE e, i, s INTEGER LET s = e*2 DISPLAY s SPACES || \"Tag: \" || n.getTagName() DISPLAY ..."
---

# Example 2: Displaying a DOM tree recursively

The following example displays a [DOM](../09_advanced-features/0948-xml-support.md "Introduces to DOM and SAX standards and describes the XML utility classes built-in the language.") tree
content recursively:

```
FUNCTION displayDomNode(n,e)
  DEFINE n om.DomNode
  DEFINE e, i, s INTEGER

  LET s = e*2
  DISPLAY s SPACES || "Tag: " || n.getTagName()

  DISPLAY s SPACES || "Attributes:"
  FOR i=1 TO n.getAttributesCount()
     DISPLAY s SPACES || "  " || n.getAttributeName(i) || 
                     "=[" || n.getAttributeValue(i) ||"]"
  END FOR
  LET n = n.getFirstChild()

  DISPLAY s SPACES || "Child Nodes:"
  WHILE n IS NOT NULL
    CALL displayDomNode(n,e+1)
    LET n = n.getNext()
  END WHILE

END FUNCTION
```
