---
title: "ui.Form.loadToolBar"
source: "fgl-topics/c_fgl_ClassForm_loadToolbar.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.loadToolBar"
type: "concept"
---

# ui.Form.loadToolBar

> Load the form toolbar.

## Syntax

```
loadToolBar(
   path STRING )
```

1. path is the name of the toolbar file without
   extension.

## Usage

Load a toolbar XML definition file into the form with the `loadToolBar()`
method.

The `loadToolBar()` method is commonly used in the form initialization
function.

Specify the toolbar filename without the ".4tb" extension.

The resource file is searched in several directories in a given order, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

When loading a toolbar file, [action defaults](../11_user-interface/2261-using-attributes-of-action-defaults.md) are automatically applied.

If the form already contains a toolbar (at the form level, not the [global default toolbar](3115-ui-interface-loadtoolbar.md "Load a default/global toolbar file for all forms of the program.")), it will be replaced by
the new toolbar loaded from this method.

## Related links

**Related concepts**  

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")

[Toolbars](../11_user-interface/1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.")
