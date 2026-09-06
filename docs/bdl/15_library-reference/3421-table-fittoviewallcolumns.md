---
title: "table.fitToViewAllColumns"
source: "fgl-topics/c_fgl_frontcall_table_fittoviewallcolumns.html"
breadcrumb: "Library reference > Built-in front calls > Table front calls > table.fitToViewAllColumns"
type: "concept"
---

# table.fitToViewAllColumns

> Adapts the width of table columns to show all columns.

## Syntax

```
ui.Interface.frontCall("table", "fitToViewAllColumns",
  [screen-record],
  []
)
```

1. screen-record - This is the name of the screen record / table.

## Usage

This front call resizes the columns of a `TABLE` container, to show all columns in
the available space of the window.

The screen-record argument identifies the `TABLE` container to
be re-organized.

> **Important:**
>
> This front call may be called before the layout is fully rendered. As a result, if incorrect
> arguments are provided, no error popup will appear (unlike other front calls). Instead, errors will
> be shown in the browser debug console.
