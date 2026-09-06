---
title: "Loading .4tm topmenu definition files"
source: "fgl-topics/c_fgl_topmenus_008.html"
breadcrumb: "User interface > Form definitions > Topmenus > Loading .4tm topmenu definition files"
type: "concept"
---

# Loading .4tm topmenu definition files

> Topmenu XML definition files can be loaded at runtime.

## Loading an XML topmenu file to define a default/global topmenu

To load a .4tm topmenu definition file as default/global topmenu for all
forms, use the [`ui.Interface.loadTopmenu()`](../15_library-reference/3116-ui-interface-loadtopmenu.md "Load a default/global topmenu file for all forms of the program.") class method:

```
CALL ui.Interface.loadTopmenu("standard")
```

The purpose of the default/global topmenu will be displayed in all forms.

## Loading a XML topmenu file for the current form/window

To load a .4tm topmenu definition file for a given form, use the [`ui.Form.loadTopmenu()`](../15_library-reference/3154-ui-form-loadtopmenu.md "Load the form topmenu.")
method:

```
DEFINE myform ui.Form
...
CALL myform.loadTopmenu("standard")
```

The topmenu will be displayed in that form only.

This method is typically used in [form
initializers](../15_library-reference/3146-ui-form-setdefaultinitializerfunction.md "Define the default initializer for all forms.").

## Related links

**Related concepts**  

[Syntax of a topmenu file (.4tm)](1869-syntax-of-a-topmenu-file-4tm.md "A .4tm topmenu file is an XML file that holds a tree of elements defining a topmenu structure.")
