---
title: "Example 1: Creating a DOM tree"
source: "fgl-topics/c_fgl_ClassDomNode_example_1.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > Examples > Example 1: Creating a DOM tree"
type: "concept"
description: "To create a DOM tree with the following structure (represented in XML format): <Vehicles> <Car name=\"Corolla\" color=\"Blue\" weight=\"1546\">Nice car!</Car> <Bus name=\"Maxibus\" color=\"Yellow\" ..."
---

# Example 1: Creating a DOM tree

To create a [DOM](../09_advanced-features/0948-xml-support.md "Introduces to DOM and SAX standards and describes the XML utility classes built-in the language.") tree with the
following structure (represented in XML format):

```
<Vehicles>
   <Car name="Corolla" color="Blue" weight="1546">Nice car!</Car>
   <Bus name="Maxibus" color="Yellow" weight="5278">
      <Wheel width="315" diameter="925" />
      <Wheel width="315" diameter="925" />
      <Wheel width="315" diameter="925" />
      <Wheel width="315" diameter="925" />
   </Bus>
</Vehicles>
```

You write the
following:

```
MAIN
  DEFINE d  om.DomDocument 
  DEFINE r, n, t, w om.DomNode
  DEFINE i INTEGER

  LET d = om.DomDocument.create("Vehicles")
  LET r = d.getDocumentElement()

  LET n = r.createChild("Car")
  CALL n.setAttribute("name","Corolla")
  CALL n.setAttribute("color","Blue")
  CALL n.setAttribute("weight","1546")

  LET t = d.createChars("Nice car!")
  CALL n.appendChild(t)
  LET t = d.createEntity("nbsp")
  CALL n.appendChild(t)
  LET t = d.createChars("Yes, very nice!")
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

  CALL r.writeXml("Vehicles.xml")

END MAIN
```
