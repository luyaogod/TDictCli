---
title: "CONTINUE DISPLAY instruction"
source: "fgl-topics/c_fgl_DisplayArray_CONTINUE_DISPLAY.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY control instructions > CONTINUE DISPLAY instruction"
type: "concept"
description: "Syntax CONTINUE DISPLAY Usage CONTINUE DISPLAY skips all subsequent statements in the current control block and gives the control back to the dialog. DISPLAY ARRAY custlist TO sr_cust.* ... ... ON ..."
---

# CONTINUE DISPLAY instruction

## Syntax

```
CONTINUE DISPLAY
```

## Usage

`CONTINUE DISPLAY` skips all subsequent statements in the current control block
and gives the control back to the dialog.

```
DISPLAY ARRAY custlist TO sr_cust.* ...
  ...
  ON ACTION check_row
    IF getCustomerStatus(arr_curr()) != "A" THEN
      ERROR "Customer status is not valid"
      CONTINUE DISPLAY
    END IF
    ...
END DISPLAY
```

The `CONTINUE DISPLAY` instruction is useful when program control is nested within
multiple conditional statements, and you want to return the control to the dialog. If this
instruction is called in a control block that is not [`AFTER DISPLAY`](1974-after-display-block.md), further control blocks might be executed depending on the
context.

Actually, `CONTINUE DISPLAY` just instructs the dialog to continue as if the code
in the control block was terminated (it's a kind of `GOTO end_of_control_block`).
However, when executed in `AFTER DISPLAY`, the focus returns to the current row in
the list, giving the user another chance to browse and select a row. In this case the `BEFORE
ROW` of the current row will be invoked.

The `CONTINUE DISPLAY` instruction can only be used in a singular `DISPLAY
ARRAY` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple dialog
block.

## Related links

**Related concepts**  

[ACCEPT DISPLAY instruction](1998-accept-display-instruction.md "ACCEPT DISPLAY instruction")

[EXIT DISPLAY instruction](1997-exit-display-instruction.md "EXIT DISPLAY instruction")
