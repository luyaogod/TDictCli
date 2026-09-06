---
title: "ui.Interface.loadTopMenu"
source: "fgl-topics/c_fgl_ClassInterface_loadTopmenu.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.loadTopMenu"
type: "concept"
---

# ui.Interface.loadTopMenu

> Load a default/global topmenu file for all forms of the program.

## Syntax

```
ui.Interface.loadTopMenu(
   path STRING )
```

1. path is the name of a topmenu file, without the extension.

## Usage

Use the `ui.Interface.loadTopMenu()` class method to load a
.4tm file defining a default/global topmenu for all forms.

The purpose of the default/global topmenu is to be displayed in all forms. However, the
default/global toolbar may be displayed in the current window or in the parent window, depending on
the front-end type and window type (modal or normal).

If the interface already contains a global topmenu, it will be replaced by the new topmenu loaded
by this method.

Specify the topmenu filename without the ".4tm" extension.

The resource file is searched in several directories in a given order, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

When loading a topmenu file, [action defaults](../11_user-interface/2261-using-attributes-of-action-defaults.md) are automatically applied.

## Example

```
CALL ui.Interface.loadTopMenu("mytopmenu")
```

## Related links

**Related concepts**  

[Topmenus](../11_user-interface/1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.")
