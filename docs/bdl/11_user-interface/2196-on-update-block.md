---
title: "ON UPDATE block"
source: "fgl-topics/c_fgl_dialog_ON_UPDATE_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG interaction blocks > ON UPDATE block"
type: "concept"
description: "Syntax ON UPDATE instruction [...] Usage The ON UPDATE trigger can be used to enable row modification during a DISPLAY ARRAY dialog. If this block is defined, the dialog will automatically create the ..."
---

# ON UPDATE block

## Syntax

```
ON UPDATE
   instruction [...]
```

## Usage

The `ON UPDATE` trigger can be used to enable row modification during a
`DISPLAY ARRAY` dialog. If this block is defined, the dialog will automatically
create the update action. This action can be decorated, enabled and disabled as regular actions.

You typically configure the `TABLE` container in the form by defining the `DOUBLECLICK` attribute to "update", in
order to trigger the update action when the user double-clicks on a row.

The action created for the `ON UPDATE` modification trigger is implicitley marked
as a [`ROWBOUND` action](2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list.").

If the dialog defines an `ON ACTION update` interaction block and the `ON
UPDATE` block is used, the compiler will stop with error [-8408](../15_library-reference/4483-genero-bdl-errors.md).

When the user fires the update action, the dialog executes the user code of the `ON
UPDATE` block.

The dialog handles only the row modification action and navigation, you must program the record
input with a regular `INPUT` statement, to let the end user modify the data of the
current row. This is typically done with an `INPUT` binding explicitly array fields
to the screen record fields, with the `WITHOUT DEFAULTS` clause. The current row in
the program array is identified with `arr_curr()`, and the current screen line in the form is defined by `scr_line()`:

```
DISPLAY ARRAY arr TO sr.*
  ...
  ON UPDATE
    INPUT arr[arr_curr()].* WITHOUT DEFAULTS FROM sr[scr_line()].* ;
    ...
```

Pay attention to the semicolon ending the `INPUT` instruction, which is usually
needed here to solve a language grammar conflict when nested dialog instructions are
implemented.

After the user code is executed, the dialog gets the control back and processes the current row
as follows:

- If the `int_flag` global
  variable is `FALSE` and `status` is zero, the modified values of the current row are kept in the program
  array.
- If the `int_flag` global variable is `TRUE` or
  `status` is different from zero, the old values of the current row are restored in
  the program array.

The `DISPLAY ARRAY` dialog always resets `int_flag` to
`FALSE` and `status` to zero before executing the user code of the
`ON UPDATE` block.

If needed, the `ON UPDATE` handler can be configured with action attributes by
added an `ATTRIBUTES()` clause, as with user-defined action
handlers:

```
  ON UPDATE ATTRIBUTES(TEXT=%"custlist.delete", IMAGE="listdel")
```

## Related links

**Related concepts**  

[Record input (INPUT)](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.")

[ON INSERT block](1985-on-insert-block.md "ON INSERT block")

[ON APPEND block](1984-on-append-block.md "ON APPEND block")

[ON DELETE block](1987-on-delete-block.md "ON DELETE block")
