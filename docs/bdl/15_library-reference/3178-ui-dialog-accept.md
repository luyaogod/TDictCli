---
title: "ui.Dialog.accept"
source: "fgl-topics/c_fgl_ClassDialog_accept.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.accept"
type: "concept"
---

# ui.Dialog.accept

> Validates and terminates the dialog.

## Syntax

```
accept()
```

## Usage

Use the `accept()` method to validate field input and terminate the dialog.

The `accept()` method is equivalent to the [`ACCEPT INPUT`](../11_user-interface/1951-accept-input-instruction.md) / [`ACCEPT DISPLAY`](../11_user-interface/1998-accept-display-instruction.md) / [`ACCEPT DIALOG`](../11_user-interface/2142-accept-dialog-instruction.md) instructions.

It can be used for example to terminate the dialog in a function, outside the context of a dialog block, where control instructions cannot be used.

> **Tip:**
>
> To skip code following the call to the `accept()` method, use the
> `CONTINUE DIALOG` instruction: Combining `accept()` + `CONTINUE
> DIALOG` is equivalent to the [`ACCEPT DIALOG`](../11_user-interface/2142-accept-dialog-instruction.md) instruction.

When calling the `accept()` method, the `AFTER INPUT`,
`AFTER DISPLAY` or `AFTER CONSTRUCT` block of the current subdialog is
executed, then the `AFTER DIALOG` block is executed.

Typical dialog validation rules are performed
when calling this method. See [`ACCEPT DIALOG`](../11_user-interface/2142-accept-dialog-instruction.md)
for more details.

## Example

```
SCHEMA stores
DEFINE rec RECORD LIKE customer.*
DIALOG mysubinput()
    INPUT BY NAME rec.*
       ON ACTION accept
          CALL DIALOG.accept()
    END INPUT
END DIALOG
```

## Related links

**Related concepts**  

[ui.Dialog.cancel](3188-ui-dialog-cancel.md "Cancels a parent dialog from a sub-dialog.")

[ui.Dialog.validate](3233-ui-dialog-validate.md "Checks form level validation rules.")
