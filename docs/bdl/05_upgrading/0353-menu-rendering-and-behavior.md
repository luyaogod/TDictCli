---
title: "MENU rendering and behavior"
source: "fgl-topics/c_fgl_MigI4GL_037.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > MENU rendering and behavior"
type: "concept"
---

# MENU rendering and behavior

> This topic describes differences between I4GL and FGL MENU instruction in TUI mode.

## Menu disappears when displaying to its lines

With IBM®
Informix® 4GL, it is possible to overwrite the menu
options and comment line with `DISPLAY text AT
line,column`: The menu elements remain visible and
will be refreshed when the control goes back to the `MENU`
dialog:

```
MAIN
   MENU "Menu1"
     COMMAND "Option1" "This is option #1"
     COMMAND "Option2" "This is option #2"
     COMMAND "Menu2"   "Executes second MENU" CALL menu2()
     COMMAND "Quit" EXIT MENU
   END MENU
END MAIN

FUNCTION menu2()
   OPTIONS MENU LINE 10
   MENU "Menu2"
     COMMAND "DisplayAt1" DISPLAY "xxxxxxxxxxxxxx" AT 1,6
     COMMAND "DisplayAt2" DISPLAY "xxxxxxxxxxxxxx" AT 2,10
     COMMAND "Quit" EXIT MENU
   END MENU
   OPTIONS MENU LINE 1
END FUNCTION
```

When
in TUI mode with Genero BDL, the first `MENU` will disappear as soon as some text is
displayed in one of the menu lines. Genero understands that you want to reuse space used by the
first menu, and clears the menu lines to avoid corrupted screens.

## Menu disappears when leaving the `MENU` instruction

With IBM Informix 4GL, the menu options and comment text remain visible even when the
`MENU` instruction ends. In the next code example, after the first
`MENU` instruction terminates, the menu display line is modified (`OPTIONS
MENU LINE`) and a second `MENU` instruction is executed. The first
`MENU` is still visible while executing the second `MENU`:

```
MAIN
   MENU "Menu1"
     COMMAND "Option1"
     COMMAND "Quit" EXIT MENU
   END MENU
   OPTIONS MENU LINE 10
   MENU "Menu2"
     COMMAND "Option1"
     COMMAND "Quit" EXIT MENU
   END MENU
END MENU
```

With Genero BDL, the first `MENU` disappears when the second `MENU`
starts.

## Using Up/Down arrow keys

With IBM Informix 4GL, when a `MENU` defines more options than can fit into a 80x24
screen, the user can press Up/Down keys to jump directly to the previous or next page of
options.

With Genero BDL, Up/Down keys move to the previous or next option, just like Left/Right keys.

Consider using short `MENU` options that can fit into the screen.

This is no longer an issue when using the GUI mode.

## Menu command beginning with space

With IBM Informix 4GL, when a `MENU` option starts with a space, the end user can select
this option by pressing the spacebar on the keyboard.

With Genero BDL, spacebar moves to the next menu option.

```
MAIN
    MENU "test"
       COMMAND "aaa"
       COMMAND "bbb"
       COMMAND " ccc"     -- starts with a space!
       COMMAND "exit"
            EXIT MENU
    END MENU
END MAIN
```

This looks more like a programming mistake and should be fixed in the user code.

## Related links

**Related concepts**  

[Ring menus (MENU)](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.")
