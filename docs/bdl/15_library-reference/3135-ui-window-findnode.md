---
title: "ui.Window.findNode"
source: "fgl-topics/c_fgl_ClassWindow_findNode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods > ui.Window.findNode"
type: "concept"
---

# ui.Window.findNode

> Search for a specific element in the window.

## Syntax

```
findNode(
   tagName STRING,
   name STRING )
  RETURNS om.DomNode
```

1. tagName defines the type of the node.
2. name defines the name of the node.

## Usage

The `findNode()` method allows you to search for a specific DOM node in the
abstract representation of the window. You search for a child node by giving its type and the name
of the element (the tagname and the value of the 'name' attribute).

The method returns the first element found matching the specified tagname and node name. Window
element names must be unique for the same type of nodes, if you want to distinguish all
elements.

The `findNode()` method is provided for
`ui.Window` class for specific
cases when the window does not contain a form. For
windows containing a form, use the
`ui.Form.findNode()` method
instead.

## Example

```
DEFINE w ui.Window, n om.DomNode
OPEN WINDOW w1 WITH FORM "custform"
LET w = ui.Window.getCurrent()
LET n = w.findNode("FormField","customer.cust_name")
...
```

## Related links

**Related concepts**  

[ui.Form.findNode](3151-ui-form-findnode.md "Search for a child node in the form.")
