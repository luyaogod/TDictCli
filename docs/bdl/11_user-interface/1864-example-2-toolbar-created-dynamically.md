---
title: "Example 2: Toolbar created dynamically"
source: "fgl-topics/c_fgl_toolbars_012.html"
breadcrumb: "User interface > Form definitions > Toolbars > Examples > Example 2: Toolbar created dynamically"
type: "concept"
description: "MAIN DEFINE aui om.DomNode DEFINE tb om.DomNode DEFINE tbi om.DomNode DEFINE tbs om.DomNode LET aui = ui.Interface.getRootNode() LET tb = aui.createChild(\"ToolBar\") LET tbi = ..."
---

# Example 2: Toolbar created dynamically

```
MAIN
  DEFINE aui om.DomNode 
  DEFINE tb  om.DomNode 
  DEFINE tbi om.DomNode 
  DEFINE tbs om.DomNode 

  LET aui = ui.Interface.getRootNode()

  LET tb = aui.createChild("ToolBar")

  LET tbi = createToolBarItem(tb,"f1","Help","Show help","help")
  LET tbs = createToolBarSeparator(tb)
  LET tbi = createToolBarItem(tb,"upd","Modify","Modify current record","change")
  LET tbi = createToolBarItem(tb,"del","Remove","Remove current record","delete")
  LET tbi = createToolBarItem(tb,"add","Append","Add a new record","add")
  LET tbs = createToolBarSeparator(tb)
  LET tbi = createToolBarItem(tb,"xxx","Exit","Quit application","quit")

  MENU "Example"
    COMMAND KEY(F1)
      DISPLAY "F1 action received"
    COMMAND "upd"
      DISPLAY "Update action received"
    COMMAND "Del"
      DISPLAY "Delete action received"
    COMMAND "Add"
      DISPLAY "Append action received"
    COMMAND "xxx"
      EXIT PROGRAM
  END MENU

END MAIN

FUNCTION createToolBarSeparator(tb)
  DEFINE tb om.DomNode 
  DEFINE tbs om.DomNode 
  LET tbs = tb.createChild("ToolBarSeparator")
  RETURN tbs 
END FUNCTION

FUNCTION createToolBarItem(tb,n,t,c,i)
  DEFINE tb om.DomNode 
  DEFINE n,t,c,i VARCHAR(100)
  DEFINE tbi om.DomNode 
  LET tbi = tb.createChild("ToolBarItem")
  CALL tbi.setAttribute("name",n)
  CALL tbi.setAttribute("text",t)
  CALL tbi.setAttribute("comment",c)
  CALL tbi.setAttribute("image",i)
  RETURN tbi 
END FUNCTION
```
