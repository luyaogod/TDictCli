---
title: "Defining default TTY attributes"
source: "fgl-topics/c_fgl_programs_007.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Defining default TTY attributes"
type: "concept"
---

# Defining default TTY attributes

> The OPTIONS {INPUT|DISPLAY} ATTRIBUTES defines default TTY attributes for dialogs and display statements.

## Syntax

```
OPTIONS {
  INPUT {ATTRIBUTE|ATTRIBUTES} ( { FORM | WINDOW | attributes } )
| DISPLAY {ATTRIBUTE|ATTRIBUTES} ( { FORM | WINDOW | attributes } )
}
```

## Usage

`OPTIONS INPUT ATTRIBUTES` defines the default color and
terminal effect attributes that will be used in subsequent dialog statement.

`OPTIONS DISPLAY ATTRIBUTES` defines the
default attributes for display statements.

The display attributes are based on dumb terminal (i.e. TTY) possibilities, but will be rendered
accordingly on GUI mode. Graphical front-ends can be configured to render TTY attributes is a
specific way. Instead of TTY based attributes, consider using [presentation styles](../11_user-interface/1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.") in new developments.

Any display attribute defined by the `OPTIONS` statement
remains in effect until the runtime system encounters a statement that redefines
the same attribute. This can be another `OPTIONS` statement,
or an `ATTRIBUTE` clause in one of the following statements:

- [`CONSTRUCT`](../11_user-interface/2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")
- [`INPUT`](../11_user-interface/1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.")
- [`DISPLAY`](../11_user-interface/1881-display-by-name.md "The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names.")
- [`DIALOG`](../11_user-interface/2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.")
- [`INPUT ARRAY`](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")
- [`DISPLAY ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")
- [`OPEN WINDOW`](../11_user-interface/1572-open-window.md "Creates and displays a new window.")

The `ATTRIBUTE` clause in these statements
only redefines the attributes temporarily. After the window
closes or after the dialog statement terminates, the runtime
system restores the attributes from the most recent `OPTIONS` statement.

The `FORM` keyword in `INPUT ATTRIBUTE` or `DISPLAY
ATTRIBUTE` clauses instructs the runtime system
to use the input or display attributes of the current form.
Similarly, you can use the `WINDOW` keyword
of the same clauses to instruct the program to use the input or
display attributes of the current window. You cannot combine
the `FORM` or `WINDOW` attributes
with any other attributes.

This table shows the valid input and display attributes:

| Attribute | Description |
| --- | --- |
| `BLACK, BLUE, CYAN, GREEN, MAGENTA, RED, WHITE, YELLOW` | The TTY color of the displayed text. |
| `BOLD, DIM, INVISIBLE, NORMAL` | The TTY font attribute of the displayed text. |
| `REVERSE, BLINK, UNDERLINE` | The TTY video attribute of the displayed text. |
