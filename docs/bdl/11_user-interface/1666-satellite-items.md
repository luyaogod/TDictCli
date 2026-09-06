---
title: "Satellite items"
source: "fgl-topics/c_fgl_FormSpecFiles_Satellite_Items.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Satellite items"
type: "concept"
---

# Satellite items

> Satellite items are display elements defined outside the LAYOUT section.

Satellite items like the [`TOOLBAR`](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions."), [`TOPMENU`](1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions.") and [`ACTION DEFAULTS`](1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements.") section are form elements independent from the main form
layout, and are defined in addition to the `LAYOUT` section.

```
TOOLBAR -- Toolbar section
...
END
LAYOUT -- Main layout section
...
END
```

## Related links

**Related concepts**  

[LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
