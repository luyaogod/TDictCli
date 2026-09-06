---
title: "ON APPEND block"
source: "fgl-topics/c_fgl_dialog_ON_APPEND_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG interaction blocks > ON APPEND block"
type: "concept"
description: "Syntax ON APPEND instruction [...] Usage Similar to the ON INSERT control block, the ON APPEND trigger can be used to enable row creation during a DISPLAY ARRAY dialog. If this block is defined, the ..."
---

# ON APPEND block

## Syntax

```
ON APPEND
   instruction [...]
```

## Usage

Similar to the `ON INSERT` control
block, the `ON APPEND` trigger can be used to enable row creation during a
`DISPLAY ARRAY` dialog. If this block is defined, the dialog will automatically
create the append action. This action can be decorated, enabled and disabled as a regular
action.

If the dialog defines an `ON ACTION append` interaction block and the `ON
APPEND` block is used, the compiler will stop with error [-8408](../15_library-reference/4483-genero-bdl-errors.md).

When the user fires the append action, the dialog first executes the user code of the
`AFTER ROW` block if defined. Then the dialog moves to the end of the list, and
creates a new row after the last existing row. After creating the row, the dialog executes the user
code of the `ON APPEND` block.

The dialog handles only row creation actions and navigation, you must program the record input
with a regular `INPUT` statement, to let the end user enter data for the
newly-created row. This is typically done with an `INPUT` binding explicitly array
fields to the screen record fields. The new current row in the program array is identified with
`arr_curr()`, and the
current screen line in the form is defined by `SCR_LINE()`:

```
DISPLAY ARRAY arr TO sr.*
  ...
  ON APPEND
    INPUT arr[arr_curr()].* FROM sr[scr_line()].* ;
    ...
```

Pay attention to the semicolon ending the `INPUT` instruction, which is usually
needed here to solve a language grammar conflict when nested dialog instructions are
implemented.

After the user code is executed, the dialog gets the control back and processes the new row as
follows:

- If the `int_flag` global
  variable is `FALSE` and `status` is zero, the new row is kept in the program array, and the [`BEFORE ROW`](1975-before-row-block.md) block is executed for the
  newly-created row.
- If the `int_flag` global variable is `TRUE` or
  `status` is different from zero, the new row is removed from the program array, and
  the `BEFORE ROW` block is executed for the row that existed at the current position,
  before the new row was created.

The `DISPLAY ARRAY` dialog always resets `int_flag` to
`FALSE` and `status` to zero before executing the user code of the
`ON APPEND` block.

The append action is disabled if the maximum number of rows is reached.

If needed, the `ON APPEND` handler can be configured with action attributes by
added an `ATTRIBUTES()` clause, as with user-defined action
handlers:

```
  ON APPEND ATTRIBUTES(TEXT=%"custlist.delete", IMAGE="listdel")
```

## Related links

**Related concepts**  

[Record input (INPUT)](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.")

[ON UPDATE block](1986-on-update-block.md "ON UPDATE block")

[ON DELETE block](1987-on-delete-block.md "ON DELETE block")
