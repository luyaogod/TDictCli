---
title: "Defining the topmenu in a form file"
source: "fgl-topics/c_fgl_topmenus_007.html"
breadcrumb: "User interface > Form definitions > Topmenus > Defining the topmenu in a form file"
type: "concept"
---

# Defining the topmenu in a form file

> Topmenus can be defined in the form specification file within the TOPMENU section.

Form topmenus will only be displayed in the window where the form is loaded. Only one topmenu can
be defined in a form file. It is recommended that topmenu item attributes that are common to
toolbar buttons are centralized in action defaults.

## Example

```
TOPMENU tm
  GROUP form (TEXT="Form")
    COMMAND help (TEXT="Help", IMAGE="quest")
    COMMAND quit (TEXT="Quit")
  END
  ...
END

LAYOUT
GRID
{
  ...
```

## Related links

**Related concepts**  

[TOPMENU section](1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions.")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
