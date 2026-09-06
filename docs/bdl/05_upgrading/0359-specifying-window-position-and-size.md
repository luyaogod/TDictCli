---
title: "Specifying WINDOW position and size"
source: "fgl-topics/c_fgl_MigI4GL_016.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Specifying WINDOW position and size"
type: "concept"
---

# Specifying WINDOW position and size

> With Genero BDL in GUI mode, window position and sizes are ignored; in TUI mode, window position and sizes are respected.

When writing a program for TUI mode, the windows can be created
with the `OPEN WINDOW name AT x,y` instruction,
specifying an position on the screen; sometimes even the
width and height of the window is specified, for example when you
don't use a form to create the window. Window position
and size is allowed by Genero Business Development Language
for TUI mode applications.

However, the window position and sizes are ignored in GUI mode.
In GUI mode, the window position is defined by the window
manager, and the size adapts to the form displayed. In this
mode, the preferred way to display application forms is to use
the `OPEN WINDOW` name `WITH
FORM` instruction.

## Related links

**Related concepts**  

[OPEN WINDOW](../11_user-interface/1572-open-window.md "Creates and displays a new window.")
