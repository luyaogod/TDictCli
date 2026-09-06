---
title: "showhelp()"
source: "fgl-topics/c_fgl_BuiltInFunctions_SHOWHELP.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > showhelp()"
type: "concept"
---

# showhelp()

> Displays a runtime help text.

## Syntax

```
FUNCTION showhelp(
   number INTEGER )
```

1. number is the help message number in the current help file.

## Usage

The `showhelp()` function displays a runtime help text, corresponding
to its specified argument, from the current help file defined by the
[`OPTIONS HELP FILE`](../09_advanced-features/0931-defining-the-message-file.md "The OPTIONS HELP FILE instruction defines the name of the message file.")
instruction.

In GUI mode, the help text will be displayed in a new pop-up window. In TUI mode, the help text
is displayed in the whole screen.

## Related links

**Related concepts**  

[Message files](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.")
