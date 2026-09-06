---
title: "scr_line()"
source: "fgl-topics/c_fgl_BuiltInFunctions_SCR_LINE.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > scr_line()"
type: "concept"
---

# scr_line()

> Returns the index of the current row in the form screen array.

## Syntax

```
FUNCTION scr_line()
  RETURNS INTEGER
```

## Usage

The `scr_line()` function returns the index of the current row in the current
screen array of the form. It is typically used inside a `DISPLAY ARRAY` or
`INPUT ARRAY` statement, especially in TUI mode.

Do not confuse `scr_line()` with `arr_curr()`; the first returns the index
of the current row in the form screen array, and the second returns the index of the current row in
the program variable.

> **Important:**
>
> With new graphical objects such as [`TABLE`](../11_user-interface/1724-table-container.md "Defines a re-sizable table designed to display a list of records.") container, the `scr_line()` function can return a
> screen array line number outside the visible page of rows, because the current row may not be in the
> visibile rows when scrolling the list. Consequently, legacy code relying on
> `scr_line()` to display data in specific screen array lines can fail. For TUI
> applications, the code can be left untouched. For GUI applications, consider using [`DIALOG.getCurrentRow()`](3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.") or
> `arr_curr()`, with the [`UNBUFFERED`](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") mode.

## Related links

**Related concepts**  

[Screen records / arrays](../11_user-interface/1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")
