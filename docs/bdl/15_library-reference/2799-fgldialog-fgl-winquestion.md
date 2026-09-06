---
title: "fgldialog.fgl_winquestion()"
source: "fgl-topics/c_fgl_utility_functions_FGL_WINQUESTION.html"
breadcrumb: "Library reference > Utility modules > fgldialog: Common dialog functions > fgldialog.fgl_winquestion()"
type: "concept"
---

# fgldialog.fgl_winquestion()

> Displays an interactive message box with configurable Ok / Yes / No / Cancel / Ignore / Abort / Retry buttons.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION fgl_winquestion(
   title STRING,
   text STRING,
   default STRING,
   buttons STRING,
   icon STRING,
   danger SMALLINT )
  RETURNS STRING
```

1. title is the message box title.
2. text is the message displayed in the message box. Use `'\n'`
   to separate lines (does not work on ASCII client).
3. default defines the default button that is preselected.
4. buttons defines the options. Must be a pipe-separated list of 2 or three
   options: `ok, yes, no, cancel, abort, retry, ignore`.
5. icon is the name of the icon to be displayed.
6. danger is supported for backward compatibility. This parameter is ignored.

## Usage

The `fgl_winquestion()` function shows a question message box to the end user and
waits for an answer.

> **Important:**
>
> With front-ends implementing this function with a system dialog box API
> creating a modal window, the end user will have to close the modal window first, before continuing
> within the window of another program. Consider using a [menu with
> "dialog" style](../11_user-interface/1906-syntax-of-the-menu-instruction.md "The MENU instruction defines a set of options the end user can select to trigger actions in a program.") instead so as not to block other programs.

The function returns the label of the option which has been selected by the user.

Supported names for the icon parameter are: `information`,
`exclamation`, `question`, `stop`. Note that on some
front-ends such as iOS devices, the native message pop-up window does not display an image.

The buttons parameter defines the list of options that the user can
select. Possible values are: `ok`, `yes`, `no`,
`cancel`, `abort`, `retry`, `ignore`.
You must specify a pipe-separated list of options, with a maximum of 3 options. For example:
`"ok"`, `"yes|no"`, `"yes|no|cancel"`,
`"abort|retry|ignore"`.

> **Important:**
>
> To display the pop-up window of this API, desktop and mobile front-ends
> use the platform specific message box API, with a predefined set of buttons. Some non-standard
> option combinations may not be supported, such as `"ok|yes|abort"`. Furthermore,
> the order of the buttons depends also on platform standards. For example, with
> `"abort|retry|ignore"`, the buttons can appear in the following order:
> `[Retry] [Ignore] [Abort]`.

On front-ends using a system dialog box API, the option buttons will be automatically localized
based on the operating system language settings. On other front-ends, the option buttons will be
decorated to action default settings.

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
  LET answer = "yes"
  WHILE answer = "yes"
    LET answer = fgl_winquestion(
        "Procedure", "Would you like to continue ? ",
        "cancel", "yes|no|cancel", "question", 0)
  END WHILE
  IF answer = "cancel" THEN
     DISPLAY "Canceled."
  END IF
END MAIN
```

## Related links

**Related concepts**  

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
