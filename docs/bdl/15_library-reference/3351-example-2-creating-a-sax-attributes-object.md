---
title: "Example 2: Creating a SAX attributes object"
source: "fgl-topics/c_fgl_ClassSaxAttributes_example_2.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > Examples > Example 2: Creating a SAX attributes object"
type: "concept"
description: "MAIN DEFINE a, c om.SaxAttributes, i, m INTEGER LET a = om.SaxAttributes.create() CALL a.addAttribute(\"name\",\"Jo\") CALL a.addAttribute(\"birth\",\"2008-12-20\") DISPLAY a.getValue(\"birth\") LET c = ..."
---

# Example 2: Creating a SAX attributes object

```
MAIN
   DEFINE a, c om.SaxAttributes,
          i, m INTEGER
   LET a = om.SaxAttributes.create()
   CALL a.addAttribute("name","Jo")
   CALL a.addAttribute("birth","2008-12-20")
   DISPLAY a.getValue("birth")
   LET c = om.SaxAttributes.copy( a )
   LET m = a.getLength()
   DISPLAY "Len = ", m
   FOR i=1 TO m
       DISPLAY i, ": ", a.getName(i),
                  " = ", a.getValueByIndex(i)
   END FOR
   CALL a.removeAttribute("birth")
   CALL c.setAttributes(a)
   CALL a.clear()
   DISPLAY c.getLength()
END MAIN
```
