---
title: "SCROLL"
source: "fgl-topics/c_fgl_record_display_SCROLL.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > SCROLL"
type: "concept"
---

# SCROLL

> The SCROLL instruction moves data rows up or down in a screen array.

## Syntax

```
SCROLL field-list { UP | DOWN } [ BY lines ]
```

where field-list is:

```
{ field-name
| table-name.*
| table-name.field-name
| screen-array[line].*
| screen-array[line].field-name
| screen-record.*
| screen-record.field-name
} [,...]
```

1. field-name is the identifier of a field of the current form.
2. table-name is the identifier of a database table of the current form.
3. screen-record is the identifier of a screen record of the current form.
4. screen-array is the name of the screen array used of the current form.
5. line is the line number in the screen array (it is ignored by
   `SCROLL`).
6. lines is an integer expression that specifies how far (in lines) to scroll
   the display. Passing zero has no effect.

## Usage

The `SCROLL` instruction specifies vertical movements of displayed values in
all or some of the fields of a screen array within the current form.

The fields to be scrolled can be specified individually or by referencing a [screen record or screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."), with the
`.*` notation to specify all fields.

The `SCROLL` instruction syntax supports a
`screen-array[line]` specification. In this
form, the *line* number will be ignored: With `SCROLL`, all rows of a screen
array are shifted.

The `SCROLL` instruction is supported for applications running in TUI mode, to
scroll screen array rows when no interactive instruction is executing. In a GUI application, use
a `TABLE` container with a [`DISPLAY
ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.") instruction.

When passing zero as offset in `SCROLL ... UP/DOWN BY lines`,
the instruction does nothing.

## Related links

**Related concepts**  

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")

[Form fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.")

[Windows and forms](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.")
