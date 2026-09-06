---
title: "Tree-view predefined actions"
source: "fgl-topics/c_fgl_treeviews_predef_actions.html"
breadcrumb: "User interface > User interface programming > Tree views > Tree-view predefined actions"
type: "concept"
---

# Tree-view predefined actions

> Several predefined actions are implicitly available with a DISPLAY ARRAY controlling a TREE.

When a `DISPLAY ARRAY` controls a `TREE` container, the following
actions are automatically created by the dialog:

- `collapseall`: Closes completely the current tree-view node.
- `expandall`: Opens the current tree-view node and all child nodes
  recursively.

These predefined actions can be configured with action defaults attributes, and can be bound to
action views such as toolbar buttons:

```
ACTION DEFAULTS
ACTION collapseall (IMAGE="fa-compress", ACCELERATOR=F10 )
ACTION expandall (IMAGE="fa-expand", ACCELERATOR=F11 )
END

TOOLBAR
ITEM collapseall
ITEM expandall
END
```

Do not create action handlers (`ON ACTION`) for these predefined actions: The
dialog implements the code to perform these actions.

## Related links

**Related concepts**  

[Action handling basics](2254-action-handling-basics.md "This topic describes the basic concepts of dialog actions.")

[Predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.")
