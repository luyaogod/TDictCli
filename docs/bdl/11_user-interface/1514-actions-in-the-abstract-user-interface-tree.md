---
title: "Actions in the abstract user interface tree"
source: "fgl-topics/c_fgl_DynamicUI_023.html"
breadcrumb: "User interface > User interface basics > The abstract user interface tree > Actions in the abstract user interface tree"
type: "concept"
description: "The abstract user interface identifies all possible actions that can be received by the current interactive instruction with a list of Action nodes. The list of possible actions are held by a Dialog ..."
---

# Actions in the abstract user interface tree

The abstract user interface identifies all possible actions that can be received by the current
interactive instruction with a list of `Action` nodes. The list of possible
actions are held by a `Dialog` node. An `Action` node is identified
by the '`name`' attribute and defines common properties such as the accelerator
key, default image, and default text.

Interactive elements are bound to `Action` nodes by the '`name`'
attribute. For example, a toolbar button (a.k.a toolbar item) with the name 'cancel' is bound to
the `Action` node having the name 'cancel', which in turn defines the accelerator
key, the default text, and the default image for the button.

![Action nodes binding in the AUI Tree binding](../_images/AUIFig01.jpg)

*Action nodes binding in the AUI Tree*

When a form element triggers an action (such as button in a grid or a toolbar button, or a
button in the action panel) an `ActionEvent` node is sent to the runtime system.
The name of the `ActionEvent` node identifies what `Action` has
been triggered.

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
