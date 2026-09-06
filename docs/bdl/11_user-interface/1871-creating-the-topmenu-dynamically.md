---
title: "Creating the topmenu dynamically"
source: "fgl-topics/c_fgl_topmenus_010.html"
breadcrumb: "User interface > Form definitions > Topmenus > Creating the topmenu dynamically"
type: "concept"
---

# Creating the topmenu dynamically

> Topmenus can be created at runtime by creating the corresponding XML representation in the AUI tree.

In the abstract user interface tree, the `TopMenu` node must be created under the
`Form` node, and must contain `TopMenuGroup` nodes. The
`TopMenuGroup` nodes group topmenu commands and other topmenu groups. A
`TopMenuCommand` is a leaf node in the topmenu tree that will trigger an action:

```
TopMenu
   +- TopMenuGroup
      +- TopMenuCommand
      +- TopMenuCommand
      +- TopMenuCommand
   +- TopMenuGroup
      +- TopMenuGroup
         +- TopMenuCommand
         +- TopMenuCommand
      +- TopMenuGroup
         +- TopMenuCommand
         +- TopMenuCommand
         +- TopMenuCommand
```

This example shows how to create a topmenu in all forms by using the default initialization
function and the `om.DomNode` class:

```
CALL ui.Form.setDefaultInitializerFunction(FUNCTION myinit)
OPEN FORM f1 FROM "form1"
DISPLAY FORM f1
...
FUNCTION myinit(form ui.Form) RETURNS ()
  DEFINE n om.DomNode
  LET n = form.getNode()
  ...
END FUNCTION
```

After getting the DOM node of the form, create a node with the "`TopMenu`" tag
name:

```
DEFINE tm om.DomNode
LET tm = n.createChild("TopMenu")
```

For each Topmenu group, create a subnode with the "`TopMenuGroup`"
tag name and set the attributes to define the group:

```
DEFINE tmg om.DomNode
LET tmg = tm.createChild("TopMenuGroup")
CALL tmg.setAttribute("text","Reports")
```

For each Topmenu option, create a sub-node in a group node with
the "`TopMenuCommand`" tag name and set the attributes
to define the option:

```
DEFINE tmi om.DomNode
LET tmi = tmg.createChild("TopMenuCommand")
CALL tmi.setAttribute("name","report")
CALL tmi.setAttribute("text","Order report")
CALL tmi.setAttribute("comment","Orders entered today")
CALL tmi.setAttribute("image","smiley")
```

If needed, you can create a "`TopMenuSeparator`"
node inside a group, to separate menu options:

```
DEFINE tms om.DomNode
LET tms = tmg.createChild("TopMenuSeparator")
```

## Related links

**Related concepts**  

[The DomNode class](../15_library-reference/3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")
