---
title: "The abstract user interface tree"
source: "fgl-topics/c_fgl_DynamicUI_017.html"
breadcrumb: "User interface > User interface basics > The abstract user interface tree"
type: "concept"
---

# The abstract user interface tree

> The abstract user interface tree is the XML representation of the application forms displayed to the end user.

The abstract user interface tree (AUI tree) is an [DOM](http://www.w3.org/DOM/) tree describing the objects of the user
interface elements of a program at a given time.

A copy of the AUI tree is held by both the front-end and the runtime system.

AUI tree synchronization is automatically done by the runtime system using the [front-end protocol](1522-the-front-end-protocol.md), when the control goes back to the end
user. If a program needs to synchronize the display while processing, it can use the [`ui.Interface.refresh()`](2224-refreshing-the-display-when-processing.md "This topic explains when to use the ui.Interface.refresh() method.")
method.

The programs can manipulate the AUI tree element by using XML utility classes or high-level
built-in classes such as [`ui.Dialog`](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.")
and [`ui.Form`](../15_library-reference/3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.").

## Child topics

- [What does the abstract user interface tree contain?](1511-what-does-the-abstract-user-interface-tree-contain.md)
- [Manipulating the abstract user interface tree](1512-manipulating-the-abstract-user-interface-tree.md)
- [XML node types and attribute names](1513-xml-node-types-and-attribute-names.md)
- [Actions in the abstract user interface tree](1514-actions-in-the-abstract-user-interface-tree.md)
- [Inspecting the AUI tree of a front-end](1515-inspecting-the-aui-tree-of-a-front-end.md)
