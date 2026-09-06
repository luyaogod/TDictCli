---
title: "Example 1: Creating a DOM document"
source: "fgl-topics/c_fgl_ClassDomDocument_example_1.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > Examples > Example 1: Creating a DOM document"
type: "concept"
description: "MAIN DEFINE d om.DomDocument DEFINE r om.DomNode LET d = om.DomDocument.create(\"MyDocument\") LET r = d.getDocumentElement() END MAIN"
---

# Example 1: Creating a DOM document

```
MAIN
  DEFINE d  om.DomDocument 
  DEFINE r om.DomNode 
  LET d = om.DomDocument.create("MyDocument")
  LET r = d.getDocumentElement()
END MAIN
```
