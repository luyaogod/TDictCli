---
title: "Example 1: Get a window by name and change the title"
source: "fgl-topics/c_fgl_ClassWindow_example_1.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > Examples > Example 1: Get a window by name and change the title"
type: "concept"
description: "FUNCTION open_customer_window() RETURNS () DEFINE w ui.Window OPEN WINDOW w1 WITH FORM \"customer\" ATTRIBUTES(TEXT=\"Unknown\") LET w = ui.Window.forName(\"w1\") CALL w.setText(\"Customer\") MENU \"Test\" ..."
---

# Example 1: Get a window by name and change the title

```
FUNCTION open_customer_window() RETURNS ()
  DEFINE w ui.Window
  OPEN WINDOW w1 WITH FORM "customer" ATTRIBUTES(TEXT="Unknown")
  LET w = ui.Window.forName("w1")
  CALL w.setText("Customer")
  MENU "Test"
     COMMAND "exit" EXIT MENU
  END MENU
  CLOSE WINDOW w1
END FUNCTION
```
