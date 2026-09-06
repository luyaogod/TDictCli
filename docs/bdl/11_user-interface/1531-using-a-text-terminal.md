---
title: "Using a text terminal"
source: "fgl-topics/c_fgl_DynamicUI_029.html"
breadcrumb: "User interface > User interface basics > Using a text terminal"
type: "concept"
---

# Using a text terminal

> This section covers topics about text terminal configuration when using the TUI mode.

The TUI mode is enabled when setting the [FGLGUI](../07_configuration/0525-fglgui.md "Defines the user interface mode to be used by the program.") environment variable to zero.

Terminal type and terminal capabilities definition is not a Genero-specific
configuration: TERM, TERMCAP and TERMINFO are also used by other UNIX™ applications and commands.

On UNIX platforms, the TERM
environment variable must be set to define the terminal type/name.
For example, if you execute the application in an xterm X11 window,
set TERM=xterm.

On Windows® platforms, you can run applications in text
mode inside a CMD console window. You must not set the TERM environment variable in this case.

Genero supports both termcap and terminfo implementations
of text terminal capabilities. The INFORMIXTERM environment variable
defines the type of library used to interact with the terminal.
When INFORMIXTERM is set to `termcap` (the default),
the runtime system reads terminal capabilities from the file defined
by the TERMCAP environment variable. When INFORMIXTERM is set to `terminfo`,
the runtime system uses the ncurses library of the operating system
to interact with the terminal. We strongly recommend you to use
the terminfo solution.

## Related links

**Related concepts**  

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")

## Child topics

- [TERMINFO terminal capabilities](1532-terminfo-terminal-capabilities.md)
- [TERMCAP terminal capabilities](1533-termcap-terminal-capabilities.md)
- [Text mode screen dump](1536-text-mode-screen-dump.md)
