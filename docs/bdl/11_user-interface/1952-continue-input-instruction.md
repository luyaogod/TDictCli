---
title: "CONTINUE INPUT instruction"
source: "fgl-topics/c_fgl_record_input_CONTINUE_INPUT.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Using simple record inputs > INPUT control instructions > CONTINUE INPUT instruction"
type: "concept"
description: "Syntax CONTINUE INPUT Usage CONTINUE INPUT skips all subsequent statements in the current control block and gives the control back to the dialog. This instruction is useful when program control is ..."
---

# CONTINUE INPUT instruction

## Syntax

```
CONTINUE INPUT
```

## Usage

`CONTINUE INPUT` skips all subsequent statements in the current control block and
gives the control back to the dialog. This instruction is useful when program control is nested
within multiple conditional statements, and you want to return the control to the dialog.

```
INPUT BY NAME cust_rec.*
  ...
  ON ACTION process_order
    IF cust_rec.cust_status != "A" THEN
      ERROR "Customer status does not allow order processing"
      CONTINUE INPUT
    END IF
    CALL process_order()
  ...
END INPUT
```

If this instruction is called in a control block that is not [`AFTER INPUT`](1941-after-input-block.md), further control blocks
might be executed depending on the context.

`CONTINUE INPUT` instructs the dialog to continue as if the code in the control
block was terminated (it's a kind of `GOTO end_of_control_block`). However, when
executed in `AFTER INPUT`, the focus returns to the most recently occupied field in
the current form, giving the user another chance to enter data in that field. In this case the
`BEFORE FIELD` of the current field will be invoked.

As alternative, use the [`NEXT
FIELD`](1955-next-field-instruction.md) control instruction to give the focus to a specific field and force the dialog
to continue. However, unlike `CONTINUE INPUT`, the `NEXT FIELD`
instruction will skip the further control blocks that are normally executed.

The `CONTINUE INPUT` instruction can only be used in a singular
`INPUT` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple
dialog block.

## Related links

**Related concepts**  

[EXIT INPUT instruction](1953-exit-input-instruction.md "EXIT INPUT instruction")

[ACCEPT INPUT instruction](1951-accept-input-instruction.md "ACCEPT INPUT instruction")

**Related reference**  

[INPUT control blocks execution order](1939-input-control-blocks-execution-order.md "INPUT control blocks execution order")
