---
title: "ON KEY / COMMAND vs ON ACTION"
source: "fgl-topics/c_fgl_MigI4GL_039.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > ON KEY / COMMAND vs ON ACTION"
type: "concept"
description: "IBM Informix 4GL applications use the ON KEY and COMMAND [KEY] clauses to handle user actions."
---

# ON KEY / COMMAND vs ON ACTION

> IBM® Informix® 4GL applications use the ON KEY and COMMAND [KEY] clauses to handle user actions.

The `ON KEY (key-name)`, `COMMAND
"cmd-name"` and `COMMAND KEY(key-name)`
instructions are legacy clauses used in dialog instructions, to execute code when the end user
presses the corresponding key on the keyboard, or when the end user selects an option in a
`MENU` instruction.

The `ON KEY` / `COMMAND` clauses are fully supported by Genero BDL
and existing code can be left untouched. However, for more flexiblity, BDL provides an abstract
solution to implement a user action handler, with the `ON ACTION
action-name` clause. In its basic form, the `ON ACTION`
clause only specifies an action name, to identify a given user action such as "accept", "cancel",
"print", insead of `F6`, `CTRL-C` and `ESCAPE` keys
with `ON KEY`, or `COMMAND` labels, which are language-specific.

`ON ACTION` decoration (labels, icons) and keyboard accelerators can then be
defined and centralized in external program files such as .4ad action defaults
files.

Using asbtract action names also simplifies application text localization.

## Related links

**Related concepts**  

[ON ACTION block](../11_user-interface/1918-on-action-block.md "ON ACTION block")
