---
title: "ui.Interface.loadToolBar"
source: "fgl-topics/c_fgl_ClassInterface_loadToolbar.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.loadToolBar"
type: "concept"
---

# ui.Interface.loadToolBar

> Load a default/global toolbar file for all forms of the program.

## Syntax

```
ui.Interface.loadToolBar(
   path STRING )
```

1. path is the name of a toolbar file, without the extension.

## Usage

Use the `ui.Interface.loadToolBar()` class method to load a
.4tb file defining a default/global toolbar for all forms.

The purpose of the default/global toolbar is to be displayed in all forms. However, the
default/global toolbar may be displayed in the current window or in the parent window, depending on
the front-end type and window type (modal or normal).

If the interface already contains a default/global toolbar, it will be replaced by the new
toolbar loaded by this method.

Specify the toolbar filename without the ".4tb" extension.

The resource file is searched in several directories in a given order, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

When loading a toolbar file, [action defaults](../11_user-interface/2261-using-attributes-of-action-defaults.md) are automatically applied.

## Example

```
CALL ui.Interface.loadToolBar("mytoolbar")
```

## Related links

**Related concepts**  

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")

[Toolbars](../11_user-interface/1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.")
