---
title: "Defining the position of reserved lines"
source: "fgl-topics/c_fgl_programs_006.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Defining the position of reserved lines"
type: "concept"
---

# Defining the position of reserved lines

> The OPTIONS element LINE defines position of dedicated screen lines.

## Syntax

```
OPTIONS
{ MENU LINE line-value
| MESSAGE LINE line-value
| COMMENT LINE {OFF|line-value}
| PROMPT LINE line-value
| ERROR LINE line-value
| FORM LINE line-value
}
```

## Usage

The `OPTIONS` statement can define the positions of reserved lines for menus,
forms and messages.

Reserved window lines are used in [TUI mode](../11_user-interface/1517-text-mode-rendering-tui-mode.md). These
options are not required in [GUI mode](../11_user-interface/1518-graphical-mode-rendering-gui-mode.md). In GUI mode,
these options have no effect, except when using the [traditional mode](../11_user-interface/1519-graphical-mode-with-traditional-display.md), where program windows are rendered as in a dumb terminal.

- `COMMENT LINE` specifies the position of the comments for fields. Field comments
  are defined with the [`COMMENT`](../11_user-interface/1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element.")
  attribute in the form specification file. The default is (`LAST-1`) for the
  `SCREEN`, and `LAST` for all other windows. The field comment display
  can be disabled with `COMMENT LINE OFF`.
- `ERROR LINE` specifies the position on the screen for the text of the [`ERROR`](../11_user-interface/1886-error.md "The ERROR instruction displays an error message to the user.") statement. The text of the
  `ERROR` statement is always displayed independently to the current window. The
  default is the `LAST` line of the screen.
- `MESSAGE LINE` specifies the position of the message line in the current window.
  This reserved line displays the text of the [`MESSAGE`](../11_user-interface/1885-message.md "The MESSAGE instruction displays a message to the user.") statement. The default is `FIRST+1` (line 2 in the
  current window). Note that the default message line position is the same as the
  `MENU` option comment line.
- `FORM LINE` specifies the window line where [forms](../11_user-interface/1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.") are displayed. The default is
  `FIRST+2` (line 3 in the current window).
- `MENU LINE` specifies the position of the menu line in the current window. This
  line displays the menu name and options, as defined by the [`MENU`](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.") statement. The default is the `FIRST` line in the
  current window.
- `PROMPT LINE` specifies the position of the prompt line where the text of [`PROMPT`](../11_user-interface/1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window.") statements is displayed. The default
  value is the `FIRST` line in the current window.

You can specify any of the following positions for each reserved line:

| Expression | Description |
| --- | --- |
| `FIRST` | The first line of the screen or window. |
| `FIRST + integer` | A relative line position from the first line. |
| `integer` | An absolute line position in the screen or window. |
| `LAST - integer` | A relative line position from the last line. |
| `LAST` | The last line of the screen or window. |

## Related links

**Related concepts**  

[Genero user interface modes](../11_user-interface/1516-genero-user-interface-modes.md "User interface modes allow you to adapt the application form rendering to different types of displays.")
