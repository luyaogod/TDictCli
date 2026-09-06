---
title: "DISPLAY FORM"
source: "fgl-topics/c_fgl_windows_and_forms_DISPLAY_FORM.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > DISPLAY FORM"
type: "concept"
---

# DISPLAY FORM

> Displays and associates a form with the current window.

## Syntax

```
DISPLAY FORM identifier
[ {ATTRIBUTE|ATTRIBUTES} ( display-attributes ) ]
```

1. identifier is the name of the form.
2. window-attributes defines the display attributes of the form.

where display-attribute is:

```
{ BLACK | BLUE | CYAN | GREEN
| MAGENTA | RED | WHITE | YELLOW
| BOLD | DIM | INVISIBLE | NORMAL
| REVERSE | BLINK | UNDERLINE
}
```

## Usage

The `DISPLAY FORM` instruction creates a form element in the current window,
from a form resource loaded by the [`OPEN
FORM`](1578-open-form.md "Declares a compiled form in the program.") instruction.

> **Important:**
>
> The `INVISIBLE` display attribute is ignored.

The runtime system applies display attributes that you specify in the `ATTRIBUTES`
clause, to any fields that have not been assigned attributes by the [`ATTRIBUTES`](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.") section of
the form specification file, or by the database schema files, or by the `OPTIONS` runtime configuration statement.
If the form is displayed in a window, color attributes from the `DISPLAY FORM`
statement supersede any from the `OPEN WINDOWOPEN
WINDOW` statement. If however subsequent `CONSTRUCT`, `DISPLAY`,
or `DISPLAY ARRAY` statements that
include an `ATTRIBUTES` clause reference the form, their attributes take precedence
over those specified in the `DISPLAY FORM` instruction.

> **Important:**
>
> In GUI mode, form elements can also be decorated with
> presentation styles. Pay attention to the specific rules that apply when [combining TTY attributes and presentation
> styles](1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.").

In graphical mode, by default, the parent window adapts its size to the content of the form
displayed with `DISPLAY FORM`.

## Related links

**Related concepts**  

[CLOSE FORM](1580-close-form.md "Closes the resources allocated by OPEN FORM.")
