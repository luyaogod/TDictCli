---
title: "EXIT DISPLAY instruction"
source: "fgl-topics/c_fgl_DisplayArray_EXIT_DISPLAY.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY control instructions > EXIT DISPLAY instruction"
type: "concept"
description: "Syntax EXIT DISPLAY Usage Use the EXIT DISPLAY to terminate the DISPLAY ARRAY instruction and resume the program execution at the instruction immediately following the DISPLAY ARRAY block. ON ACTION ..."
---

# EXIT DISPLAY instruction

## Syntax

```
EXIT DISPLAY
```

## Usage

Use the `EXIT DISPLAY` to terminate the `DISPLAY ARRAY` instruction
and resume the program execution at the instruction immediately following the `DISPLAY
ARRAY` block.

```
ON ACTION leave_now EXIT DISPLAY
```

`EXIT DISPLAY` does not set `int_flag` to `TRUE` as when the cancel action is fired.

The `EXIT DISPLAY` instruction can only be used in a singular `DISPLAY
ARRAY` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple dialog
block.

## Related links

**Related concepts**  

[ACCEPT DISPLAY instruction](1998-accept-display-instruction.md "ACCEPT DISPLAY instruction")

[CONTINUE DISPLAY instruction](1996-continue-display-instruction.md "CONTINUE DISPLAY instruction")
