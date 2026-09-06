---
title: "Understanding declarative dialogs"
source: "fgl-topics/c_fgl_declarative_dialogs_intro.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Understanding declarative dialogs"
type: "concept"
---

# Understanding declarative dialogs

> Declarative dialogs are defined with DIALOG blocks at the module level.

A declarative dialog block is a module element defined at the same level as a
`FUNCTION` or `REPORT` routine:

```
-- Module orders.4gl
SCHEMA stock
PUBLIC TYPE t_order RECORD LIKE custorder.*
PUBLIC TYPE t_orders DYNAMIC ARRAY OF RECORD LIKE custorder.*
DIALOG orders_subdialog(ordlist t_orders)
  DEFINE x INT
  DISPLAY ARRAY ordlist TO sr_orders.*
     ...
  END DISPLAY 
END DIALOG
```

The name of a declarative dialog is mandatory. A [`SUBDIALOG`](2087-the-subdialog-clause.md) clause will reference
a declarative dialog by its name, and can identify sub-dialog actions with a prefix. For more
details, read [Sub-dialog actions in procedural DIALOG blocks](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.").

The declarative dialog can take parameters, typically to bind program variables defined in the
module of the parent dialog to the interactive instruction of the sub-dialog. Declarative dialog
parameters will be passed when the parent dialog starts. Variables of type `RECORD`
must be passed by reference with the [`INOUT`](../09_advanced-features/0835-passing-records-as-parameter.md) keyword. Variables of type `DYNAMIC ARRAY` are
implicitely passed by reference.

When using the `DIALOG` keyword inside a declarative dialog block to use
`ui.Dialog` class methods, it references the current instance of the dialog
object.

## Related links

**Related concepts**  

[Windows and forms](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.")
