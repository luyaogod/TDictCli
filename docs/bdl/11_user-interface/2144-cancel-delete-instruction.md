---
title: "CANCEL DELETE instruction"
source: "fgl-topics/c_fgl_dialog_CANCEL_DELETE_2.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control instructions > CANCEL DELETE instruction"
type: "concept"
description: "Syntax CANCEL DELETE Usage In a list controlled by an INPUT ARRAY , row deletion can be canceled by using the CANCEL DELETE instruction in the BEFORE DELETE block. Using this instruction in a ..."
---

# CANCEL DELETE instruction

## Syntax

```
CANCEL DELETE
```

## Usage

In a list controlled by an `INPUT ARRAY`, row deletion can be canceled by using
the `CANCEL DELETE` instruction in the `BEFORE DELETE` block. Using
this instruction in a different place will generate a compilation error.

When the `CANCEL DELETE` instruction is executed, the current `BEFORE
DELETE` block is terminated without any other trigger execution (no `BEFORE
ROW` or `BEFORE FIELD` is executed), and the program execution continues in
the user event loop.

You can, for example, prevent row deletion based on some condition:

```
   BEFORE DELETE
      IF user_can_delete() == FALSE THEN
         ERROR "You are not allowed to delete rows"
         CANCEL DELETE
      END IF
```

The instructions that appear after `CANCEL DELETE` will be skipped.

If the row deletion condition is known before the delete action occurs, disable the delete action
to prevent the user from performing a delete row action with the
`DIALOG.setActionActive()` method:

```
   CALL DIALOG.setActionActive("delete", FALSE)
```

It is also possible to prevent the user from deleting rows with the `DELETE ROW =
FALSE` option in the `ATTRIBUTE` clause.

## Related links

**Related concepts**  

[BEFORE DELETE block](2022-before-delete-block.md "BEFORE DELETE block")
