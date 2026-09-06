---
title: "ON DELETE block"
source: "fgl-topics/c_fgl_dialog_ON_DELETE_3.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG interaction blocks > ON DELETE block"
type: "concept"
description: "Syntax ON DELETE instruction [...] Usage The ON DELETE trigger can be used to enable row deletion during a DISPLAY ARRAY dialog. If this block is defined, the dialog will automatically create the ..."
---

# ON DELETE block

## Syntax

```
ON DELETE
   instruction [...]
```

## Usage

The `ON DELETE` trigger can be used to enable row deletion during a
`DISPLAY ARRAY` dialog. If this block is defined, the dialog will automatically
create the delete action. This action can be decorated, enabled and disabled as regular actions.

The action created for the `ON DELETE` modification trigger is implicitley marked
as a [`ROWBOUND` action](2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list.").

If the dialog defines an `ON ACTION delete` interaction block and the `ON
DELETE` block is used, the compiler will stop with error [-8408](../15_library-reference/4483-genero-bdl-errors.md).

When the user fires the delete action, the dialog executes the user code of the `ON
DELETE` block.

The dialog handles only the row deletion action and navigation, you can typically program a
validation dialog box to let the user confirm the deletion. The current row in the program array is
identified with `arr_curr()`:

```
DISPLAY ARRAY arr TO sr.*
  ...
  ON DELETE
    IF fgl_winQuestion("Delete",
          "Do you want to delete this record?",
          "yes", "no|yes", "help", 0) == "no"
    THEN
       LET int_flag = TRUE
    END IF
    ...
```

After the user code is executed, the dialog gets the control back and processes the current row
as follows:

- If the `int_flag` global
  variable is `FALSE` and `status` is zero, the current row is deleted from the program array, and the
  `BEFORE ROW` block is executed for the next row in the list.
- If the `int_flag` global variable is `TRUE` or
  `status` is different from zero, the current row is kept in the program array, and
  the `BEFORE ROW` block is executed again for the current row.

The `DISPLAY ARRAY` dialog always resets `int_flag` to
`FALSE` and `status` to zero before executing the user code of the
`ON DELETE` block.

If needed, the `ON DELETE` handler can be configured with action attributes by
adding an `ATTRIBUTES()` clause, as with user-defined action
handlers:

```
  ON DELETE ATTRIBUTES(TEXT=%"custlist.delete", IMAGE="listdel")
```

## Related links

**Related concepts**  

[ON APPEND block](1984-on-append-block.md "ON APPEND block")

[ON INSERT block](1985-on-insert-block.md "ON INSERT block")

[ON UPDATE block](1986-on-update-block.md "ON UPDATE block")
