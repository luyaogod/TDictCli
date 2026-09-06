---
title: "Example 1: Displaying SAX attributes of an XML node"
source: "fgl-topics/c_fgl_ClassSaxAttributes_example_1.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > Examples > Example 1: Displaying SAX attributes of an XML node"
type: "concept"
description: "FUNCTION displayAttributes( a ) DEFINE a om.SaxAttributes DEFINE i, m INTEGER LET m = a.getLength() FOR i=1 to m DISPLAY a.getName(i) || \"=[\" || a.getValueByIndex(i) || \"]\" END FOR END FUNCTION"
---

# Example 1: Displaying SAX attributes of an XML node

```
FUNCTION displayAttributes( a )
   DEFINE a om.SaxAttributes 
   DEFINE i, m INTEGER
   LET m = a.getLength()
   FOR i=1 to m
      DISPLAY a.getName(i) || "=[" || a.getValueByIndex(i) || "]"
   END FOR
END FUNCTION
```
