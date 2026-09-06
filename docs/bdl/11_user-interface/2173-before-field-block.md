---
title: "BEFORE FIELD block"
source: "fgl-topics/c_fgl_dialog_BEFORE_FIELD_6.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > DIALOG control blocks > BEFORE FIELD block"
type: "concept"
description: "Syntax BEFORE FIELD field-spec [ , ...] instruction [...] Usage In dialog instructions INPUT , INPUT ARRAY , CONSTRUCT or in a DISPLAY ARRAY using the FOCUSONFIELD attribute, the BEFORE FIELD block is ..."
---

# BEFORE FIELD block

## Syntax

```
BEFORE FIELD field-spec [,...]
   instruction [...]
```

## Usage

In dialog instructions `INPUT`, `INPUT ARRAY`,
`CONSTRUCT` or in a `DISPLAY ARRAY` using the
`FOCUSONFIELD` attribute, the `BEFORE FIELD` block is executed every
time the specified field gets the focus.

For single record inputs driven by `INPUT` or query by example (QBEs) driven by
`CONSTRUCT`, the `BEFORE FIELD` block is executed when moving the
focus from field to field.

For editable lists driven by `INPUT ARRAY`, the `BEFORE FIELD`
block is executed when moving the focus from field to field in the same row, or when moving to
another row in the same column.

For record lists driven by `DISPLAY ARRAY` using the `FOCUSONFIELD`
attribute, the `BEFORE FIELD` block is executed when moving the focus from field to
field. However, the fields will not be editable as in an `INPUT ARRAY`.

The `BEFORE FIELD` block is also executed when performing a `NEXT FIELD` instruction.

The `BEFORE FIELD` keywords must be followed by a list of form field
specification. The screen-record name can be omitted.

`BEFORE FIELD` is executed after `BEFORE INPUT`,
`BEFORE CONSTRUCT`, `BEFORE ROW` and `BEFORE
INSERT`.

Use this block to do some field value initialization, or to display a message to the user:

```
INPUT BY NAME p_cust.* ...
   BEFORE FIELD cust_status
     LET p_cust.cust_comment = NULL
     MESSAGE "Enter customer status"
```

When using the default `FIELD ORDER CONSTRAINT` mode, the dialog executes the
`BEFORE FIELD` block of the field corresponding to the first variable of an
`INPUT` or `INPUT ARRAY`, even if that field is not editable
(`NOENTRY`, hidden or disabled). The block is executed when you enter the dialog and
every time you create a new row in the case of `INPUT ARRAY`. This behavior is
supported for backward compatibility. The block is not executed when using the `FIELD
ORDER FORM`, the mode recommended for `DIALOG` instructions.

With the `FIELD ORDER FORM` mode, for each dialog executing for the first time
with a specific form, the `BEFORE FIELD` block will be invoked for the first field of
the initial tabbing list defined by the form, even if that field was hidden or moved around in a
table. The dialog then behaves as if a `NEXT FIELD first-visible-column` had been
executed in the `BEFORE FIELD` of that field.

When form-level validation occurs and a field contains an invalid value, the dialog gives the
focus to the field, but no `BEFORE FIELD` trigger will be executed.

## Related links

**Related concepts**  

[Form-level validation rules](2240-form-level-validation-rules.md "Form-level validation rules can be defined for each field controlled by a dialog.")

[AFTER FIELD block](1944-after-field-block.md "AFTER FIELD block")
