---
title: "Field-level focus in DISPLAY ARRAY"
source: "fgl-topics/c_fgl_ui_tables_focusonfield.html"
breadcrumb: "User interface > User interface programming > List dialogs > Field-level focus in DISPLAY ARRAY"
type: "concept"
---

# Field-level focus in DISPLAY ARRAY

> The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.

## Enabling focusable cells in `DISPLAY ARRAY`

When using a [`DISPLAY ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")
dialog to control a list view, you can enable cell-level focus handling with the [`FOCUSONFIELD`](1965-display-array-instruction-configuration.md) attribute.

With a graphical front-end, this feature allows cell mouse clicks or tabbing between cells,
instead of having the whole current row highlighted.

This feature can also be used in text mode (FGLGUI=0), to move in `DISPLAY ARRAY`
fields with the keyboard using arrow keys or tab.

To enable field-level focus handling in a `DISPLAY ARRAY`, add the
`FOCUSONFIELD` attribute in the dialog definition:

```
DISPLAY ARRAY arr TO sr.* ATTRIBUTES(FOCUSONFIELD)
```

When using the `FOCUSONFIELD` option, the `DISPLAY ARRAY` will
automatically create the `left` and `right` [predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions."), to navigate in the cells
with the left/right keys.

## Detecting cell focus changes

When the `FOCUSONFIELD` attribute is defined, [`BEFORE FIELD`](1942-before-field-block.md) and [`AFTER FIELD`](1944-after-field-block.md) blocks can be used, to
detect field focus
changes:

```
DISPLAY ARRAY arr TO sr.* ATTRIBUTES(FOCUSONFIELD)
   ...
   BEFORE FIELD cust_id
      MESSAGE "focus in cust_id, row = ", arr_curr()
   AFTER FIELD cust_id
      MESSAGE "focus left cust_id, row = ", arr_curr()
   BEFORE FIELD cust_name
      MESSAGE "focus in cust_name, row = ", arr_curr()
   ...
END DISPLAY
```

Note that if defined, the `AFTER ROW` control block will execute after the
`AFTER FIELD` block and the `BEFORE ROW` control block will execute
before the `BEFORE FIELD` block. The code blocks execute in the following order:

1. `AFTER FIELD` (for the field that loses the focus)
2. `AFTER ROW` (for the previous current row)
3. `BEFORE ROW` (for the new current row)
4. `BEFORE FIELD` (for the field that gets the focus in the new row)

## What is the current cell?

The current cell of `DISPLAY ARRAY` with `FOCUSONFIELD` attribute
can be found by using the [`ui.Dialog.getCurrentItem()`](../15_library-reference/3194-ui-dialog-getcurrentitem.md "Returns the current item having focus.") method, in conjunction with the [`ui.Dialog.getCurrentRow()`](../15_library-reference/3195-ui-dialog-getcurrentrow.md "Returns the current row of the specified list.")
method:

```
DISPLAY ARRAY arr TO sr.* ATTRIBUTES(FOCUSONFIELD)
   ...
    ON ACTION show_current_cell
        MESSAGE SFMT("Current cell: field=%1 / row=%2",
                     DIALOG.getCurrentItem(),DIALOG.getCurrentRow("sr"))
```

## Setting the current cell

To set the focus to a specific cell with program code, use the [`NEXT FIELD`](2245-giving-the-focus-to-a-form-element.md "How to force the focus to move or stay in a specific form element using program code.")
instruction, or the [`ui.Dialog.nextField()`](../15_library-reference/3210-ui-dialog-nextfield.md "Registers the next field to go to.") method, in conjunction with the
[`ui.Dialog.setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.")
method:

```
DISPLAY ARRAY arr TO sr.* ATTRIBUTES(FOCUSONFIELD)
   ...
    ON ACTION top_left ATTRIBUTES(TEXT = "TOP LEFT")
        CALL DIALOG.setCurrentRow("sr", 1)
        NEXT FIELD first_field
    ON ACTION bottom_right ATTRIBUTES(TEXT = "BOTTOM RIGHT")
        CALL DIALOG.setCurrentRow("sr", sr.getLength())
        NEXT FIELD last_field
   ...
END DISPLAY
```

## Related links

**Related concepts**  

[DISPLAY ARRAY control blocks](1971-display-array-control-blocks.md "DISPLAY ARRAY control blocks")
