---
title: "ERROR"
source: "fgl-topics/c_fgl_message_display_ERROR.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > ERROR"
type: "concept"
---

# ERROR

> The ERROR instruction displays an error message to the user.

## Syntax

```
ERROR expression [,...]
  [ {ATTRIBUTE|ATTRIBUTES} ( display-attribute [,...] ) ]
```

where display-attribute
is:

```
{ BLACK | BLUE | CYAN | GREEN
| MAGENTA | RED | WHITE | YELLOW
| BOLD | DIM | INVISIBLE | NORMAL
| REVERSE | BLINK | UNDERLINE
| STYLE = "style-name"
}
```

1. expression is any expression supported by the
   language.
2. style-name is a presentation style name.

## Usage

The `ERROR` instruction displays an error message to the user in an interactive
program.

The error message will remain visible to the end user until the next user interaction like a move
to another field or an action execution.

To cleanup the error message text, use `ERROR ""` .

In TUI mode, the error text is displayed in the error line of the screen. The text of the
`ERROR` statement is always displayed independently to the current window. The error
line can be defined by the [`OPTIONS ERROR
LINE`](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.") instruction.

In GUI mode, you can use the `STYLE` attribute to
[reference a style
definition](1644-message-style-attributes.md "Message presentation style attributes apply to an ERROR or MESSAGE instruction."). This allows you to display errors or messages in GUI mode with more sophisticated
visual effects as the regular TTY attributes. If you want to apply a style automatically to all
program warnings displayed with the `ERROR` instruction, you can use the [`:error`](1612-pseudo-selectors.md "Pseudo selectors can be used to apply style only when some conditions are fulfilled.") pseudo selector in the
style definition.

## Example

```
...
UPDATE tab1 SET col1 = ...
IF sqlca.sqlcode < 0 THEN
  ERROR SFMT("Row update failed (err=%1)", sqlca.sqlcode)
     ATTRIBUTES(STYLE="important")
    ...
END IF
...
```

## Related links

**Related concepts**  

[MESSAGE](1885-message.md "The MESSAGE instruction displays a message to the user.")

[Windows and forms](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.")

[Defining the position of reserved lines](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.")

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")
