---
title: "Example 1: Simple MessageServer usage"
source: "fgl-topics/c_fgl_ClassMessageServer_example_1.html"
breadcrumb: "Library reference > Built-in packages > The base package > The MessageServer class > Examples > Example 1: Simple MessageServer usage"
type: "concept"
description: "Compile the following program and start multiple instances of the program with fglrun . Then hit the \"Send F1\" button in one of the program instances, and check that an error message is displayed by ..."
---

# Example 1: Simple MessageServer usage

Compile the following program and start multiple instances of the program with
fglrun. Then hit the "Send F1" button in one of the program instances, and check
that an error message is displayed by other program instances.

```
MAIN
  CALL base.MessageServer.connect()
  MENU "test"
    COMMAND "Send F1"
      CALL base.MessageServer.send("f1")
    COMMAND KEY (F1)
      ERROR SFMT("Key F1 received at %1",CURRENT HOUR TO SECOND)
    COMMAND "quit"
      EXIT MENU
  END MENU
END MAIN
```
