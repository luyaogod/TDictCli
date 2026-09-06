---
title: "ui.Interface.loadStartMenu"
source: "fgl-topics/c_fgl_ClassInterface_loadStartMenu.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.loadStartMenu"
type: "concept"
---

# ui.Interface.loadStartMenu

> Load the start menu file.

## Syntax

```
ui.Interface.loadStartMenu(
   path STRING )
```

1. path is the name of a start menu file, without the extension.

## Usage

Use the `ui.Interface.loadStartMenu()` class method to load a
.4sm file defining a start menu.

If the interface already contains a start menu, it will be replaced by the new start menu loaded
by this method.

Specify the start menu filename without the ".4sm" extension.

The resource file is searched in several directories in a given order, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

## Example

```
CALL ui.Interface.loadStartMenu("mystartmenu")
```

## Related links

**Related concepts**  

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")

[Start menus](../11_user-interface/2446-start-menus.md "Start menus define a tree of application programs that can be started.")
