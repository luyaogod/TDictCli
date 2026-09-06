---
title: "ui.Form.loadActionDefaults"
source: "fgl-topics/c_fgl_ClassForm_loadActionDefaults.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.loadActionDefaults"
type: "concept"
---

# ui.Form.loadActionDefaults

> Load form action defaults.

## Syntax

```
loadActionDefaults(
   path STRING )
```

1. path is the name of the action defaults
   file without extension.

## Usage

Load form specific action defaults at runtime with the `loadActionDefaults()`
method.

The `loadActionDefaults()` method is commonly used in the form initialization
function.

Specify the action defaults filename without the .4ad extension.

The resource file is searched in several directories in a given order, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

If the form already contains action defaults, it will be replaced by the new action defaults
loaded by this method.

The `loadActionDefaults()` method of a form object is typically used in a [generic form initializer function](3146-ui-form-setdefaultinitializerfunction.md "Define the default initializer for all forms.").

## Related links

**Related concepts**  

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
