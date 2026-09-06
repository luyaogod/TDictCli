---
title: "AFTER FIELD block"
source: "fgl-topics/c_fgl_dialog_AFTER_FIELD_5.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control blocks > AFTER FIELD block"
type: "concept"
description: "Syntax AFTER FIELD field-spec [ , ...] instruction [...] Usage In dialog instructions INPUT , INPUT ARRAY , CONSTRUCT or in a DISPLAY ARRAY using the FOCUSONFIELD attribute, the AFTER FIELD block is ..."
---

# AFTER FIELD block

## Syntax

```
AFTER FIELD field-spec [,...]
   instruction [...]
```

## Usage

In dialog instructions `INPUT`, `INPUT ARRAY`,
`CONSTRUCT` or in a `DISPLAY ARRAY` using the
`FOCUSONFIELD` attribute, the `AFTER FIELD` block is executed every
time the focus leaves the specified field.

For single record inputs driven by `INPUT` or query by example (QBEs) driven by
`CONSTRUCT`, the `AFTER FIELD` block is executed when moving the focus
from field to field.

For editable lists driven by `INPUT ARRAY`, the `AFTER FIELD` block
is executed when moving the focus from field to field in the same row, or when moving to another row
in the same column.

For record lists driven by `DISPLAY ARRAY` using the `FOCUSONFIELD`
attribute, the `AFTER FIELD` block is executed when moving the focus from field to
field. However, the fields will not be editable as in an `INPUT ARRAY`.

The `AFTER FIELD` keywords must be followed by a list of form field
specifications. The screen-record name can be omitted.

`AFTER FIELD` is executed before `AFTER INSERT`, `ON
ROW CHANGE`, `AFTER ROW`, `AFTER INPUT` or `AFTER
CONSTRUCT`.

When a `NEXT FIELD` instruction
is executed in an `AFTER FIELD` block, the cursor moves to the specified field, which
can be the current field. This can be used to prevent the user from moving to another field / row
during data input. Note that the `BEFORE
FIELD` block is also executed when `NEXT FIELD` is invoked.

The `AFTER FIELD` block of the current field is not executed when performing a
`NEXT FIELD`; only `BEFORE INPUT`, `BEFORE CONSTRUCT`,
`BEFORE ROW`, and `BEFORE FIELD` of the target item might be executed,
depending on the sub-dialog type.

When `ACCEPT DIALOG`, `ACCEPT INPUT`, or `ACCEPT
CONSTRUCT` is performed, the `AFTER FIELD` trigger of the current field is
executed.

Use the `AFTER FIELD` block to implement field validation rules:

```
INPUT BY NAME p_item.* ...
   AFTER FIELD item_quantity
     IF p_item.item_quantity <= 0 THEN
        ERROR "Item quantity cannot be negative or zero"
        LET p_item.item_quantity = 0
        NEXT FIELD item_quantity
     END IF
```

## Related links

**Related concepts**  

[ACCEPT DIALOG instruction](2142-accept-dialog-instruction.md "ACCEPT DIALOG instruction")

[ON CHANGE block](1943-on-change-block.md "ON CHANGE block")
