---
title: "BEFORE MENU block"
source: "fgl-topics/c_fgl_dialog_BEFORE_MENU.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > MENU control blocks > BEFORE MENU block"
type: "concept"
description: "Syntax BEFORE MENU instruction [...] Usage If the MENU block contains a BEFORE MENU clause, statements within this clause are executed before the menu dialog starts. This block is typically used to ..."
---

# BEFORE MENU block

## Syntax

```
BEFORE MENU
   instruction [...]
```

## Usage

If the `MENU` block contains a `BEFORE MENU` clause, statements
within this clause are executed before the menu dialog starts.

This block is typically used to hide or disable some menu options depending on the current
context of the program. For example, when the current user is not allowed to create new records, the
menu options can be disabled as follows:

```
MENU "Orders"
   BEFORE MENU
      CALL DIALOG.setActionActive("append", can_user_append() )
      ...
   COMMAND "Append" -- creates "append" action (lowercase)
      ...
   ...
END MENU
```

In TUI mode, the menu options can also be disabled, but they will still be displayed on the
screen. The end user will see the option, but cannot select it. In this case it's more convenient to
hide the option from the end user with the `DIALOG.setActionHidden()` method,
instead of disabling the action.

## Related links

**Related concepts**  

[The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.")
