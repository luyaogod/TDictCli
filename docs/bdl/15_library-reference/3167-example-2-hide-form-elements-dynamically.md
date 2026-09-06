---
title: "Example 2: Hide form elements dynamically"
source: "fgl-topics/c_fgl_ClassForm_example_2.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > Examples > Example 2: Hide form elements dynamically"
type: "concept"
description: "MAIN DEFINE w ui.Window DEFINE f ui.Form DEFINE rec RECORD custid INTEGER, custname VARCHAR(40) END RECORD OPEN FORM f1 FROM \"customer\" DISPLAY FORM f1 LET w = ui.Window.getCurrent() LET f = ..."
---

# Example 2: Hide form elements dynamically

```
MAIN
  DEFINE w ui.Window
  DEFINE f ui.Form
  DEFINE rec RECORD
             custid INTEGER,
             custname VARCHAR(40)
         END RECORD
  OPEN FORM f1 FROM "customer"
  DISPLAY FORM f1
  LET w = ui.Window.getCurrent()
  LET f = w.getForm()
  INPUT BY NAME rec.*
    ON ACTION hide
      CALL f.setFieldHidden("customer.custid",1)
      CALL f.setElementHidden("label_custid",1)
    ON ACTION show
      CALL f.setFieldHidden("customer.custid",0)
      CALL f.setElementHidden("label_custid",0)
  END INPUT
END MAIN
```
