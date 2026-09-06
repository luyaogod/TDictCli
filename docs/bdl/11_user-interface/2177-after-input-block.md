---
title: "AFTER INPUT block"
source: "fgl-topics/c_fgl_dialog_AFTER_INPUT_4.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control blocks > AFTER INPUT block"
type: "concept"
description: "Syntax AFTER INPUT instruction [...] AFTER INPUT block in singular INPUT, INPUT ARRAY dialogs In a singular INPUT or INPUT ARRAY instruction, the AFTER INPUT is only executed once when dialog ends. ..."
---

# AFTER INPUT block

## Syntax

```
AFTER INPUT
   instruction [...]
```

## AFTER INPUT block in singular INPUT, INPUT ARRAY dialogs

In a singular `INPUT` or `INPUT ARRAY` instruction, the
`AFTER INPUT` is only executed once when dialog ends.

The `AFTER INPUT` block is executed after the user has validated or canceled the
`INPUT` or `INPUT ARRAY` dialog with the accept or cancel default
actions, or when the `ACCEPT INPUT` instruction is executed.

The `AFTER INPUT` block is not executed when the `EXIT INPUT`
instruction is performed.

In singular and parallel dialogs, this block is typically used to implement global dialog
validation rules for several fields. If the values entered by the user do not satisfy the
constraints, use the `NEXT FIELD` instruction to force the dialog to continue. The
`CONTINUE INPUT` instruction can be used instead of `NEXT FIELD`, when
no particular field has to be selected.

Before checking the validation rules, make sure that the [`int_flag`](../09_advanced-features/0940-int-flag.md "int_flag is a predefined variable set to TRUE when an interruption event is detected or when a cancel action is fired in a singular dialog.") variable is
`FALSE`: because if the user cancels the dialog, the validation rules must be
skipped.

```
LET int_flag = FALSE
INPUT BY NAME cust_rec.*
   WITHOUT DEFAULTS ATTRIBUTES ( UNBUFFERED )
  ...

  AFTER INPUT
    IF NOT int_flag THEN
       IF cust_rec.cust_address IS NOT NULL
          AND cust_rec.cust_zipcode IS NULL THEN
          ERROR "Address is incomplete, enter a zipcode."
          NEXT FIELD zipcode
       END IF
    END IF
END INPUT
```

To limit the validation to fields that have been modified by the end user, you can call the
`FIELD_TOUCHED()` function or the `DIALOG.getFieldTouched()` method to
check if a field has changed during the dialog execution. This will make your validation code faster
if the user has only modified a couple of fields in a large form.

## AFTER INPUT block in INPUT and INPUT ARRAY of procedural DIALOG

In an `INPUT` or `INPUT ARRAY` sub-dialog of a procedural
`DIALOG` instruction, the `AFTER INPUT` block is executed when the
focus is lost by a group of fields driven by an `INPUT` or `INPUT
ARRAY` sub-dialog. This trigger is invoked if a field of the sub-dialog loses the focus, and
a field of a different sub-dialog gets the focus. When the focus is in a list driven by an
`INPUT ARRAY` sub-dialog, moving to a different row will not invoke the `AFTER
INPUT` block.

If the focus leaves the current group and goes to an action view, this trigger is not executed,
because the focus has not gone to another sub-dialog yet.

`AFTER INPUT` is executed after the `AFTER FIELD`,
`AFTER ROW` blocks and before the `AFTER DIALOG` block.

Executing a `NEXT FIELD` in the `AFTER INPUT` control block will
keep the focus in the group of fields. Within an `INPUT ARRAY` sub-dialog,
`NEXT FIELD` will keep the focus in the list and stay in the current row. You
typically use this behavior to control user input.

In this example, the `AFTER INPUT` block is used to validate data and disable an
action that can only be used in the current group:

```
INPUT BY NAME p_order.*
   AFTER INPUT
      IF NOT check_order_data(DIALOG) THEN
        NEXT FIELD CURRENT
      END IF
      CALL DIALOG.setFieldActive("validate_order", FALSE)
```

## Related links

**Related concepts**  

[AFTER DISPLAY block](1974-after-display-block.md "AFTER DISPLAY block")

[AFTER CONSTRUCT block](2058-after-construct-block.md "AFTER CONSTRUCT block")

[BEFORE INPUT block](1940-before-input-block.md "BEFORE INPUT block")
