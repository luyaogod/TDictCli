---
title: "What does the abstract user interface tree contain?"
source: "fgl-topics/c_fgl_DynamicUI_018.html"
breadcrumb: "User interface > User interface basics > The abstract user interface tree > What does the abstract user interface tree contain?"
type: "concept"
description: "The abstract user interface defines a tree of objects organized by parent/child relationship. The different kinds of user interface objects are defined by attributes. The AUI tree can be serialized as ..."
---

# What does the abstract user interface tree contain?

The abstract user interface defines a tree of objects organized by parent/child relationship. The
different kinds of user interface objects are defined by attributes. The AUI tree can be serialized
as text based on the [XML](http://www.w3c.org/XML)
standard notation.

The following example shows a part of an AUI tree defining a [toolbar](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.") serialized with the XML
notation:

```
<ToolBar>
  <ToolBarItem name="f5" text="List" image="list" />
  <ToolBarSeparator/>
  <ToolBarItem name="Query" text="Query" image="search" />
  <ToolBarItem name="Add" text="Append" image="add" />
  ...
</ToolBar>
```

## Related links

**Related concepts**  

[Manipulating the abstract user interface tree](1512-manipulating-the-abstract-user-interface-tree.md "Manipulating the abstract user interface tree")
