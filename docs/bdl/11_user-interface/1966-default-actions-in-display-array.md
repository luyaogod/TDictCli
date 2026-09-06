---
title: "Default actions in DISPLAY ARRAY"
source: "fgl-topics/c_fgl_DisplayArray_017.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > Default actions in DISPLAY ARRAY"
type: "concept"
description: "When a DISPLAY ARRAY instruction executes, the runtime system creates a set of default actions . Field validation occurs and different DISPLAY ARRAY control blocks are executed, based on the invoked ..."
---

# Default actions in DISPLAY ARRAY

When a `DISPLAY ARRAY` instruction executes, the runtime system creates a set of
[default actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.").

Field validation occurs and different `DISPLAY ARRAY` control blocks are executed,
based on the invoked default action.

This table lists the default actions created for this dialog:

| Default action | Description |
| --- | --- |
| `accept` | Validates the `DISPLAY ARRAY` dialog (validates current row selection)Creation can be avoided with `ACCEPT = FALSE` attribute. |
| `cancel` | Cancels the `DISPLAY ARRAY` dialog ([`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") is set to `TRUE`)Creation can be avoided with `CANCEL = FALSE` attribute. |
| `close` | By default, cancels the `DISPLAY ARRAY` dialog (no validation, `int_flag` is set to `TRUE`)Default action view is hidden. See [Implementing the close action](2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button)."). |
| `help` | Shows the help topic defined by the `HELP` clause.Only created when a `HELP` clause is defined. |
| `nextrow` | Moves to the next row in a list displayed in one row of fields.See note (1). |
| `prevrow` | Moves to the previous row in a list displayed in one row of fields.See note (1). |
| `firstrow` | Moves to the first row in a list displayed in one row of fields.See note (1). |
| `lastrow` | Moves to the last row in a list displayed in one row of fields.See note (1). |
| `find` | Opens the fglfind dialog window to let the user enter a search value, and searches for the row matching the value.See note (2). |
| `findnext` | Searches for the next row matching the value entered during the fglfind dialog.See note (2). |

Notes:

1. The action is only created with a `DISPLAY ARRAY` or `INPUT ARRAY`
   using a [screen record](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") bound to a set of
   form fields in a `GRID` container, and this set of fields show only a single row of
   the list. The action is not created when using a screen array bound to a list container such as
   `TABLE`, `TREE` and `SCROLLGRID`, or to a set of fields
   in a `GRID` container, that repeat on several lines to show more than one single
   row.
2. The action is only created if the context allows [built-in find](2325-finding-rows-matching-a-pattern.md "List controllers implement a built-in find. This feature can be disabled if not required.").

The accept and cancel default actions can be avoided with the `ACCEPT` and
`CANCEL` dialog control
attributes:

```
DISPLAY ARRAY arr TO sr.* ATTRIBUTES( CANCEL=FALSE, ... )
  ...
```

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
