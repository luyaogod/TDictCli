---
title: "fgldialog.fgl_winmessage()"
source: "fgl-topics/c_fgl_utility_functions_FGL_WINMESSAGE.html"
breadcrumb: "Library reference > Utility modules > fgldialog: Common dialog functions > fgldialog.fgl_winmessage()"
type: "concept"
---

# fgldialog.fgl_winmessage()

> Displays an interactive message box containing text and OK button.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
FUNCTION fgl_winmessage(
   title STRING,
   text STRING,
   icon STRING )
```

1. title defines message box title.
2. text is the text displayed in the message box. Use '\n'
   to separate lines.
3. icon is the name of the icon to be displayed.

## Usage

The `fgl_winmessage()` function displays a message box to the end user.

> **Important:**
>
> With front-ends implementing this function with a system dialog box API
> creating a modal window, the end user will have to close the modal window first, before continuing
> within the window of another program. Consider using a [menu with
> "dialog" style](../11_user-interface/1906-syntax-of-the-menu-instruction.md "The MENU instruction defines a set of options the end user can select to trigger actions in a program.") instead so as not to block other programs.

Supported names for the icon parameter are: `information`,
`exclamation`, `question`, `stop`. Note that on some
front-ends such as iOS devices, the native message pop-up window does not display an image.

On front-ends using a system dialog box API, the OK buttons will be automatically localized based
on the operating system language settings. On other front-ends, the option buttons will be decorated
depending on action default settings.

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
  CALL fgl_winmessage( "Title", "This is a critical message.", "stop")
END MAIN
```

## Related links

**Related concepts**  

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
