---
title: "EXIT MENU instruction"
source: "fgl-topics/c_fgl_menus_EXIT_MENU.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > MENU control instructions > EXIT MENU instruction"
type: "concept"
description: "Syntax EXIT MENU Usage EXIT MENU statement terminates the MENU block and continues the program flow with the statement after the menu block. MENU \"Stock\" ... COMMAND \"Exit\" EXIT MENU END MENU The EXIT ..."
---

# EXIT MENU instruction

## Syntax

```
EXIT MENU
```

## Usage

`EXIT MENU` statement terminates the `MENU` block and continues the
program flow with the statement after the menu block.

```
MENU "Stock"
  ...
  COMMAND "Exit"
     EXIT MENU
END MENU
```

The `EXIT MENU` instruction can only be used in a singular `MENU`
dialog, it cannot be used in a `DIALOG / END DIALOG` multiple dialog block.

## Related links

**Related concepts**  

[EXIT block-name](../08_language-basics/0679-exit-block-name.md "The EXIT block instruction transfers control out of the current program block.")
