---
title: "MENU programming steps"
source: "fgl-topics/t_fgl_menus_007.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > MENU programming steps"
type: "task"
---

# MENU programming steps

> Follow this procedure to use the MENU dialog instruction.

To implement a `MENU` statement:

1. Create a `MENU` block with a title and write
   the end of the menu block with the `END
   MENU` keywords.
2. Depending on the type of menu rendering you need, add an `ATTRIBUTES` clause
   with the required [`STYLE`](1909-rendering-modes-of-a-menu.md) attribute.
3. List all the options that you want to offer to the end
   user when the menu executes. Typical CRUD programs
   will implement "Append", "Modify", "Delete" operations
   for a given database application entity (customers, orders, items
   tables). Typical dialog box menus have "Yes" /
   "No" / "Cancel" options.
4. Whether TUI or GUI mode, define action views (`TOPMENU`, `TOOLBAR` or
   form `BUTTON`) for each menu
   action, and use either `COMMAND [KEY]` or `ON ACTION` clauses to
   define the menu options.
5. When the menu is not a pop-up or dialog menu, do not forget to implement an option to leave
   the menu with the `EXIT MENU` control instruction.
6. Implement the code to be executed in every option.
