---
title: "ON INSERT block"
source: "fgl-topics/c_fgl_dialog_ON_INSERT_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG interaction blocks > ON INSERT block"
type: "concept"
description: "Syntax ON INSERT instruction [...] Usage Similar to the ON APPEND control block, the ON INSERT trigger can be used to enable row creation during a DISPLAY ARRAY dialog. If this block is defined, the ..."
---

# ON INSERT block

## Syntax

```
ON INSERT
   instruction [...]
```

## Usage

Similar to the `ON APPEND` control
block, the `ON INSERT` trigger can be used to enable row creation during a
`DISPLAY ARRAY` dialog. If this block is defined, the dialog will automatically
create the insert action. This action can be decorated, enabled and disabled as a regular
action.

If the dialog defines an `ON ACTION insert` interaction block and the `ON
INSERT` block is used, the compiler will stop with error [-8408](../15_library-reference/4483-genero-bdl-errors.md).

When the user fires the insert action, the dialog first executes the user code of the
`AFTER ROW` block if defined. Then the new row is created: The insert action creates
a new row before current row in the list. After creating the row, the dialog executes the user code
of the `ON INSERT` block.

The dialog handles only row creation actions and navigation, you must program the record input
with a regular `INPUT` statement, to let the end user enter data for the
newly-created row. This is typically done with an `INPUT` binding explicitly array
fields to the screen record fields. The new current row in the program array is identified with
`arr_curr()`, and the
current screen line in the form is defined by `scr_line()`:

```
DISPLAY ARRAY arr TO sr.*
 ...
  ON INSERT
    INPUT arr[arr_curr()].* FROM sr[scr_line()].* ;
    ...
```

Pay attention to the semicolon ending the `INPUT` instruction, which is usually
needed here to solve a language grammar conflict when nested dialog instructions are
implemented.

After the user code is executed, the dialog gets the control back and processes the new row as
follows:

- If the `int_flag` global
  variable is `FALSE` and `status` is zero, the new row is kept in the program array, and the `BEFORE ROW` block is executed for the new created
  row.
- If the `int_flag` global variable is `TRUE` or
  `status` is different from zero, the new row is removed from the program array, and
  the `BEFORE ROW` block is executed for the row that was existing at the current
  position, before the new row was created.

The `DISPLAY ARRAY` dialog always resets `int_flag` to
`FALSE` and `status` to zero before executing the user code of the
`ON INSERT` block.

The insert action is disabled if the maximum number of rows is reached.

If needed, the `ON INSERT` handler can be configured with action attributes by
added an `ATTRIBUTES()` clause, as with user-defined action handlers:

```
  ON INSERT ATTRIBUTES(TEXT=%"custlist.delete", IMAGE="listdel")
```

## Related links

**Related concepts**  

[Record input (INPUT)](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.")

[ON UPDATE block](1986-on-update-block.md "ON UPDATE block")

[ON DELETE block](1987-on-delete-block.md "ON DELETE block")
