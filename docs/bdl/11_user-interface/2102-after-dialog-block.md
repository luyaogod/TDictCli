---
title: "AFTER DIALOG block"
source: "fgl-topics/c_fgl_DIALOG_block_AFTER_DIALOG.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control blocks > AFTER DIALOG block"
type: "concept"
description: "Syntax AFTER DIALOG instruction [...] Usage The AFTER DIALOG block is executed one time as the last trigger when the DIALOG instruction terminates, when performing an ACCEPT DIALOG instruction. Dialog ..."
---

# AFTER DIALOG block

## Syntax

```
AFTER DIALOG
   instruction [...]
```

## Usage

The `AFTER DIALOG` block is executed one time as the last trigger when the
`DIALOG` instruction terminates, when performing an `ACCEPT DIALOG`
instruction. Dialog finalization code can be implemented in this block.

The dialog terminates when an `ACCEPT DIALOG` or `EXIT DIALOG`
control instruction is executed. However, the `AFTER DIALOG` block is  not
executed if an `EXIT DIALOG` is performed.

If you execute one of the following control instructions in an `AFTER DIALOG`
block, the dialog will not terminate and it will give control back to the user:

1. `NEXT FIELD`
2. `NEXT OPTION`
3. `CONTINUE DIALOG`

In the next example, the `AFTER DIALOG` block checks whether a field value is
correct and gives control back to the dialog if the value is wrong:

```
   ON ACTION accept
      ACCEPT DIALOG
   ...
   AFTER DIALOG
     IF NOT cust_is_valid_status(p_cust.cust_status) THEN
        ERROR "Customer state is not valid"
        NEXT FIELD cust_status
     END IF
```

## Related links

**Related concepts**  

[ACCEPT DIALOG instruction](2142-accept-dialog-instruction.md "ACCEPT DIALOG instruction")

[EXIT DIALOG instruction](2141-exit-dialog-instruction.md "EXIT DIALOG instruction")

[BEFORE DIALOG block](2101-before-dialog-block.md "BEFORE DIALOG block")
