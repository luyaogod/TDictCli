---
title: "Example 3: Writing a DOM tree to a SAX handler"
source: "fgl-topics/c_fgl_ClassDomNode_example_3.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > Examples > Example 3: Writing a DOM tree to a SAX handler"
type: "concept"
description: "The following example outputs a DOM tree without indentation. MAIN DEFINE d om.DomDocument DEFINE r, n, t, w om.DomNode DEFINE dh om.SaxDocumentHandler DEFINE i INTEGER LET dh = ..."
---

# Example 3: Writing a DOM tree to a SAX handler

The following example outputs a DOM tree without
indentation.

```
MAIN
  DEFINE d  om.DomDocument 
  DEFINE r, n, t, w om.DomNode 
  DEFINE dh om.SaxDocumentHandler 

  DEFINE i INTEGER

  LET dh = om.XmlWriter.createPipeWriter("cat")
  CALL dh.setIndent(FALSE)

  LET d = om.DomDocument.create("Vehicles")
  LET r = d.getDocumentElement()

  LET n = r.createChild("Car")
  CALL n.setAttribute("name","Corolla")
  CALL n.setAttribute("color","Blue")
  CALL n.setAttribute("weight","1546")

  LET t = d.createChars("Nice car!")
  CALL n.appendChild(t)

  LET n = r.createChild("Bus")
  CALL n.setAttribute("name","Maxibus")
  CALL n.setAttribute("color","yellow")
  CALL n.setAttribute("weight","5278")
  FOR i=1 TO 4
    LET w = n.createChild("Wheel")
    CALL w.setAttribute("width","315")
    CALL w.setAttribute("diameter","925")
  END FOR

  CALL r.write(dh)

END MAIN
```
