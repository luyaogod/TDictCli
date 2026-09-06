---
title: "AUTONEXT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_AUTONEXT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > AUTONEXT attribute"
type: "concept"
---

# AUTONEXT attribute

> The AUTONEXT attribute forces the focus to automatically leave the current field when completed.

## Syntax

```
AUTONEXT
```

## Usage

When using the `AUTONEXT` attribute in a form item type that has a text
editor, the focus goes automatically to the next field, when the user enters a character
that completely fills the field. The maximum number of characters is defined by the [field input length](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

`AUTONEXT` is useful with character fields in which the input data is of a fixed
length, such as postal codes. It is also useful if a character field has a length of 1, as only one
keystroke is required to enter data and move to the next field.

When the form item type has a value picker such as calendars or drop down lists, the focus also
goes automatically to the next field, when a value is selected in the value picker.

When `AUTONEXT` is used for a [`COMBOBOX`](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.") field, the focus goes automatically to the next input field, when
selecting an item in the combobox drop-down list.

When `AUTONEXT` is used for an [`EDIT`](1690-edit-item-type.md "Defines a simple line-edit field.") or [`BUTTONEDIT`](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.") with a [`COMPLETER`](1770-completer-attribute.md "The COMPLETER attribute enables autocompletion for the edit field.") attribute, the focus goes automatically to the next input field,
when selecting an item from the completion list.

If the value entered in the field does not meet the requirements of other field attributes like
[`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field."), the focus stays in
the current field, and an error message displays.

The next field is defined by the field input order, see [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.").

## Related links

**Related concepts**  

[EDIT item type](1690-edit-item-type.md "Defines a simple line-edit field.")
