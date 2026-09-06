---
title: "ui.Dialog.nextField"
source: "fgl-topics/c_fgl_ClassDialog_nextField.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.nextField"
type: "concept"
---

# ui.Dialog.nextField

> Registers the next field to go to.

## Syntax

```
nextField(
   name STRING )
```

1. name is the form field name, see [Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md), or
   `"+NEXT"`,`"+PREV"`, `"+CURR"`, to specify respectively
   the next, previous and current field.

## Usage

The `nextField()` method registers the name of the next field that must get the
focus when control goes back to the dialog.

Like `NEXT FIELD`, the `nextField()` method is typically used in an
editable dialog (`INPUT`, `CONSTRUCT`, `INPUT
ARRAY`).

In a `DISPLAY ARRAY` using the [`FOCUSONFIELD` attribute](../11_user-interface/2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD."), `nextField()` can be used in
conjunction with `DIALOG.setCurrentRow()`, to set the focus to a specific cell in the
list.

The `nextField()` method is similar to the [`NEXT FIELD` instruction](../11_user-interface/1955-next-field-instruction.md), except that it
does not implicitly break the program flow. To get the same behavior as `NEXT FIELD`,
the method call must be followed by a [`CONTINUE DIALOG`](../11_user-interface/2140-continue-dialog-instruction.md) instruction, or an equivalent instruction such as
`CONTINUE INPUT`, in case of singular dialog.

Since this method takes an expression as parameter, you can write generic code, when the name of
the target field is not known at compile time. In the next example, the
`check_value()` function returns a field name where the value does not satisfy the
validation rules.

```
DEFINE fn STRING
...
ON ACTION save 
  IF ( fn:= check_values() ) IS NOT NULL THEN
    CALL DIALOG.nextField(fn)
    CONTINUE DIALOG
  END IF
  CALL save_data()
  ...
```

When specifying `"+NEXT"` as parameter, the dialog will register the field next to
the current field (like `NEXT FIELD NEXT`). When passing `"+PREV"`,
the dialog will register the previous field (like `NEXT FIELD PREVIOUS`). When
passing `"+CURR"`, it is equivalent to a `NEXT FIELD CURRENT`
instruction:

```
...
ON ACTION save 
  IF NOT valid_input() THEN
    CALL DIALOG.nextField("+CURR")
    CONTINUE DIALOG
  END IF
  CALL save_data()
  ...
```

## Related links

**Related concepts**  

[Giving the focus to a form element](../11_user-interface/2245-giving-the-focus-to-a-form-element.md "How to force the focus to move or stay in a specific form element using program code.")
