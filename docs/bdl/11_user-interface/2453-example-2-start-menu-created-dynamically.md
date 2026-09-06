---
title: "Example 2: Start menu created dynamically"
source: "fgl-topics/c_fgl_startmenus_011.html"
breadcrumb: "User interface > User interface programming > Start menus > Examples > Example 2: Start menu created dynamically"
type: "concept"
description: "MAIN DEFINE aui om.DomNode DEFINE sm om.DomNode DEFINE smg om.DomNode DEFINE smc om.DomNode LET aui = ui.Interface.getRootNode() LET sm = aui.createChild(\"StartMenu\") LET smg = ..."
---

# Example 2: Start menu created dynamically

```
MAIN
  DEFINE aui om.DomNode 
  DEFINE sm  om.DomNode 
  DEFINE smg om.DomNode 
  DEFINE smc om.DomNode 

  LET aui = ui.Interface.getRootNode()

  LET sm = aui.createChild("StartMenu")

  LET smg = createStartMenuGroup(sm,"Ordering")
  LET smc = createStartMenuCommand(smg,"Orders","fglrun orders",NULL)
  LET smc = createStartMenuCommand(smg,"Customers","fglrun custs",NULL)
  LET smc = createStartMenuCommand(smg,"Items","fglrun items",NULL)
  LET smc = createStartMenuCommand(smg,"Reports","fglrun reports",NULL)
  LET smg = createStartMenuGroup(sm,"Configuration")
  LET smc = createStartMenuCommand(smg,"Database","fglrun dbseconf",NULL)
  LET smc = createStartMenuCommand(smg,"Users","fglrun userconf",NULL)
  LET smc = createStartMenuCommand(smg,"Printers","fglrun prntconf",NULL)

  MENU "Example"
    COMMAND "Quit"
      EXIT PROGRAM
  END MENU

END MAIN

FUNCTION createStartMenuGroup(p,t)
  DEFINE p om.DomNode 
  DEFINE t STRING
  DEFINE s om.DomNode 
  LET s = p.createChild("StartMenuGroup")
  CALL s.setAttribute("text",t)
  RETURN s 
END FUNCTION

FUNCTION createStartMenuCommand(p,t,c,i)
  DEFINE p om.DomNode 
  DEFINE t,c,i STRING
  DEFINE s om.DomNode 
  LET s = p.createChild("StartMenuCommand")
  CALL s.setAttribute("text",t)
  CALL s.setAttribute("exec",c)
  CALL s.setAttribute("image",i)
  RETURN s 
END FUNCTION
```
