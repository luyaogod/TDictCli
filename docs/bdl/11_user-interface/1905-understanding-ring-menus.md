---
title: "Understanding ring menus"
source: "fgl-topics/c_fgl_menus_002.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Understanding ring menus"
type: "concept"
---

# Understanding ring menus

> The MENU instruction implements a set of choices, also known as action handlers.

A `MENU` dialog defines a list of options that can trigger actions to execute
associated program code. Ring menus are implemented with the `MENU` interactive
instruction.

A `MENU` block lists the possible actions that can be triggered in a given place
in the program, with the associated program code to be
executed.

```
MENU "Sample"
    COMMAND "Say hello"
       DISPLAY "Hello, world!"
    COMMAND "Exit"
       EXIT MENU 
END MENU
```

A ring menu can only define a set of options for a given level
of the program. You cannot define all menu options of your program in
a single `MENU` instruction; you must implement nested menus.

The `MENU` instruction is mainly designed for text mode applications, displaying
ring menus at the top of the screen. A typical TUI mode application starts with a global menu,
defining general options to access subroutines, which in turn implement specific menus with database
record handling options such as 'Append', 'Delete', 'Modify', and 'Search'.

Ring menus can also be used in a GUI application. However, as this instruction does not
handle form fields, other parts of the form are disabled during the menu dialog execution. In GUI
applications, ring menus are typically used to open a modal window with Yes / No / Cancel options,
where options are typically defined with the `ON ACTION` clauses, to keep the code
abstract and define decoration in form or with
attributes:

```
MENU ""
    ON ACTION print ATTRIBUTES(TEXT=%"print",IMAGE="printer")
       CALL print_document()
    ON ACTION save -- Get decoration from form or action defaults 
       CALL save_document()
    ON ACTION close -- Binds to GUI window close button
       EXIT PROGRAM
END MENU
```

## Related links

**Related concepts**  

[Action handling basics](2254-action-handling-basics.md "This topic describes the basic concepts of dialog actions.")

[The model-view-controller paradigm](2219-the-model-view-controller-paradigm.md "The dynamic user interface architecture is based on the Model-View-Controller (MVC) paradigm.")
