---
title: "ON CHANGE block"
source: "fgl-topics/c_fgl_dialog_ON_CHANGE_3.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > CONSTRUCT control blocks > ON CHANGE block"
type: "concept"
description: "Syntax ON CHANGE field-spec [ , ...] instruction [...] Usage The ON CHANGE block can be used to detect when a field has been changed by user input. The ON CHANGE block is executed, if the value has ..."
---

# ON CHANGE block

## Syntax

```
ON CHANGE field-spec [,...]
   instruction [...]
```

## Usage

The `ON CHANGE` block can be used to detect when a field has been changed by user
input. The `ON CHANGE` block is executed, if the value has changed since the field
got the focus and the modification flag is set.

The `ON CHANGE` block can be used in `INPUT`, `INPUT
ARRAY` and `CONSTRUCT` dialogs.

For editable fields defined as `EDIT`, `TEXTEDIT` or
`BUTTONEDIT`, the `ON CHANGE` block is executed when leaving a field,
if the value of the specified field has changed since the field got the focus and the [modification flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.") is set for the field. The field
is left when user validates the dialog, when moving to another field, or when moving to another row
in an `INPUT ARRAY`. However, if the text edit field is defined with the
`COMPLETER` attribute to
enable [autocompletion](2247-enabling-autocompletion.md "Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field."), the `ON
CHANGE` trigger will be fired after a short period of time, when the user has typed
characters in.

For editable fields defined as `CHECKBOX`, `COMBOBOX`,
`DATEEDIT`, `DATETIMEEDIT`, `TIMEEDIT`,
`RADIOGROUP`, `SPINEDIT`, `SLIDER` or URL-based
`WEBCOMPONENT` (when the `COMPONENTTYPE` attribute is not used), the
`ON CHANGE` block is invoked immediately when the user changes the value with the
widget edition feature. For example, when toggling the state of a `CHECKBOX`, when
selecting an item in a `COMBOBOX` list, or when choosing a date in the calendar of a
`DATEEDIT`. Note that for such item types, when `ON CHANGE` is fired,
the [modification flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.") is always set.

```
   ON CHANGE order_checked -- Defined as CHECKBOX
      CALL setup_dialog(DIALOG)
```

If both an `ON CHANGE` block and `AFTER FIELD` block are defined
for a field, the `ON CHANGE` block is executed before the `AFTER
FIELD` block.

When changing by program the value of the current field in an `ON ACTION` block,
the `ON CHANGE` block will be executed when leaving the field, if the value is
different from the previous value and the modification flag is set (after previous user input
or when the touched flag has been changed by program).

In an `INPUT` or `INPUT ARRAY`, the field value change is related
to the value of the variable bound to the field. In a `CONSTRUCT` dialog, the field
value change is related to the input buffer / displayed value.

With a [`NEXT FIELD`](1955-next-field-instruction.md) instruction,
the state of the field value change is reset, as if the user had left and reentered the field. When
using `NEXT FIELD` in `ON CHANGE` block or in an `ON
ACTION` block, the `ON CHANGE` block will only be re-executed, if the value
is changed since `NEXT FIELD`, and the [modification flag](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.") is set. Therefore, `ON
CHANGE` should not be used for field validation with a `NEXT FIELD`, because
`ON CHANGE` will not get triggered again, if the (invalid) value has not changed;
Field validation rules must be implemented in `AFTER FIELD` blocks and/or
`AFTER INPUT` blocks.

## Related links

**Related concepts**  

[Form item types](1683-form-item-types.md "The form item types defines the purpose of form elements.")

[AFTER FIELD block](1944-after-field-block.md "AFTER FIELD block")
