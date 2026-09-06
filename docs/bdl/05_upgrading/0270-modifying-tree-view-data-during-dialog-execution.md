---
title: "Modifying tree view data during dialog execution"
source: "fgl-topics/c_fgl_Migrate_to_230_modif_treeview.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide > Modifying tree view data during dialog execution"
type: "concept"
---

# Modifying tree view data during dialog execution

> Use ui.Dialog methods to insert/append/delete treeview nodes.

Before version 2.30, it was possible to use the `insertRow()` /
`appendRow()` / `deleteRow()` / `deleteAllRows()`
dialog class methods to modify the tree array during the dialog execution. But these methods were
not designed to handle tree data properly. An alternative was to use program array methods, but
when modifying the program array directly, multi-range selection flags or cell attributes were
not synchronized.

Starting with 2.30.02, you can now use the `insertNode()`,
`appendNode()` and `deleteNode()` methods of the
`ui.Dialog` class. You can still directly fill the program array before
the dialog execution, but it is recommended to use dialog methods during the dialog
execution.

## Related links

**Related concepts**  

[The Dialog class](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.")

[Tree views](../11_user-interface/2352-tree-views.md "Describes how to implement tree views.")
