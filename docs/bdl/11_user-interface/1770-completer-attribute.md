---
title: "COMPLETER attribute"
source: "fgl-topics/c_fgl_FSFAttributes_COMPLETER.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > COMPLETER attribute"
type: "concept"
---

# COMPLETER attribute

> The COMPLETER attribute enables autocompletion for the edit field.

## Syntax

```
COMPLETER
```

## Usage

Form fields with `COMPLETER` attribute provide suggestions while the end-user
types text into the field, it can be used in text edit fields such as `EDIT` and
`BUTTONEDIT` item types.

Normally, the `ON CHANGE` trigger
is fired for text edit fields when leaving the field and if the content was modified. Form fields
defined with the `COMPLETER` attribute will trigger the `ON CHANGE`
control block when the end user modifies the content of the field.

The `COMPLETER` attribute can be combined with the [`AUTONEXT`](1763-autonext-attribute.md "The AUTONEXT attribute forces the focus to automatically leave the current field when completed.") attribute, to
automatically give the focus to the next input fied, when an item is selected in the completion
list.

See [Enabling autocompletion](2247-enabling-autocompletion.md "Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field.") for more details.

## Example

```
EDIT f1 = FORMONLY.custname, COMPLETER;
```

## Related links

**Related concepts**  

[ui.Dialog.setCompleterItems](../15_library-reference/3223-ui-dialog-setcompleteritems.md "Define autocompletion items for a field defined with COMPLETER attribute.")

[EDIT item type](1690-edit-item-type.md "Defines a simple line-edit field.")

[BUTTONEDIT item type](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.")
