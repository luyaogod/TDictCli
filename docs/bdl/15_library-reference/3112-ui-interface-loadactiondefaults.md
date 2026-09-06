---
title: "ui.Interface.loadActionDefaults"
source: "fgl-topics/c_fgl_ClassInterface_loadActionDefaults.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.loadActionDefaults"
type: "concept"
---

# ui.Interface.loadActionDefaults

> Load the default action defaults file.

## Syntax

```
ui.Interface.loadActionDefaults(
   path STRING )
```

1. path is the name of the action defaults file, without the extension.

## Usage

Use the `ui.Interface.loadActionDefaults()` class method to load a
.4ad file defining action defaults for all program forms.

If the interface already contains action defaults, these will be replaced by the new action
defaults loaded by this method.

Specify the action defaults filename without the ".4ad"
extension.

The resource file is searched in several directories in a given order, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

## Example

```
CALL ui.Interface.loadActionDefaults("mydefaults")
```

For a complete example, see [Example 3: Loading custom resources](3127-example-3-loading-custom-resources.md).

## Related links

**Related concepts**  

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
