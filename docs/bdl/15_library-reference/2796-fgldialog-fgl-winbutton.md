---
title: "fgldialog.fgl_winbutton()"
source: "fgl-topics/c_fgl_utility_functions_FGL_WINBUTTON.html"
breadcrumb: "Library reference > Utility modules > fgldialog: Common dialog functions > fgldialog.fgl_winbutton()"
type: "concept"
---

# fgldialog.fgl_winbutton()

> Displays an interactive message box containing multiple choices, in a pop-up window.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION fgl_winbutton(
   title STRING,
   text STRING,
   default STRING,
   buttons STRING,
   icon STRING,
   danger SMALLINT )
  RETURNS STRING
```

1. title defines the title of the message window.
2. text specifies the string displayed in message window.
3. default indicates the default button to be pre-selected.
4. buttons defines a set of button labels separated by "|".
5. icon is the name of the icon to be displayed.
6. danger (for X11 only), number of warning items. Otherwise, this parameter is
   ignored.

## Usage

Use the `fgl_winbutton()` function to open a message
box and let the end user select an option in a set of buttons.
The function returns the label of the button which has been selected
by the user.

Use '\n' in text to separate lines (this does
not work in TUI mode).

Supported names for the icon parameter are:
`information`, `exclamation`,
`question`, `stop`.

You can define up to 7 buttons that each have 10 characters.

If two buttons start with the same letter, the user will not be able
to select one of them in the TUI mode.

The "&" before a letter for a button is displayed in TUI mode,
or underlines the next letter in graphical front-ends.

This function is provided for backward compatibility. Consider using a [`MENU` with `STYLE="dialog"`](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.") as in
the next
example:

```
PUBLIC FUNCTION mbox_yn(t STRING, m STRING, i STRING) RETURNS BOOLEAN
  DEFINE r BOOLEAN
  MENU t ATTRIBUTES(STYLE="dialog",COMMENT=m, IMAGE=i)
      ON ACTION opt_yes ATTRIBUTES(TEXT=%"opt_yes")
        LET r = TRUE
      ON ACTION opt_no ATTRIBUTES(TEXT=%"opt_no")
        LET r = FALSE
  END MENU
  RETURN r
END FUNCTION
```

## Example

```
IMPORT FGL fgldialog
MAIN
  DEFINE answer STRING
  LET answer = fgl_winbutton( "Media selection", "What is your favorite media?",
      "Lynx", "Floppy Disk|CD-ROM|DVD-ROM|Other", "question", 0)
  DISPLAY "Selected media is: " || answer 
END MAIN
```
