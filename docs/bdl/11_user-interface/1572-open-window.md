---
title: "OPEN WINDOW"
source: "fgl-topics/c_fgl_windows_and_forms_OPEN_WINDOW.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Instructions for windows and forms > OPEN WINDOW"
type: "concept"
---

# OPEN WINDOW

> Creates and displays a new window.

## Syntax

```
OPEN WINDOW identifier
   [ AT line, column ]
   WITH { FORM form-file
        | height ROWS, width COLUMNS
        }
   [ {ATTRIBUTE|ATTRIBUTES} ( window-attribute [,...] ) ]
```

where
window-attribute
is:

```
{ BLACK | BLUE | CYAN | GREEN
| MAGENTA | RED | WHITE | YELLOW
| BOLD | DIM | INVISIBLE | NORMAL
| REVERSE | BLINK | UNDERLINE
| BORDER
| TEXT = "string"
| STYLE = "string"
| PROMPT LINE integer
| MENU LINE integer
| MESSAGE LINE integer
| ERROR LINE integer
| COMMENT LINE { OFF | integer }
}
```

1. identifier is the name of the window. It is always converted to lowercase by
   the compiler.
2. line is the integer defining the top position of the window. The first line
   in the screen is 1, while the relative line number inside the window is zero.
3. column is the integer defining the position of the left margin. The first
   column in the screen is 1, while the relative column number inside the window is zero.
4. form-file defines the name of the compiled form file, without
   .42f extension.
5. height defines the number of lines of the window in character units; includes
   the borders in character mode.
6. width defines the number of lines of the window in character units; includes
   the borders in character mode.

## Usage

An `OPEN WINDOW` statement
can have the following effects:

- Declares a name (the identifier) for the window.
- Indicates which form has to be used in that window.
- Specifies the display attributes of the window.
- When using character mode, specifies the position and dimensions
  of the window, in character units.

For graphical applications, use this instruction without the `AT` clause, and with
the `WITH FORM` clause.

The window identifier must follow the rules for identifiers and be unique among all windows
defined in the program. Its scope is the entire program. You can use this identifier to reference
the same Window in other modules with other statements (for example, `CURRENT WINDOW` and `CLOSE WINDOW`).

The compiler converts the window identifier to lowercase for internal storage. When using
functions or methods receiving the window identifier as a string parameter, the window name is case
sensitive. We recommend that you always specify the window identifier in lowercase letters.

The runtime system maintains a stack of
all open windows. If you execute `OPEN WINDOW` to
open a new window, it is added to the window stack and becomes the
current window. Other statements that can modify the window
stack are `CURRENT WINDOW` and `CLOSE
WINDOW`.

## Example

```
MAIN
  OPEN WINDOW w1 WITH FORM "customer"
  MENU "Test"
    COMMAND KEY(INTERRUPT) "exit" EXIT MENU
  END MENU
  CLOSE WINDOW w1
END MAIN
```

## Related links

**Related concepts**  

[Graphical mode with Traditional Display](1519-graphical-mode-with-traditional-display.md "Graphical mode with Traditional Display")

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")

[The Window class](../15_library-reference/3128-the-window-class.md "The ui.Window class provides an interface to the window objects created with the OPEN WINDOW instruction.")

[The Form class](../15_library-reference/3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction.")

## Child topics

- [OPEN WINDOW attributes](1573-open-window-attributes.md): List of attributes for the OPEN WINDOW instruction.
- [WITH FORM clause](1574-with-form-clause.md): Creating a window object with a form.
