---
title: "Example 2: Get a the current form and hide a groupbox"
source: "fgl-topics/c_fgl_ClassWindow_example_2.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > Examples > Example 2: Get a the current form and hide a groupbox"
type: "concept"
description: "MAIN DEFINE w ui.Window DEFINE f ui.Form OPEN FORM f1 FROM \"customer\" DISPLAY FORM f1 LET w = ui.Window.getCurrent() LET f = w.getForm() MENU \"Test\" COMMAND \"hide\" CALL f.setElementHidden(\"gb1\",1) ..."
---

# Example 2: Get a the current form and hide a groupbox

```
MAIN
  DEFINE w ui.Window
  DEFINE f ui.Form
  OPEN FORM f1 FROM "customer"
  DISPLAY FORM f1
  LET w = ui.Window.getCurrent()
  LET f = w.getForm()
  MENU "Test"
    COMMAND "hide" CALL f.setElementHidden("gb1",1)
    COMMAND "exit" EXIT MENU
  END MENU
END MAIN
```
