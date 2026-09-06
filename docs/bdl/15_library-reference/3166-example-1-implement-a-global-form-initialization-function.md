---
title: "Example 1: Implement a global form initialization function"
source: "fgl-topics/c_fgl_ClassForm_example_1.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > Examples > Example 1: Implement a global form initialization function"
type: "concept"
description: "MAIN CALL ui.Form.setDefaultInitializerFunction(FUNCTION init) OPEN FORM f1 FROM \"items\" DISPLAY FORM f1 -- Form appears in the default SCREEN window OPEN WINDOW w1 WITH FORM \"customer\" OPEN WINDOW w2 ..."
---

# Example 1: Implement a global form initialization function

```
MAIN
  CALL ui.Form.setDefaultInitializerFunction(FUNCTION init)
  OPEN FORM f1 FROM "items"
  DISPLAY FORM f1 -- Form appears in the default SCREEN window 
  OPEN WINDOW w1 WITH FORM "customer"
  OPEN WINDOW w2 WITH FORM "orders"
  DISPLAY FORM f1 -- Form appears in w2 window 
  MENU "Test"
    COMMAND "exit" EXIT MENU
  END MENU
END MAIN

FUNCTION init(form ui.Form) RETURNS ()
  DEFINE n om.DomNode
  CALL form.loadTopMenu("mymenu")
  LET n = form.getNode()
  DISPLAY "Init: ", n.getAttribute("name")
END FUNCTION
```
