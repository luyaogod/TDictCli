---
title: "Defining toolbars in the form file"
source: "fgl-topics/c_fgl_toolbars_006.html"
breadcrumb: "User interface > Form definitions > Toolbars > Defining toolbars in the form file"
type: "concept"
---

# Defining toolbars in the form file

> Toolbars can be defined in the form specification file within the TOOLBAR section.

Form toolbars are only displayed in the window where the form is loaded. Only one toolbar can be
defined in a form file. It is recommended that toolbar button attributes that are common to
topmenu options are centralized in action defaults.

## Example

```
TOOLBAR tb
  ITEM accept ( TEXT="Ok", IMAGE="ok" )
  ITEM cancel ( TEXT="Cancel", IMAGE="cancel" )
  SEPARATOR
  ...
END

LAYOUT
GRID
{
  ...
```

## Related links

**Related concepts**  

[TOOLBAR section](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions.")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
