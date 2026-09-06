---
title: "ACCEPT DISPLAY instruction"
source: "fgl-topics/c_fgl_DisplayArray_ACCEPT_DISPLAY.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY control instructions > ACCEPT DISPLAY instruction"
type: "concept"
description: "Syntax ACCEPT DISPLAY Usage The ACCEPT DISPLAY instruction validates and exists the DISPLAY ARRAY instruction. DISPLAY ARRAY custlist TO sr_cust.* ... ... ON ACTION validate_row IF ..."
---

# ACCEPT DISPLAY instruction

## Syntax

```
ACCEPT DISPLAY
```

## Usage

The `ACCEPT DISPLAY` instruction validates and exists the `DISPLAY
ARRAY` instruction.

```
DISPLAY ARRAY custlist TO sr_cust.* ...
  ...
  ON ACTION validate_row
    IF validateRow(arr_curr()) THEN
      ACCEPT DISPLAY
    ELSE
      CONTINUE DISPLAY
    END IF
  ...
END DISPLAY
```

The [`AFTER ROW`](1976-after-row-block.md) and [`AFTER DISPLAY`](1974-after-display-block.md) control block will be
executed.

Statements after `ACCEPT DISPLAY` will be skipped.

The `ACCEPT DISPLAY` instruction can only be used in a singular `DISPLAY
ARRAY` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple dialog
block.

## Related links

**Related concepts**  

[EXIT DISPLAY instruction](1997-exit-display-instruction.md "EXIT DISPLAY instruction")

[CONTINUE DISPLAY instruction](1996-continue-display-instruction.md "CONTINUE DISPLAY instruction")
