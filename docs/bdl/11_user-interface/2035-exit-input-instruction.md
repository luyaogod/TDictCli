---
title: "EXIT INPUT instruction"
source: "fgl-topics/c_fgl_InputArray_EXIT_INPUT.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Using editable record lists > INPUT ARRAY control instructions > EXIT INPUT instruction"
type: "concept"
description: "Syntax EXIT INPUT Usage Use the EXIT INPUT to terminate the INPUT ARRAY instruction and resume the program execution at the instruction following the INPUT ARRAY block. ON ACTION leave_now EXIT INPUT ..."
---

# EXIT INPUT instruction

## Syntax

```
EXIT INPUT
```

## Usage

Use the `EXIT INPUT` to terminate the `INPUT ARRAY` instruction and
resume the program execution at the instruction following the `INPUT ARRAY`
block.

```
ON ACTION leave_now EXIT INPUT
```

When leaving the `INPUT ARRAY` instruction, all form items used by the dialog will
be disabled until another interactive statement takes control.

`EXIT INPUT` does not set `int_flag` to `TRUE` as when the cancel action is fired.

The `EXIT INPUT` instruction can only be used in a singular `INPUT
ARRAY` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple dialog
block.

## Related links

**Related concepts**  

[ACCEPT INPUT instruction](2034-accept-input-instruction.md "ACCEPT INPUT instruction")
