---
title: "ui.Window.getNode"
source: "fgl-topics/c_fgl_ClassWindow_getNode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods > ui.Window.getNode"
type: "concept"
---

# ui.Window.getNode

> Get the DOM node of a window.

## Syntax

```
getNode()
  RETURNS om.DomNode
```

## Usage

The `getNode()` method returns the `om.DomNode` object
corresponding to the window object.

Declare a variable of type `om.DomNode` to hold the DOM node object
reference.

Consider using the `ui.Dialog.getForm()` method to get the form used
by the current dialog.

## Example

```
DEFINE w ui.Window, n om.DomNode
OPEN WINDOW w1 WITH FORM "custform"
LET w = ui.Window.getCurrent()
LET n = w.getNode()
...
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")

[ui.Dialog.getForm](3200-ui-dialog-getform.md "Returns the current form used by the dialog.")
