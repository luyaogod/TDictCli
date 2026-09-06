---
title: "Default actions created by a DIALOG block"
source: "fgl-topics/c_fgl_DIALOG_block_default_actions.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > Default actions created by a DIALOG block"
type: "concept"
---

# Default actions created by a DIALOG block

> Default actions ease the implementation of the controller by providing expected actions.

The runtime system creates a set of [default
actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.") based on the sub-dialogs defined in a (declarative or procedural)
`DIALOG` block. These actions are provided to ease the implementation of the
controller. For example, when using an `INPUT ARRAY` sub-dialog, the dialog
instruction will automatically create the insert, append and delete default actions.

Table 1 lists the default actions created for the `DIALOG` interactive instruction, for
each type of sub-dialogs:

| Default action | Control Block execution order |
| --- | --- |
| `help` | Shows the help topic defined by the `HELP` clause.Only created when a `HELP` clause or option is defined for the sub-dialog. |
| `insert` | Inserts a new row before current row.Only for `INPUT ARRAY` dialogs. Action creation can be avoided with `INSERT ROW = FALSE` attribute. |
| `append` | Appends a new row at the end of the list.Only for `INPUT ARRAY` dialogs. Action creation can be avoided with `APPEND ROW = FALSE` attribute. |
| `delete` | Deletes the current row.Only for `INPUT ARRAY` dialogs. Action creation can be avoided with `DELETE ROW = FALSE` attribute. |
| `nextrow` | Moves to the next row in a list displayed in one row of fields.See note (1). |
| `prevrow` | Moves to the previous row in a list displayed in one row of fields.See note (1). |
| `firstrow` | Moves to the first row in a list displayed in one row of fields.See note (1). |
| `lastrow` | Moves to the last row in a list displayed in one row of fields.See note (1). |
| `find` | Opens the fglfind dialog window to let the user enter a search value, and seeks to the row matching the value.See note (2). |
| `findnext` | Seeks to the next row matching the value entered during the fglfind dialog.See note (2). |

Notes:

1. The action is only created with a `DISPLAY ARRAY` or `INPUT ARRAY`
   using a [screen record](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") bound to a set of
   form fields in a `GRID` container, and this set of fields show only a single row of
   the list. The action is not created when using a screen array bound to a list container such as
   `TABLE`, `TREE` and `SCROLLGRID`, or to a set of fields
   in a `GRID` container, that repeat on several lines to show more than one single
   row.
2. The action is only created if the context allows [built-in find](2325-finding-rows-matching-a-pattern.md "List controllers implement a built-in find. This feature can be disabled if not required.").

The `insert`, `append` and `delete` default actions
can be avoided with dialog control
attributes:

```
INPUT ARRAY arr TO sr.* ATTRIBUTES( INSERT ROW=FALSE, APPEND ROW=FALSE, ... )
     ...
```

## Related links

**Related concepts**  

[DISPLAY ARRAY ATTRIBUTES clause](2091-display-array-attributes-clause.md "DISPLAY ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.")

[INPUT ARRAY ATTRIBUTES clause](2092-input-array-attributes-clause.md "INPUT ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.")
