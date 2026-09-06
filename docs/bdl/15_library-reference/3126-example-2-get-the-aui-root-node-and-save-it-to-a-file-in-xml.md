---
title: "Example 2: Get the AUI root node and save it to a file in XML format"
source: "fgl-topics/c_fgl_ClassInterface_example_2.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > Examples > Example 2: Get the AUI root node and save it to a file in XML format"
type: "concept"
description: "MAIN DEFINE n om.DomNode MENU \"Test\" COMMAND \"SaveUI\" LET n = ui.Interface.getRootNode() CALL n.writeXml(\"auitree.xml\") COMMAND \"Exit\" EXIT MENU END MENU END MAIN"
---

# Example 2: Get the AUI root node and save it to a file in XML format

```
MAIN
  DEFINE n om.DomNode 
  MENU "Test"
    COMMAND "SaveUI"
      LET n = ui.Interface.getRootNode()
      CALL n.writeXml("auitree.xml")
    COMMAND "Exit"
      EXIT MENU
  END MENU
END MAIN
```
