---
title: "The default SCREEN window"
source: "fgl-topics/c_fgl_MigI4GL_015.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > The default SCREEN window"
type: "concept"
---

# The default SCREEN window

> When the first interactive instruction is reached in a Genero BDL program, a default window named SCREEN is created.

The default `SCREEN` window can be used to open one or more successive forms; it
can also be closed, with the `CLOSE WINDOW SCREEN` instruction. If the default
`SCREEN` window is not closed, and a new window is created with the `OPEN
WINDOW` command, an empty default `SCREEN` window will be displayed.

When writing a GUI application, you typically open the main form
in the `SCREEN` window, and display other forms
with the `OPEN WINDOW` name `WITH FORM`
instruction:

```
MAIN
  DEFER INTERRUPT
  OPTIONS INPUT WRAP
   ...
  OPEN FORM f_main FROM "custfrm"
  DISPLAY FORM f_main 
   ...
END MAIN
```

The `SCREEN` window is not visible in TUI mode because
program windows are rendered as simple boxes and `SCREEN` is
created without borders. The size of the `SCREEN` window
is 80x25 in TUI mode.

## Related links

**Related concepts**  

[OPEN WINDOW](../11_user-interface/1572-open-window.md "Creates and displays a new window.")
