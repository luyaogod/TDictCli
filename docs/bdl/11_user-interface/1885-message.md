---
title: "MESSAGE"
source: "fgl-topics/c_fgl_message_display_MESSAGE.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > MESSAGE"
type: "concept"
---

# MESSAGE

> The MESSAGE instruction displays a message to the user.

## Syntax

```
MESSAGE expression [,...]
  [ {ATTRIBUTE|ATTRIBUTES} ( display-attribute [,...] ) ]
```

where display-attribute is:

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

The `MESSAGE` instruction
displays a message to the user in an interactive program.

The message will remain visible to the end user until the next `MESSAGE`
instruction.

To cleanup the message text, use `MESSAGE ""` .

In TUI mode, the text is displayed in the message line of the current window. The message line
can be defined by the [`OPTIONS MESSAGE
LINE`](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.") instruction. Note that the default message line position is the same as the
`MENU` option comment line.

In GUI mode, you can use the `STYLE` attribute to
[reference a style
definition](1644-message-style-attributes.md "Message presentation style attributes apply to an ERROR or MESSAGE instruction."). This allows you to display errors or messages in GUI mode with more sophisticated
visual effects as the regular TTY attributes. If you want to apply a style automatically to all
program messages displayed with the `MESSAGE` instruction, you can use the [`:message`](1612-pseudo-selectors.md "Pseudo selectors can be used to apply style only when some conditions are fulfilled.") pseudo selector in the
style definition.

## Example

```
INPUT BY NAME custrec.* ...
  BEFORE INPUT
    MESSAGE "Enter customer data."
  ...
```

## Related links

**Related concepts**  

[ERROR](1886-error.md "The ERROR instruction displays an error message to the user.")

[Windows and forms](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.")

[Defining the position of reserved lines](../09_advanced-features/0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.")

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")
