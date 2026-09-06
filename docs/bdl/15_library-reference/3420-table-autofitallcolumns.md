---
title: "table.autoFitAllColumns"
source: "fgl-topics/c_fgl_frontcall_table_autofitallcolumns.html"
breadcrumb: "Library reference > Built-in front calls > Table front calls > table.autoFitAllColumns"
type: "concept"
---

# table.autoFitAllColumns

> Adapts the width of table columns to the displayed data.

## Syntax

```
ui.Interface.frontCall("table", "autoFitAllColumns",
  [screen-record],
  []
)
```

1. screen-record - This is the name of the screen record / table.

## Usage

This front call resizes the columns of a `TABLE` container, to make the column
contents fully visible.

The screen-record argument identifies the `TABLE` container to
be re-organized.

> **Important:**
>
> This front call may be called before the layout is fully rendered. As a result, if incorrect
> arguments are provided, no error popup will appear (unlike other front calls). Instead, errors will
> be shown in the browser debug console.
