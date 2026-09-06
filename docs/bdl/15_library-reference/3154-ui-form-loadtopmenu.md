---
title: "ui.Form.loadTopMenu"
source: "fgl-topics/c_fgl_ClassForm_loadTopmenu.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.loadTopMenu"
type: "concept"
---

# ui.Form.loadTopMenu

> Load the form topmenu.

## Syntax

```
loadTopMenu(
   path STRING )
```

1. path is the name of the topmenu file without extension.

## Usage

Load a topmenu XML definition file into the form with the
`loadTopMenu()` method.

The `loadTopMenu()`
method is commonly used in the form initialization function.

Specify the topmenu filename without the .4tm extension.

The resource file is searched in several directories in a given order, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

When loading a topmenu file, [action defaults](../11_user-interface/2261-using-attributes-of-action-defaults.md) are automatically applied.

If the form already contains a topmenu (at the form level, not the [global default topmenu](3116-ui-interface-loadtopmenu.md "Load a default/global topmenu file for all forms of the program.")), it will be replaced by
the new topmenu loaded by this method.

## Related links

**Related concepts**  

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")

[Topmenus](../11_user-interface/1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.")
