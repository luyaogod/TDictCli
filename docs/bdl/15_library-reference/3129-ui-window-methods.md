---
title: "ui.Window methods"
source: "fgl-topics/c_fgl_ClassWindow_methods.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Window class > ui.Window methods"
type: "concept"
---

# ui.Window methods

> Methods of the ui.Window class.

| Name | Description |
| --- | --- |
| ui.Window.forName( name STRING ) RETURNS ui.Window | Get a window object by name. |
| ui.Window.getCurrent() RETURNS ui.Window | Get the current window object. |

| Name | Description |
| --- | --- |
| createForm( name STRING ) RETURNS ui.Form | Create a new empty form in a window. |
| findNode( tagName STRING, name STRING ) RETURNS om.DomNode | Search for a specific element in the window. |
| getForm() RETURNS ui.Form | Get the current form of a window. |
| getNode() RETURNS om.DomNode | Get the DOM node of a window. |
| getImage() RETURNS STRING | Get the window icon. |
| getText() RETURNS STRING | Get the window title. |
| setImage( image STRING ) | Set the window icon. |
| setText( title STRING ) | Set the window title. |

## Related links

**Related concepts**  

[Defining the window title](../11_user-interface/1568-defining-the-window-title.md "Use the TEXT attribute to define a title for a window.")
