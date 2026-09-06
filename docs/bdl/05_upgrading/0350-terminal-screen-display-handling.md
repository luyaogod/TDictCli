---
title: "Terminal screen display handling"
source: "fgl-topics/c_fgl_MigI4GL_052.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Terminal screen display handling"
type: "concept"
---

# Terminal screen display handling

> Compared to I4GL, Genero BDL uses an optimized and specific design to handle terminal display, that introduces some differences.

## TERMCAP/TERMINFO

Genero BDL uses termcap/terminfo based solutions and therefore supports all kind of terminals the
IBM® Informix® 4GL
supports.

For more details, see [Using a text terminal](../11_user-interface/1531-using-a-text-terminal.md "This section covers topics about text terminal configuration when using the TUI mode.").

## Remaining DISPLAY AT text

With IBM Informix 4GL, displaying text to the current window is a direct rendering method, and no trace
of the displayed text is keep. When chaining interactive instructions and various forms with code
using [`DISPLAY AT`](../11_user-interface/1582-display-at.md "The DISPLAY ... AT instruction displays text at a given line/column position in the current window.")
instructions, it is possible that some phantom text remains with I4GL when a previous window/form is
re-activated.

Genero BDL implements the user interface with the concept of [Abstract User Interface](../11_user-interface/1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user."), a XML tree that keeps
the elements that have been displayed by the program. Consequently, a text shown with a
`DISPLAY AT` instruction will remain, even after closing a window/form that was
hiding this text. With I4GL, the text disappears.

Regarding `DISPLAY AT`, the behavior of Genero BDL is more consistent than
I4GL.

## TTY attributes not reused in new INPUT

With IBM Informix 4GL, an `INPUT` dialog without `ATTRIBUTES` clause
defining TTY attributes, reuses the TTY attributes that were set by a previous dialog. For example,
when an `OPEN WINDOW` defines TTY attributes, then an `INPUT` dialog
without TTY attributes is executed, followed by another `INPUT` with TTY attributes,
then again by the `INPUT` without TTY attributes: The last `INPUT`
uses the TTY attributes defined the second
`INPUT`:

```
OPEN WINDOW w1 AT 2,2 WITH FORM "form1"
     ATTRIBUTES(BORDER,BLUE,UNDERLINE)
...
INPUT BY NAME rec.* WITHOUT DEFAULTS
...
INPUT BY NAME rec.* WITHOUT DEFAULTS ATTRIBUTES(RED,REVERSE)
...
INPUT BY NAME rec.* WITHOUT DEFAULTS -- gets RED,REVERSE
...
```

With Genero BDL, any new `INPUT` without TTY attributes always uses the attributes
defined by the `OPEN WINDOW`. This behavior is more consistent than I4GL.

If needed, review the program logic and explicitely use the TTY attribute you want directly in
the `INPUT` dialog.
