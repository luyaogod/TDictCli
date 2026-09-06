---
title: "Function key mapping"
source: "fgl-topics/c_fgl_Mig0000_023.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics > Function key mapping"
type: "concept"
---

# Function key mapping

> When migrating to Genero BDL, special care must be taken for programs that uses function keys greater than F12.

With Four Js Business Development Suite (BDS), when the user pressed a key modifier plus a
function key (like `Shift-F4` or `Ctrl-F6`), the key combination was
mapped to a regular function key `F(nn+offset)`, because `Shift` and
`Control` key modifiers are not handled in the 4GL language.

The number of function keys of the keyboard was defined by the
`gui.key.add_function` FGLPROFILE entry. For example, when this entry is set to 12
(the default), a Shift-F4 was received as F16 (4 + 12) in
the program, to be handled with the `ON KEY (F16)` clause.

This feature and FGLPROFILE entry is still supported when using the traditional mode.

When using the standard GUI mode, special consideration needs to be taken regarding function keys
above F12.

Function keys from F1 to F12 are common keys found on the keyboard and do not need any particular
configuration regarding accelerators: The `ON KEY` clause defines an action object
identified by the key name (in lowercase) and the first accelerator attribute defined with the same
name. For example, `ON KEY (F10)` creates an action f10, with
accelerator F10.

However, if the program uses `ON KEY (Fnn)` clauses where nn is above 12, in order
to have `Shift-F(nn-12)` key combinations working, you need to define this
accelerator in the corresponding action default entry.

> **Important:**
>
> Since the `ON KEY (Fnn)` clause defines an automatic shortcut
> key with the first accelerator attribute, it overwrites the value of
> `acceleratorName` attribute defined in the action defaults. In order to associate the
> actual `Shift-F(nn-12)` key combination for the action `fnn`, you must
> use the second accelerator attribute. For example: `ON KEY (F14)` creates an action
> f14, with first accelerator F14. Specify
> `acceleratorName2="Shift-F2"` in the .4ad action defaults file,
> to define the `Shift-F2` key combination for the f14
> action.

## Related links

**Related concepts**  

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Graphical mode with Traditional Display](../11_user-interface/1519-graphical-mode-with-traditional-display.md "Graphical mode with Traditional Display")
