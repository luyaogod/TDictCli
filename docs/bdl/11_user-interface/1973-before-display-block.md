---
title: "BEFORE DISPLAY block"
source: "fgl-topics/c_fgl_dialog_BEFORE_DISPLAY.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY control blocks > BEFORE DISPLAY block"
type: "concept"
description: "Syntax BEFORE DISPLAY instruction [...] BEFORE DISPLAY block in singular DISPLAY ARRAY dialogs In a singular DISPLAY ARRAY instruction, the BEFORE DISPLAY is only executed once when the dialog is ..."
---

# BEFORE DISPLAY block

## Syntax

```
BEFORE DISPLAY
   instruction [...]
```

## BEFORE DISPLAY block in singular DISPLAY ARRAY dialogs

In a singular `DISPLAY ARRAY` instruction, the `BEFORE DISPLAY` is
only executed once when the dialog is started.

The `BEFORE DISPLAY` block is executed once at dialog startup, before the
runtime system gives control to the user. This block can be used to display messages to the user,
initialize program variables, and set up the dialog instance by deactivating actions the user
is not allowed to execute.

```
DISPLAY ARRAY p_items TO s_items.*
   BEFORE DISPLAY
       CALL DIALOG.setActionActive("clear_item_list", is_super_user())
```

## BEFORE DISPLAY block DISPLAY ARRAY of procedural DIALOG

In a `DISPLAY ARRAY` sub-dialog of a procedural `DIALOG`
instruction, the `BEFORE DISPLAY` block is executed when a `DISPLAY
ARRAY` list gets the focus.

`BEFORE DISPLAY` is executed before the `BEFORE ROW`
block.

In this example the `BEFORE DISPLAY` block enables
an action and displays a message:

```
DISPLAY ARRAY p_items TO s_items.*
   BEFORE DISPLAY
       CALL DIALOG.setActionActive("print_list", TRUE)
       MESSAGE "You are now in the list of items"
```

## Related links

**Related concepts**  

[BEFORE INPUT block](1940-before-input-block.md "BEFORE INPUT block")

[BEFORE CONSTRUCT block](2057-before-construct-block.md "BEFORE CONSTRUCT block")

[AFTER DISPLAY block](1974-after-display-block.md "AFTER DISPLAY block")
