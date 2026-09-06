---
title: "Creating the start menu dynamically"
source: "fgl-topics/c_fgl_startmenus_008.html"
breadcrumb: "User interface > User interface programming > Start menus > Creating the start menu dynamically"
type: "concept"
---

# Creating the start menu dynamically

> Start menu can be created dynamically with the om.DomNode class.

First, get the abstract user interface root node:

```
DEFINE aui om.DomNode 
LET aui = ui.Interface.getRootNode()
```

Next, create a node with the "`StartMenu`" tag name:

```
DEFINE sm om.DomNode 
LET sm = aui.createChild("StartMenu")
```

Next, create a "`StartMenuGroup`" node to group a couple of command nodes:

```
DEFINE smg om.DomNode 
LET smg = sm.createChild("StartMenuGroup")
CALL smg.setAttribute("text","Programs")
```

Then, create "`StartMenuCommand`" nodes for each program and, if needed, add
"`StartMenuSeparator`" nodes to separate entries:

```
DEFINE smc, sms om.DomNode 
LET smc = smg.createChild("StartMenuCommand")
CALL smc.setAttribute("text","Orders")
CALL smc.setAttribute("exec","fglrun orders.42r")
LET smc = smg.createChild("StartMenuCommand")
CALL smc.setAttribute("text","Customers")
CALL smc.setAttribute("exec","fglrun customers.42r")
LET sms = smg.createChild("StartMenuSeparator")
LET smc = smg.createChild("StartMenuCommand")
CALL smc.setAttribute("text","Items")
CALL smc.setAttribute("exec","fglrun items.42r")
```

## Related links

**Related concepts**  

[The DomNode class](../15_library-reference/3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")

[The Interface class](../15_library-reference/3095-the-interface-class.md "The ui.Interface class provides methods to manipulate the user interface.")

[The abstract user interface tree](1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")
