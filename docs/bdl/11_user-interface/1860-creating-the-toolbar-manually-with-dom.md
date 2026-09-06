---
title: "Creating the toolbar manually with DOM"
source: "fgl-topics/c_fgl_toolbars_009.html"
breadcrumb: "User interface > Form definitions > Toolbars > Creating the toolbar manually with DOM"
type: "concept"
---

# Creating the toolbar manually with DOM

> Toolbars can be created at runtime by creating the corresponding XML representation in the AUI tree.

This example shows how to create a toolbar in all forms by using the default initialization
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

After getting the DOM node of the form, create a node with the "`ToolBar`" tag
name:

```
DEFINE tb om.DomNode
LET tb = n.createChild("ToolBar")
```

For each toolbar button, create a sub-node with the "`ToolBarItem`"
tag name and set the attributes to define the button:

```
DEFINE tbi om.DomNode
LET tbi = tb.createChild("ToolBarItem")
CALL tbi.setAttribute("name","update")
CALL tbi.setAttribute("text","Modify")
CALL tbi.setAttribute("comment","Modify the current record")
CALL tbi.setAttribute("image","change")
```

If needed, you can create a "`ToolBarSeparator`" node to separate toolbar
buttons:

```
DEFINE tbs om.DomNode
LET tbs = tb.createChild("ToolBarSeparator")
```

## Related links

**Related concepts**  

[The DomNode class](../15_library-reference/3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")

[Example 2: Toolbar created dynamically](1864-example-2-toolbar-created-dynamically.md "Example 2: Toolbar created dynamically")
