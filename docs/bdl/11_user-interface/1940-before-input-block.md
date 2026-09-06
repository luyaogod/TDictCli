---
title: "BEFORE INPUT block"
source: "fgl-topics/c_fgl_dialog_BEFORE_INPUT.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Using simple record inputs > INPUT control blocks > BEFORE INPUT block"
type: "concept"
description: "Syntax BEFORE INPUT instruction [...] BEFORE INPUT block in singular INPUT, INPUT ARRAY dialogs In a singular INPUT or INPUT ARRAY instruction, the BEFORE INPUT is only executed once when the dialog ..."
---

# BEFORE INPUT block

## Syntax

```
BEFORE INPUT
   instruction [...]
```

## BEFORE INPUT block in singular INPUT, INPUT ARRAY dialogs

In a singular `INPUT` or `INPUT ARRAY` instruction, the
`BEFORE INPUT` is only executed once when the dialog is started.

The `BEFORE INPUT` block is executed once at dialog startup, before the
runtime system gives control to the user. This block can be used to display messages to the
user, initialize program variables and set up the dialog instance by deactivating unused
fields or actions the user is not allowed to execute.

```
INPUT BY NAME cust_rec.* ...
   BEFORE INPUT
       MESSAGE "Input customer information"
       CALL DIALOG.setActionActive("check_info", is_super_user() )
       CALL DIALOG.setFieldActive("cust_comment", is_super_user() )
   ...
```

The fields are initialized with the default values, before the `BEFORE
INPUT` code is executed. When the `INPUT` instruction uses the
`WITHOUT DEFAULTS` option or when inserting/appending a new row in an `INPUT
ARRAY`, the default values are taken from the program variables bound to the fields.
Otherwise (with defaults), the [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") attributes of the form fields are used.

Use the `NEXT FIELD` control instruction in the `BEFORE
INPUT` block, to jump to a specific field when the dialog starts.

## BEFORE INPUT block in INPUT and INPUT ARRAY of procedural DIALOG

In an `INPUT` or `INPUT ARRAY` sub-dialog of a procedural
`DIALOG` instruction, the `BEFORE INPUT` block is executed when the
focus goes to a group of fields driven by the sub-dialog. This trigger is only invoked if a field of
the sub-dialog gets the focus, and none of the other fields had the focus.

When the focus is in a list driven by an `INPUT ARRAY` sub-dialog, moving to
a different row will not invoke the `BEFORE INPUT` block.

`BEFORE INPUT` is executed after the `BEFORE DIALOG`
block andbefore the `BEFORE ROW`, `BEFORE FIELD` blocks.

In this example, the `BEFORE INPUT` block is used to set up a specific action
and display a message:

```
INPUT BY NAME p_order.*
   BEFORE INPUT
       CALL DIALOG.setActionActive("validate_order", TRUE)
```

With a `DIALOG` block, fields are initialized with the default values,
before the `BEFORE DIALOG` code is executed. Fields are not re-initialized
when re-entering the `INPUT` or `INPUT ARRAY` sub-dialog. Regarding
`WITHOUT DEFAULTS` clause, the same rules apply as in a singular
`INPUT` or `INPUT ARRAY`.

## Related links

**Related concepts**  

[BEFORE CONSTRUCT block](2057-before-construct-block.md "BEFORE CONSTRUCT block")

[BEFORE DISPLAY block](1973-before-display-block.md "BEFORE DISPLAY block")

[AFTER INPUT block](1941-after-input-block.md "AFTER INPUT block")
