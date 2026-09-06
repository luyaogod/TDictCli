---
title: "Example 1: Get the type and version of the front-end"
source: "fgl-topics/c_fgl_ClassInterface_example_1.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > Examples > Example 1: Get the type and version of the front-end"
type: "concept"
description: "MAIN MENU \"Test\" COMMAND \"Get\" DISPLAY \"Name = \" || ui.Interface.getFrontEndName() DISPLAY \"Version = \" || ui.Interface.getFrontEndVersion() COMMAND \"Exit\" EXIT MENU END MENU END MAIN"
---

# Example 1: Get the type and version of the front-end

```
MAIN
  MENU "Test"
    COMMAND "Get"
      DISPLAY "Name = " || ui.Interface.getFrontEndName()
      DISPLAY "Version = " || ui.Interface.getFrontEndVersion()
    COMMAND "Exit"
      EXIT MENU
  END MENU
END MAIN
```
