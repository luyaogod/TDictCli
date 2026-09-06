---
title: "DISPLAY (to stdout)"
source: "fgl-topics/c_fgl_message_display_DISPLAY.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > DISPLAY (to stdout)"
type: "concept"
---

# DISPLAY (to stdout)

> The DISPLAY instruction displays text in line mode to the standard output channel.

## Syntax

```
DISPLAY expression [,...]
```

1. expression is any expression supported by the language.

## Usage

The `DISPLAY` instruction can be used to print information to the standard output
channel (stdout) of the terminal the program is attached to.

Before displaying to the standard output channel, the expression is converted to a character
string. The values contained in variables are formatted depending on the data types and environment
settings such as [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") and [DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.").

When using the TUI mode (FGLGUI=0), a `DISPLAY` to stdout switches the program to
`LINE MODE`, which impacts subsequent instructions such as the `RUN`
command. For more details about the terminal mode, see [IN LINE MODE and IN FORM MODE](../09_advanced-features/0830-run.md) .

## Example

```
MAIN
  DISPLAY "Today's date is: ", TODAY
END MAIN
```

## Related links

**Related concepts**  

[The Channel class](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
