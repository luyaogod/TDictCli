---
title: "CANCEL INSERT instruction"
source: "fgl-topics/c_fgl_dialog_CANCEL_INSERT.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Using editable record lists > INPUT ARRAY control instructions > CANCEL INSERT instruction"
type: "concept"
description: "Syntax CANCEL INSERT Usage In a list controlled by an INPUT ARRAY , row creation can be canceled by the program with the CANCEL INSERT instruction. This instruction can only be used in the BEFORE ..."
---

# CANCEL INSERT instruction

## Syntax

```
CANCEL INSERT
```

## Usage

In a list controlled by an `INPUT ARRAY`, row creation can be canceled
by the program with the `CANCEL INSERT` instruction. This instruction can
only be used in the `BEFORE INSERT` and `AFTER INSERT`
control blocks. If it appears at a different place, the compiler will generate an error.

The instructions that appear after `CANCEL INSERT` will be skipped.

If the row creation condition is known before the insert/append action occurs, disable
the insert and/or append actions to prevent the user from creating new rows, with
`DIALOG.setActionActive()`:

```
   CALL DIALOG.setActionActive("insert", FALSE)
   CALL DIALOG.setActionActive("append", FALSE)
```

However, this will not prevent the user from appending a new temporary
row at the end of the list, when moving down after the last
row. To prevent row creation completely, use the `INSERT
ROW = FALSE` and `APPEND ROW =FALSE` options
in the `ATTRIBUTE` clause of `INPUT
ARRAY`, or combine with the `AUTO APPEND = FALSE`
attribute.

## CANCEL INSERT in BEFORE INSERT

A `CANCEL INSERT` executed inside a `BEFORE INSERT`
block prevents the new row creation. The following tasks are performed:

1. No new row will be created (the new row is not yet shown to the user).
2. The `BEFORE INSERT` block is terminated (further instructions are
   skipped).
3. The `BEFORE ROW` and `BEFORE FIELD` triggers
   are executed.
4. Control goes back to the user.

You can, for example, cancel a row creation if the user is not allowed to create
rows:

```
   BEFORE INSERT
      IF NOT user_can_insert THEN
         ERROR "You are not allowed to insert rows"
         CANCEL INSERT
      END IF
```

Executing `CANCEL INSERT` in `BEFORE INSERT` will
also cancel a temporary row creation, except when there are no more rows in the list.
In this case, `CANCEL INSERT` will just be ignored and leave the
new row as is (otherwise, the instruction would loop without end). You can prevent
automatic temporary row creation with the `AUTO APPEND=FALSE` attribute.
If `AUTO APPEND=FALSE` and a `CANCEL INSERT` is executed
in `BEFORE INSERT` (user has invoked an append action), the temporary
row will be deleted and list will remain empty if it was the last row.

## CANCEL INSERT in AFTER INSERT

A `CANCEL INSERT` executed inside an `AFTER INSERT`
block removes the newly created row. The following tasks are performed:

1. The newly created row is removed from the list (the row exists now and user has
   entered data).
2. The `AFTER INSERT` block is terminated (further instructions are
   skipped).
3. The `BEFORE ROW` and `BEFORE FIELD` triggers are
   executed.
4. The control goes back to the user.

You can, for example, cancel a row insertion if a database error occurs when you
try to insert the row into a database table:

```
   AFTER INSERT
      WHENEVER ERROR CONTINUE
      LET r = DIALOG.getCurrentRow("s_items")
      INSERT INTO items VALUES ( p_items[r].* )
      WHENEVER ERROR STOP
      IF sqlca.sqlcode<>0 THEN
         ERROR SQLERRMESSAGE
         CANCEL INSERT
      END IF
```

## Related links

**Related concepts**  

[BEFORE DELETE block](2022-before-delete-block.md "BEFORE DELETE block")

[Appending rows in INPUT ARRAY](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.")
