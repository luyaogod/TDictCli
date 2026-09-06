---
title: "Loading .4tb toolbar definition files"
source: "fgl-topics/c_fgl_toolbars_007.html"
breadcrumb: "User interface > Form definitions > Toolbars > Loading .4tb toolbar definition files"
type: "concept"
---

# Loading .4tb toolbar definition files

> Toolbar XML definition files can be loaded at runtime.

## Loading an XML toolbar file to define a default/global toolbar

To load a .4tb toolbar definition file as default toolbar for all forms, use
the [`ui.Interface.loadToolbar()`](../15_library-reference/3115-ui-interface-loadtoolbar.md "Load a default/global toolbar file for all forms of the program.") class method:

```
CALL ui.Interface.loadToolbar("standard")
```

The purpose of the default/global toolbar will be displayed in all forms.

## Loading a XML toolbar file for the current form/window

To load a .4tb toolbar definition file for a given form, use the [`ui.Form.loadToolbar()`](../15_library-reference/3153-ui-form-loadtoolbar.md "Load the form toolbar.")
method:

```
DEFINE myform ui.Form
...
CALL myform.loadToolbar("standard")
```

The toolbar will be displayed in that form only.

This method is typically used in [form initializers](../15_library-reference/3146-ui-form-setdefaultinitializerfunction.md "Define the default initializer for all forms.").

## Related links

**Related concepts**  

[Syntax of a toolbar file (.4tb)](1858-syntax-of-a-toolbar-file-4tb.md "A .4tb toolbar file is an XML file that holds a tree of elements defining a toolbar structure.")
