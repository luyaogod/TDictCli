---
title: "CONTINUE CONSTRUCT instruction"
source: "fgl-topics/c_fgl_Construct_CONTINUE_CONSTRUCT.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > CONSTRUCT control instructions > CONTINUE CONSTRUCT instruction"
type: "concept"
description: "CONTINUE CONSTRUCT skips all subsequent statements in the current control block and gives the control back to the dialog. CONSTRUCT BY NAME where_part ON ... ... ON ACTION check_condition IF ..."
---

# CONTINUE CONSTRUCT instruction

`CONTINUE CONSTRUCT` skips all subsequent statements in the current control block
and gives the control back to the dialog.

```
CONSTRUCT BY NAME where_part ON ...
  ...
  ON ACTION check_condition
    IF getCustomerStatus() != "A" THEN
      ERROR "Customer status does not allow query"
      CONTINUE CONSTRUCT
    END IF
    ...
END CONSTRUCT
```

This instruction is useful when program control is nested within multiple conditional statements,
and you want to return the control to the dialog. If this instruction is called in a control block
that is not [`AFTER CONSTRUCT`](2058-after-construct-block.md),
further control blocks might be executed depending on the context.

`CONTINUE CONSTRUCT` instructs the dialog to continue as if the code in the
control block was terminated (therefore it acts as a kind of `GOTO
end_of_control_block`). However, when executed in `AFTER CONSTRUCT`, the
focus returns to the most recently occupied field in the current form, giving the user another
chance to enter data in that field. In this case the [`BEFORE FIELD`](1942-before-field-block.md) of the current field will be invoked.

As alternative, use the [`NEXT
FIELD`](1955-next-field-instruction.md) control instruction to give the focus to a specific field and force the dialog
to continue. However, unlike `CONTINUE CONSTRUCT`, the `NEXT FIELD`
instruction will skip the next control blocks that are normally executed.

The `CONTINUE CONSTRUCT` instruction can only be used in a singular
`CONSTRUCT` dialog, it cannot be used in a `DIALOG / END DIALOG`
multiple dialog block.

## Related links

**Related concepts**  

[ACCEPT CONSTRUCT instruction](2068-accept-construct-instruction.md "ACCEPT CONSTRUCT instruction")

[EXIT CONSTRUCT instruction](2070-exit-construct-instruction.md "EXIT CONSTRUCT instruction")
