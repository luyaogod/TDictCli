---
title: "AFTER CONSTRUCT block"
source: "fgl-topics/c_fgl_dialog_AFTER_CONSTRUCT_2.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control blocks > AFTER CONSTRUCT block"
type: "concept"
description: "Syntax AFTER CONSTRUCT instruction [...] AFTER CONSTRUCT block in singular and parallel CONSTRUCT dialogs In a singular CONSTRUCT instruction, the AFTER CONSTRUCT is only executed once when dialog is ..."
---

# AFTER CONSTRUCT block

## Syntax

```
AFTER CONSTRUCT
   instruction [...]
```

## AFTER CONSTRUCT block in singular and parallel CONSTRUCT dialogs

In a singular `CONSTRUCT` instruction, the `AFTER CONSTRUCT` is
only executed once when dialog is ended.

Use an `AFTER CONSTRUCT` block to execute instructions after the user
has finished search criteria input.

`AFTER CONSTRUCT` is not executed if an `EXIT CONSTRUCT` is
performed.

Use the `AFTER CONSTRUCT` block to check user input and and stay in the dialog
with `NEXT FIELD` or `CONTINUE CONSTRUCT`.

Before checking the content of the fields used in the `CONSTRUCT`, make sure that
the [`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") variable is
`FALSE`. In the case that the user cancels the dialog, the validation rules must be
skipped.

Since no program variables are associated with the form fields, you must query the input
buffers of the fields to get the values entered by the user.

```
LET int_flag = FALSE
CONSTRUCT BY NAME where_part ON ...
  ...
  AFTER CONSTRUCT
    IF NOT int_flag THEN
       IF length(DIALOG.getFieldBuffer(cust_name))==0
       OR length(DIALOG.getFieldBuffer(cust_addr))==0 THEN
          ERROR "Enter a search criteria for customer name and address fields."
          NEXT FIELD CURRENT
       END IF
    END IF
END CONSTRUCT
```

To limit the validation to fields that have been modified by the end user, you can call the
`FIELD_TOUCHED()` function or the `DIALOG.getFieldTouched()` method to
check if a field has changed during the dialog execution. This makes your validation code execute
faster if the user has only modified a couple of fields in a large form.

## AFTER CONSTRUCT block in CONSTRUCT of procedural DIALOG

In a `CONSTRUCT` sub-dialog of a procedural `DIALOG`
instruction, the `AFTER CONSTRUCT` block is executed when the focus is lost by a
group of fields driven by a `CONSTRUCT` sub-dialog. This trigger is invoked if a
field of the sub-dialog loses the focus, and a field of a different sub-dialog gets the focus.

If the focus leaves the current group and goes to an action view, this trigger is not
executed, because the focus did not go to another sub-dialog yet.

`AFTER CONSTRUCT` is executed after the `AFTER FIELD`
and before the `AFTER DIALOG` block.

Executing a `NEXT FIELD` in the `AFTER CONSTRUCT` control
block will keep the focus in the group of fields.

In this example, the `AFTER CONSTRUCT` block is used to build the SELECT
statement:

```
CONSTRUCT BY NAME sql ON customer.*
   AFTER CONSTRUCT
       LET sql = "SELECT * FROM customers WHERE " || sql
```

## Related links

**Related concepts**  

[AFTER DISPLAY block](1974-after-display-block.md "AFTER DISPLAY block")

[AFTER INPUT block](1941-after-input-block.md "AFTER INPUT block")

[BEFORE CONSTRUCT block](2057-before-construct-block.md "BEFORE CONSTRUCT block")
