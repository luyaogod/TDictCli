---
title: "CLEAR FORM"
source: "fgl-topics/c_fgl_record_display_CLEAR_FORM.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > CLEAR FORM"
type: "concept"
---

# CLEAR FORM

> The CLEAR FORM instruction clears all fields in the current form.

## Syntax

```
CLEAR FORM
```

## Usage

The `CLEAR FORM` instruction clears all form fields of the current form. It has
no effect on any part of the screen display except the form fields.

> **Important:**
>
> Unlike `CLEAR
> field-list`, the `CLEAR FORM` instruction does not
> set the [field modification flags](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.").

In a similar way to `CLEAR
field-list`, the `CLEAR FORM` instruction is
typically used when the program is outside a dialog block execution controlling the form fields. For
example, after a database query with a [`CONSTRUCT`](2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.") instruction, you might want to clear all search criteria entered
by the user with this instruction, to cleanup the form.

The `CLEAR FORM` instruction clears the field values and resets the TTY attributes
to `NORMAL`.

The `CLEAR FORM` instruction is not needed, if the program is always in the
context of a [dialog controlling the form
fields](2219-the-model-view-controller-paradigm.md "The dynamic user interface architecture is based on the Model-View-Controller (MVC) paradigm.").

## Example

```
  CONSTRUCT BY NAME sql
     ON cust_name, cust_address, ...
     ...
  END CONSTRUCT
  CLEAR FORM
  ...
```

## Related links

**Related concepts**  

[CLEAR WINDOW](1577-clear-window.md "Clears the contents of a window.")

[CLEAR SCREEN](1581-clear-screen.md "Clears the complete application screen.")
