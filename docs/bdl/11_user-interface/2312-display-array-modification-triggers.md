---
title: "DISPLAY ARRAY modification triggers"
source: "fgl-topics/c_fgl_prog_dialogs_modif_rolist.html"
breadcrumb: "User interface > User interface programming > List dialogs > DISPLAY ARRAY modification triggers"
type: "concept"
---

# DISPLAY ARRAY modification triggers

> Using dedicated interaction blocks to allow the user to modify a read-only record list.

The `DISPLAY ARRAY` block implements by default
a read-only list of records. The end user can navigate in
the list, but cannot modify the rows.

The traditional way to implement an editable list of record is to use
`INPUT ARRAY`. However, `INPUT ARRAY` uses
ergonomics that may not correspond to the end user expectations.
Basically, a list controlled by an `INPUT ARRAY` is always
in "edit mode": the focus is in a field and the user can modify the current
field. When moving up or down in the list, the edit cursor jumps to the
upper or lower cell.

Other GUI applications use a different pattern, with read-only lists
that can switch to edit mode when a specific action is fired.
To implement such ergonomics, use the `ON INSERT`,
`ON APPEND`, `ON UPDATE`,
`ON DELETE` modification triggers to control row insertion,
appending, modification and deletion in a `DISPLAY ARRAY` block.

## Related links

**Related concepts**  

[Example 3: DISPLAY ARRAY using modification triggers](2003-example-3-display-array-using-modification-triggers.md "Example 3: DISPLAY ARRAY using modification triggers")
