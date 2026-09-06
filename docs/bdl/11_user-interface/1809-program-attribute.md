---
title: "PROGRAM attribute"
source: "fgl-topics/c_fgl_FSFAttributes_PROGRAM.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > PROGRAM attribute"
type: "concept"
---

# PROGRAM attribute

> The PROGRAM attribute can specify an external application program to edit TEXT or BYTE fields.

## Syntax

```
PROGRAM = "editor"
```

1. editor is the name of the program that must be used to edit the special field data.

## Usage

You can assign the `PROGRAM` attribute to a `TEXT` or `BYTE`
field to call an external program to work with the `BYTE` or `TEXT` values.

This attribute works in TUI mode only.

Users can invoke the external program by pressing the exclamation point (!) key while the screen cursor
is in the field.

The external program then takes over control of the screen. When the user exits from the external program,
the form is redisplayed with any display attributes besides `PROGRAM` in effect.

When no `PROGRAM` attribute is used, the DBEDIT environment variable defines the default
editor.

## Related links

**Related concepts**  

[TEXT](../08_language-basics/0569-text.md "The TEXT data type stores large text data.")

[BYTE](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.")

[DBEDIT](../07_configuration/0510-dbedit.md "Defines the editor program for TEXT fields in TUI mode.")

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")
