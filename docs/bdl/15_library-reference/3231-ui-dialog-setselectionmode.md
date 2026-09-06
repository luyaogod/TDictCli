---
title: "ui.Dialog.setSelectionMode"
source: "fgl-topics/c_fgl_ClassDialog_setSelectionMode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setSelectionMode"
type: "concept"
---

# ui.Dialog.setSelectionMode

> Defines the row selection mode for the specified list.

## Syntax

```
setSelectionMode(
   name STRING,
   mode INTEGER )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. mode defines the selection mode (0, 1).

## Usage

In [`DISPLAY ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.") instructions,
the `setSelectionMode()` method can be used to enable/disable
multi-row selection.

Possible values of the mode parameter are 0 (single row selection) or 1
(multi-range selection). Other values are reserved for future use.

If multi-row selection is switched off, selected rows get deselected.

For
more details about multi-row selection, see [Multiple row selection](../11_user-interface/2314-multiple-row-selection.md "Multiple row selection allows the end user to select several rows within a list of records.").
