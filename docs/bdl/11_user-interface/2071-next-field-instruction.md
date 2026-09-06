---
title: "NEXT FIELD instruction"
source: "fgl-topics/c_fgl_dialog_NEXT_FIELD_4.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > CONSTRUCT control instructions > NEXT FIELD instruction"
type: "concept"
description: "Syntax NEXT FIELD { CURRENT | NEXT | PREVIOUS | field-spec } where field-spec can be one of: field-name : The name of the field as defined in the form file. table-name.field-name : The name of SQL ..."
---

# NEXT FIELD instruction

## Syntax

```
NEXT FIELD
   { CURRENT
   | NEXT
   | PREVIOUS
   | field-spec
   }
```

where field-spec can be one of:

- field-name: The name of the field as defined in the form file.
- table-name.field-name: The name of SQL table and field name as defined in the
  form file.
- screen-record-name.field-name: The screen record name and field name as
  defined in the form file.
- `FORMONLY.field-name`: The FORMONLY keyword
  followed by the field name as defined in the form file.

## Purpose of NEXT FIELD instruction

The `NEXT FIELD` instruction gives the focus to the specified field. This
instruction can be used to control field input, in [`BEFORE FIELD`](1942-before-field-block.md), [`ON
CHANGE`](1943-on-change-block.md) or [`AFTER
FIELD`](1944-after-field-block.md) blocks, it can also force a `DISPLAY ARRAY` or `INPUT
ARRAY` to stay in the current row when `NEXT FIELD` is used in the [`AFTER ROW`](1976-after-row-block.md) block.

```
  AFTER FIELD cust_status
    IF cust_rec.cust_status != "A" AND cust_rec.cust_creadate < TODAY THEN
      ERROR "Customer created today must have status A"
      NEXT FIELD cust_status
    END IF
```

If it exists, the `BEFORE FIELD` block of the destination field is executed.
However, in the context of a singular dialog, or when in a `DIALOG` block and the
next field is in the same sub-dialog as the current field, the `AFTER FIELD` of the
current field is not executed.

In `INPUT`, `INPUT ARRAY` and `CONSTRUCT` dialogs,
the purpose of the `NEXT FIELD` instruction is to give the focus to an focusable
field. Make sure that the field specified in `NEXT FIELD` is active and focusable, or
use `NEXT FIELD CURRENT`.

In a `DISPLAY ARRAY` using the [`FOCUSONFIELD`](1965-display-array-instruction-configuration.md) attribute, `NEXT FIELD` can be used in
conjunction with [`DIALOG.setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list."), to set the focus to a specific cell in the list.

Instead of the `NEXT FIELD` instruction, you can use the [`DIALOG.nextField("field-name")`](../15_library-reference/3210-ui-dialog-nextfield.md "Registers the next field to go to.") method to register a field, for
example when the name is not known at compile time. However, this method only registers the field.
It does not stop code execution, like the `NEXT FIELD` instruction does. You must
execute a `CONTINUE DIALOG` to get the same behavior as `NEXT
FIELD`.

## Form field identification with NEXT FIELD

With the `NEXT FIELD form-field` instruction, fields are
identified by the [form field](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.") name
specification, not the program variable name used by the dialog.

Abstract field identification is also supported with the `CURRENT`,
`NEXT` and `PREVIOUS` keywords.

The field-name must be a field name as defined in the form file, and must be
used by the current dialog.

Form fields are [bound](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.") to program
variables with the binding clause of the dialog instruction (`INPUT
variable-list FROM field-list`, `INPUT BY NAME
variable-list`, `CONSTRUCT BY NAME sql ON
column-list`, `CONSTRUCT sql ON
column-list FROM field-list`, `INPUT ARRAY
array-name FROM screen-array.*`).

Here are some examples of field name specifications:

- `cust_name`
- `customer.cust_name`
- `cust_screen_record.cust_name`
- `item_screen_array.item_label`
- `formonly.total`

When no field name prefix is used, the first form field matching that simple field name is
used.

When using a prefix in the field name specification, it must match the field prefix
assigned by the dialog field binding method used at the beginning of the interactive statement: When
no screen-record has been explicitly specified in the field binding clause (for example, when using
`INPUT BY NAME variable-list`), the field prefix must be
the database table name or [`FORMONLY`](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns."), as defined in the form file, or any valid screen-record using
that field. When the `FROM` clause of the dialog specifies an explicit screen-record
(for example, in `INPUT variable-list FROM screen-record.* /
field-list-with-screen-record-prefix` or `INPUT ARRAY
array-name FROM screen-array.*`), the field
prefix must be the screen-record name used in the `FROM` clause.

Abstract field identification is supported with the `CURRENT`,
`NEXT` and `PREVIOUS` keywords. `NEXT FIELD CURRENT`
forces to stay in the current field; `NEXT FIELD NEXT` selects the field right after
the current field; `NEXT FIELD PREVIOUS` selects the field right before the current
field. When using `FIELD ORDER FORM`, the `NEXT` and
`PREVIOUS` options follow the tabbing order defined by the form. Otherwise, they
follow the order defined by the input binding list (with the `FROM` or `BY
NAME` clause). The field selected with `NEXT` or `PREVIOUS`
keywords may be non-focusable.

Usage of `NEXT FIELD PREVIOUS` and `NEXT FIELD NEXT` is
discouraged: The target field may be non-focusable (it can be enabled/disabled during dialog
execution). Dependending on the context, the field selection can be based on the last field
navigation direction, and can lead to infinite loops. Always use `NEXT FIELD CURRENT`
or `NEXT FIELD field-name` of a field that is focusable.

## NEXT FIELD to a non-focusable field

Non-focusable fields are:

- Fields defined with the `NOENTRY` attribute in the form definition file.
- Fields defined with a widget that does not allow input, such as a [`LABEL`](1696-label-item-type.md "Defines a simple text area to display a read-only value."), [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource.").
- Fields defined as `PHANTOM` fields in the form.
- Fields disabled at runtime with [`DIALOG.setFieldActive("field-spec",FALSE)`](../15_library-reference/3226-ui-dialog-setfieldactive.md "Enable and disable form fields.").
- Fields that are hidden or when a parent container (`TABLE`, `GRID`)
  is hidden, in the form definition file with a [`HIDDEN`](1783-hidden-attribute.md "The HIDDEN attribute indicates if an element must be visible to the user.") attribute, or by program with [`ui.Form.setFieldHidden()`](../15_library-reference/3161-ui-form-setfieldhidden.md "Show or hide a form field.") or [`ui.Form.setElementHidden()`](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements.").

`NEXT FIELD field-name` may explicitly specify a non-focusable
field. A non-focusable field can also be the target, when `NEXT FIELD NEXT` or
`NEXT FIELD PREVIOUS` selects a non-focusable field, based on the current [tabbing order](2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.").

> **Important:**
>
> Doing a `NEXT FIELD` to a non-focusable field can lead to
> infinite loops in the dialog; This practice is discouraged. To stay in the current field / row, use
> `NEXT FIELD CURRENT` instead.

## NEXT FIELD in procedural DIALOG blocks

In a [procedural dialog block](2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling."), the
`NEXT FIELD field-name` instruction gives the focus to the specified field
controlled by `INPUT`, `INPUT ARRAY` or `CONSTRUCT`, or
to a read-only list when using `DISPLAY ARRAY`.

When using a `DISPLAY ARRAY` sub-dialog, it is possible to give the focus to the
list, by specifying the name of the first column as argument for `NEXT FIELD`.

If the target field specified in the `NEXT FIELD` instruction is inside the
current sub-dialog, neither `AFTER FIELD` nor [`AFTER ROW`](1976-after-row-block.md) will be invoked for the field
or list you are leaving. However, the `BEFORE FIELD` control blocks of the
destination field (or the [`BEFORE
ROW`](1975-before-row-block.md) in case of read-only list) will be executed.

If the target field specified in the `NEXT FIELD` instruction is outside
the current sub-dialog, the `AFTER FIELD`, [`AFTER INSERT`](2021-after-insert-block.md), `AFTER
ROW`, and `AFTER INPUT`, `AFTER DISPLAY`, `AFTER
/CONSTRUCT` control blocks will be invoked for the field or list you are leaving. [Form-level validation rules](2240-form-level-validation-rules.md "Form-level validation rules can be defined for each field controlled by a dialog.") will also be
checked, as if the user had selected the new sub-dialog himself. This guarantees the current
sub-dialog is left in a consistent state. The `BEFORE INPUT`, `BEFORE
DISPLAY`, `BEFORE CONSTRUCT`, `BEFORE ROW` and the
`BEFORE FIELD` control blocks of the destination field / list will then be
executed.

## NEXT FIELD in record list control blocks

When using `NEXT FIELD` in [`AFTER
ROW`](1976-after-row-block.md) or in [`ON ROW
CHANGE`](2018-on-row-change-block.md) of a `DISPLAY ARRAY` or `INPUT ARRAY`, the
dialog will stay in the current row and give control back to the user.

This behavior allows you to implement data input rules:

```
  AFTER ROW
    IF NOT int_flag AND arr_count()<=arr_curr() THEN
       IF arr[arr_curr()].it_count * arr[arr_curr()].it_value > maxval THEN
          ERROR "Amount of line exceeds max value."
          NEXT FIELD item_count
       END IF
    END IF
```

If a condition is not met for a given field, do a `NEXT FIELD` to that field to
point the user to the problem. If a global condition for the current row is not met, use
`NEXT FIELD CURRENT` to stay in the current row and current field.

## NEXT FIELD in DISPLAY ARRAY with FOCUSONFIELD

`DISPLAY ARRAY` dialogs support cell-oriented navigation with the [`FOCUSONFIELD`](2091-display-array-attributes-clause.md) dialog attribute.

The `NEXT FIELD` instruction can be used in `DISPLAY ARRAY` with
`FOCUSONFIELD`, to give the focus to a specific field in the current row.
`NEXT FIELD` can be combined with [`DIALOG.setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list."), to give the focus to a specific cell in the
table.

## Related links

**Related concepts**  

[Giving the focus to a form element](2245-giving-the-focus-to-a-form-element.md "How to force the focus to move or stay in a specific form element using program code.")

[Understanding multiple dialogs](2077-understanding-multiple-dialogs.md "Multiple dialogs are defined with DIALOG blocks inside a FUNCTION.")
