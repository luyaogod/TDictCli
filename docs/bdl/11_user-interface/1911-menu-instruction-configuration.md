---
title: "MENU instruction configuration"
source: "fgl-topics/c_fgl_menus_008.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > MENU instruction configuration"
type: "concept"
description: "The rendering and behavior of a MENU instruction can be configured with the ATTRIBUTES clause: MENU \"Question\" ATTRIBUTES ( STYLE=\"dialog\", COMMENT=\"Do you want to commit your changes?\" ) When the ..."
---

# MENU instruction configuration

The rendering and behavior of a `MENU` instruction can be configured
with the `ATTRIBUTES` clause:

```
MENU "Question"
     ATTRIBUTES (
         STYLE="dialog",
         COMMENT="Do you want to commit your changes?"
      )
```

When the `STYLE` dialog attribute is set to '`default`'
or when you do not specify the menu type, the runtime system generates a default decoration
as a set of buttons in a specific area of the current window.

When the `STYLE` attribute is set to '`dialog`', the menu
options appear as buttons at the bottom in a temporary modal window, in which you can
define the message and the icon with the `COMMENT` and `IMAGE`
attributes.

When the `STYLE` is set to '`popup`', the menu appears as
a pop-up menu (a context menu).

If the menu `STYLE` is "`dialog`" or "`popup`",
the dialog is automatically exited after any action clause such as `ON ACTION`,
`COMMAND`, `ON IDLE` or `ON TIMER`.

## Related links

**Related concepts**  

[Rendering modes of a menu](1909-rendering-modes-of-a-menu.md "Rendering modes of a menu")

[Syntax of the MENU instruction](1906-syntax-of-the-menu-instruction.md "The MENU instruction defines a set of options the end user can select to trigger actions in a program.")
