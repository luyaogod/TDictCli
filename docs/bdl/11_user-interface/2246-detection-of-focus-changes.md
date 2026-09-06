---
title: "Detection of focus changes"
source: "fgl-topics/c_fgl_prog_dialogs_focus_detect.html"
breadcrumb: "User interface > User interface programming > Input fields > Detection of focus changes"
type: "concept"
---

# Detection of focus changes

> Describes how to detect when the focus goes from field to field or to a read-only list.

## Detecting focus changes in a singular INPUT or CONSTRUCT

An singular `INPUT` or `CONSTRUCT` controls several fields that can
get the focus and become current. In order to execute some code when a field gets (or loses) the
focus, use the following control blocks:

- [`BEFORE FIELD`](1942-before-field-block.md) (a specific
  field (or group of fields) gets the focus)
- [`AFTER FIELD`](1944-after-field-block.md) (the field (or
  group of fields) loses focus)

## Detecting focus changes in a singular DISPLAY ARRAY

A singular `DISPLAY ARRAY` controls rows of a list, that can get the focus and
become current. In order to execute some code when a row gets (or loses) the focus,
use the following control blocks:

- [`BEFORE ROW`](1975-before-row-block.md) (a new row gets the
  focus inside a `DISPLAY ARRAY` or `INPUT ARRAY` list)
- [`AFTER ROW`](1976-after-row-block.md) (a row inside a
  `DISPLAY ARRAY` or `INPUT ARRAY` list loses focus)

## Detecting focus changes in a singular INPUT ARRAY

An singular `INPUT ARRAY` controls several fields and rows of a list, that can get
the focus and become current. In order to execute some code when a field or a row gets (or loses)
the focus, use the following control blocks:

- [`BEFORE ROW`](1975-before-row-block.md) (a new row gets the
  focus inside a `DISPLAY ARRAY` or `INPUT ARRAY` list)
- [`BEFORE FIELD`](1942-before-field-block.md) (a specific
  field (or group of fields) gets the focus)
- [`AFTER FIELD`](1944-after-field-block.md) (the field (or
  group of fields) loses focus)
- [`AFTER ROW`](1976-after-row-block.md) (a row inside a
  `DISPLAY ARRAY` or `INPUT ARRAY` list loses focus)

## Detecting focus changes in a DIALOG

A `DIALOG` interaction block can handle different parts of a form
simultaneously. In order to execute some code when a part of the form gets (or
loses) the focus, use the following control blocks:

- [`BEFORE INPUT`](1940-before-input-block.md)
  (a field of this `INPUT` or `INPUT ARRAY`
  sub-dialog gets the focus and none of its fields had focus before)
- [`BEFORE
  CONSTRUCT`](2057-before-construct-block.md) (a field of this `CONSTRUCT`
  sub-dialog gets the focus and none of its fields had focus before)
- [`BEFORE
  DISPLAY`](1973-before-display-block.md) (this `DISPLAY ARRAY` sub-dialog gets
  the focus and none of its fields had focus before)
- [`BEFORE ROW`](1975-before-row-block.md) (a
  new row gets the focus inside a `DISPLAY ARRAY` or `INPUT
  ARRAY` list)
- [`BEFORE FIELD`](1942-before-field-block.md)
  (a specific field (or group of fields) gets the focus)
- [`AFTER FIELD`](1944-after-field-block.md)
  (the field (or group of fields) loses focus)
- [`AFTER ROW`](1976-after-row-block.md) (a
  row inside a `DISPLAY ARRAY` or `INPUT ARRAY` list
  loses focus)
- [`AFTER
  DISPLAY`](1974-after-display-block.md) (this `DISPLAY ARRAY` sub-dialog
  loses the focus = focus goes to another sub-dialog)
- [`AFTER
  CONSTRUCT`](2058-after-construct-block.md) (this `CONSTRUCT` sub-dialog loses
  the focus = focus goes to another sub-dialog)
- [`AFTER INPUT`](1941-after-input-block.md)
  (this `INPUT` or `INPUT ARRAY` sub-dialog loses
  focus = focus goes to another sub-dialog)

These triggers are also executed by [`NEXT FIELD`](1955-next-field-instruction.md).

## Related links

**Related concepts**  

[Which form item has the focus?](2244-which-form-item-has-the-focus.md "Identify what element of the current form has the focus.")

[Field-level focus in DISPLAY ARRAY](2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.")

[DIALOG control blocks](2099-dialog-control-blocks.md "Dialog control blocks are predefined dialog triggers where you can implement specific code to control the interactive instruction.")
