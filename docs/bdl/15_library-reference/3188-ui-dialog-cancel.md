---
title: "ui.Dialog.cancel"
source: "fgl-topics/c_fgl_ClassDialog_cancel.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.cancel"
type: "concept"
---

# ui.Dialog.cancel

> Cancels a parent dialog from a sub-dialog.

## Syntax

```
cancel()
```

## Usage

The `cancel()` method can be used to terminate a dialog by ignoring the current
input.

This method can be used to terminate the dialog in a function, outside the context of a dialog
block, where control instructions cannot be used.

> **Tip:**
>
> To skip code following the call to the `cancel()` method, use the
> `CONTINUE DIALOG` instruction: Combining `cancel()` + `CONTINUE
> DIALOG` is equivalent to the [`CANCEL DIALOG`](../11_user-interface/2143-cancel-dialog-instruction.md) instruction.

When calling the `cancel()` method, the `int_flag` is set to
`TRUE`, the `AFTER INPUT`, `AFTER DISPLAY` or
`AFTER CONSTRUCT` block of the current subdialog is executed, then the `AFTER
DIALOG` block is executed.

The default settings regarding action attributes for the `cancel` action define
the `validate` attribute to `"no"`, in order to avoid current field
validation for this action. This is important when using the `UNBUFFERED` mode. For
more details, see [Actions configuration for field validation](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md).

## Example

```
SCHEMA stores
DEFINE rec RECORD LIKE customer.*
DIALOG mysubinput()
    INPUT BY NAME rec.*
       ON ACTION cancel
          CALL DIALOG.cancel()
    END INPUT
END DIALOG
```

## Related links

**Related concepts**  

[ui.Dialog.accept](3178-ui-dialog-accept.md "Validates and terminates the dialog.")
