---
title: "AFTER DISPLAY block"
source: "fgl-topics/c_fgl_dialog_AFTER_DISPLAY_2.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG control blocks > AFTER DISPLAY block"
type: "concept"
description: "Syntax AFTER DISPLAY instruction [...] AFTER DISPLAY block in singular DISPLAY ARRAY dialogs In a singular DISPLAY ARRAY instruction, the AFTER DISPLAY is only executed once when the dialog is ended. ..."
---

# AFTER DISPLAY block

## Syntax

```
AFTER DISPLAY
   instruction [...]
```

## AFTER DISPLAY block in singular DISPLAY ARRAY dialogs

In a singular `DISPLAY ARRAY` instruction, the `AFTER DISPLAY` is
only executed once when the dialog is ended.

You typically implement dialog finalization in this block.

```
DISPLAY ARRAY p_items TO s_items.*
   AFTER DISPLAY
       DISPLAY "Current row is: ", arr_curr()
```

## AFTER DISPLAY block in DISPLAY ARRAY of procedural DIALOG

In a `DISPLAY ARRAY` sub-dialog of a procedural `DIALOG`
instruction, the `AFTER DISPLAY` block is executed when a `DISPLAY
ARRAY` list loses the focus and the focus goes to another sub-dialog.

If the focus leaves the current group and goes to an action view, this trigger is not
executed, because the focus has not gone to another sub-dialog yet.

`AFTER DISPLAY` is executed after the `AFTER
ROW` block.

In this example, the `AFTER DISPLAY` block disables
an action that is specific to the current list:

```
DISPLAY ARRAY p_items TO s_items.*
   AFTER DISPLAY
       CALL DIALOG.setActionActive("clear_item_list", FALSE)
```

## Related links

**Related concepts**  

[AFTER INPUT block](1941-after-input-block.md "AFTER INPUT block")

[AFTER CONSTRUCT block](2058-after-construct-block.md "AFTER CONSTRUCT block")

[BEFORE DISPLAY block](1973-before-display-block.md "BEFORE DISPLAY block")
