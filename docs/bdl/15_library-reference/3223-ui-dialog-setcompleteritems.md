---
title: "ui.Dialog.setCompleterItems"
source: "fgl-topics/c_fgl_ClassDialog_setCompleterItems.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setCompleterItems"
type: "concept"
---

# ui.Dialog.setCompleterItems

> Define autocompletion items for a field defined with COMPLETER attribute.

## Syntax

```
setCompleterItems(
   items DYNAMIC ARRAY OF STRING )
```

1. items defines the list of completion proposals to be passed to the front-end.

## Usage

The `setCompleterItems()` dialog method defines the list of completion proposals
for the current field, to implement autocompletion.

The field must be defined in the form with the `COMPLETER` attribute.

The list of completion proposal items is passed as a dynamic array of
strings:

```
DEFINE items DYNAMIC ARRAY OF STRING
```

To clean up the completion proposal list for a given field, pass `NULL` as second
parameter to the function.

The method will raise error [-8114](4483-genero-bdl-errors.md), if the list of items contains more than 50 elements. Note that this error is not
trappable with exception handlers like `TRY/CATCH`, the code must avoid reaching the
limit.

See [Enabling autocompletion](../11_user-interface/2247-enabling-autocompletion.md "Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field.") for more details.

## Example

```
DEFINE items DYNAMIC ARRAY OF STRING
...
   ON CHANGE firstname
      -- fill the array with items
      LET items[1] = "Ann"
      LET items[2] = "Anna"
      LET items[3] = "Annabel"
      CALL DIALOG.setCompleterItems(items)
```

## Related links

**Related concepts**  

[Enabling autocompletion](../11_user-interface/2247-enabling-autocompletion.md "Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field.")

[COMPLETER attribute](../11_user-interface/1770-completer-attribute.md "The COMPLETER attribute enables autocompletion for the edit field.")
