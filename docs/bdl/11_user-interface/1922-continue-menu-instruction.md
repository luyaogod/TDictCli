---
title: "CONTINUE MENU instruction"
source: "fgl-topics/c_fgl_menus_CONTINUE_MENU.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > MENU control instructions > CONTINUE MENU instruction"
type: "concept"
description: "Syntax CONTINUE MENU Usage The CONTINUE MENU instruction resumes the execution of a classical MENU dialog. MENU \"Stock\" ... COMMAND \"Exit\" IF question(\"Exit the program?\")==FALSE THEN CONTINUE MENU ..."
---

# CONTINUE MENU instruction

## Syntax

```
CONTINUE MENU
```

## Usage

The `CONTINUE MENU` instruction resumes the execution of a classical
`MENU` dialog.

```
MENU "Stock"
  ...
  COMMAND "Exit"
    IF question("Exit the program?")==FALSE THEN
       CONTINUE MENU
    END IF
    CALL commit_changes()
    EXIT MENU
END MENU
```

`CONTINUE MENU` skips the instructions remaing in the current code block (that
belongs to the current dialog), redisplays the menu options, and gives the control back to the
user.

The `CONTINUE MENU` instruction can only be used in a singular
`MENU` dialog, it cannot be used in a `DIALOG / END DIALOG` multiple
dialog block.

## Related links

**Related concepts**  

[CONTINUE block-name](../08_language-basics/0678-continue-block-name.md "The CONTINUE block-name instruction resumes execution of a loop or dialog statement.")
