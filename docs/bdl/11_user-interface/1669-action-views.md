---
title: "Action views"
source: "fgl-topics/c_fgl_FormSpecFiles_Action_Views.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Action views"
type: "concept"
---

# Action views

> An action view defines a form item that can trigger an action in the program.

## Action views as satellite items like `TOOLBAR`

Below is [`TOOLBAR`](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions.")
section defining a toolbar button using the `close` action name. Here no layout
tag is used because the toolbar item is part of the toolbar graphical object (it will not appear
in the form layout area):

```
TOOLBAR
  ITEM close (TEXT="Close")
END
```

## Action views in a `GRID` container

In a [`GRID`](1722-grid-container.md "Defines a layout area based on a grid of cells.") container,
the position and size of the element is defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container."), while the rendering and behavior is
defined in the [`ATTRIBUTES`](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.") section. Both parts are bound by the name of the item tag. The
item tag name is local to the .per file and is not available at runtime.

The example defines a `BUTTON` form item, where the item tag name is
"`b_close`", and the button name (and the action name) is
"`close`":

```
LAYOUT
GRID
{
  ...
  [b_close     ]
}
END
END
...
ATTRIBUTES
BUTTON b_close: close, TEXT="Close";
END
```

## Related links

**Related concepts**  

[Dialog actions](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.")
