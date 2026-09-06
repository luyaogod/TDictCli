---
title: "Using regular DISPLAY ARRAY control blocks"
source: "fgl-topics/c_fgl_treeviews_009.html"
breadcrumb: "User interface > User interface programming > Tree views > Using regular DISPLAY ARRAY control blocks"
type: "concept"
---

# Using regular DISPLAY ARRAY control blocks

> Simple table DISPLAY ARRAY control blocks can be used with tree-views.

If needed, you can implement traditional `DISPLAY ARRAY` control blocks like
`BEFORE ROW` or `AFTER
ROW`:

```
DISPLAY ARRAY tree_arr TO sr.* ATTRIBUTES(UNBUFFERED)
  BEFORE ROW
    DISPLAY "BEFORE ROW - Current row is: ", DIALOG.getCurrentRow("sr")
  AFTER ROW
    DISPLAY "AFTER ROW  - Current row is: ", DIALOG.getCurrentRow("sr")
END DISPLAY
```

## Related links

**Related concepts**  

[DISPLAY ARRAY control blocks](1971-display-array-control-blocks.md "DISPLAY ARRAY control blocks")

[DIALOG control blocks](2099-dialog-control-blocks.md "Dialog control blocks are predefined dialog triggers where you can implement specific code to control the interactive instruction.")

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
