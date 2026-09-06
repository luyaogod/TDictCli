---
title: "DISPLAY AT"
source: "fgl-topics/c_fgl_windows_and_forms_DISPLAY_AT_2.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > DISPLAY AT"
type: "concept"
---

# DISPLAY AT

> The DISPLAY ... AT instruction displays text at a given line/column position in the current window.

## Syntax

```
DISPLAY text AT line, column [ {ATTRIBUTE|ATTRIBUTES} ( display-attributes ) ]
```

1. text is any expression to be evaluated and displayed at the given position
   in the current window.
2. line is an integer expression defining the line position in the current
   window.
3. column is an integer expression defining the column position on the screen.
4. display-attributes defines the display attributes for the
   text.

## Usage

The `DISPLAY AT` instruction evaluates a string expression and displays the result
at a given line and column in the current window. This instruction is typically used in text user
interface (TUI) text-based applications to display static text on the screen such as messages or
decoration lines with `-` (hyphen) or `_`(underscore) characters.

Use of the `DISPLAY AT` instruction is recommended only in TUI mode or GUI/traditional
mode. To display data at a given place in a graphical form, use form fields and the
`DISPLAY BY NAME` or
`DISPLAY TO` instructions, or use interactive instructions with the [`UNBUFFERED`](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") mode to automatically
display program variable data to form fields.

When using `DISPLAY AT` in GUI mode, the text will only be displayed if the
current window contains no form, or contains a form defined with the [`SCREEN`](1714-screen-section.md "The SCREEN section defines the form layout for TUI mode forms.") layout.

| Attribute | Description |
| --- | --- |
| `BLACK, BLUE, CYAN, GREEN, MAGENTA, RED, WHITE, YELLOW` | The TTY color of the displayed text. |
| `BOLD, DIM, INVISIBLE, NORMAL` | The TTY font attribute of the displayed text. |
| `REVERSE, BLINK, UNDERLINE` | The TTY video attribute of the displayed text. |
