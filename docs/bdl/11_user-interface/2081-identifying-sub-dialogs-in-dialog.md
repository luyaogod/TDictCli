---
title: "Identifying sub-dialogs in DIALOG"
source: "fgl-topics/c_fgl_prog_dialogs_subdialog_ident.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > Identifying sub-dialogs in DIALOG"
type: "concept"
---

# Identifying sub-dialogs in DIALOG

> Sub-dialogs need to be identified by a name to distinguish the different contexts.

A procedural `DIALOG` block is a collection of sub-dialogs that act as controllers
for different parts of a form. In order to program a procedural `DIALOG` block,
there must be a unique identifier for each sub-dialog.

For example, to set the current row of a screen array with the
`DIALOG.setCurrentRow()` method, you pass the name of the screen
array to specify the sub-dialog to be affected. Sub-dialog identifiers are also
used as a prefix to specify actions for the sub-dialog.

The following topics describe how to specify the names of the different types
of `DIALOG` sub-dialogs:

- [Identifying an INPUT sub-dialog](2083-the-input-sub-dialog.md)
- [Identifying a DISPLAY ARRAY sub-dialog](2085-the-display-array-sub-dialog.md)
- [Identifying an INPUT ARRAY sub-dialog](2086-the-input-array-sub-dialog.md)
- [Identifying a CONSTRUCT sub-dialog](2084-the-construct-sub-dialog.md)
- [The SUBDIALOG clause](2087-the-subdialog-clause.md).

## Related links

**Related concepts**  

[Structure of a procedural DIALOG block](2082-structure-of-a-procedural-dialog-block.md "Structure of a procedural DIALOG block")

[The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.")

[Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")
